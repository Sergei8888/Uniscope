import sqlite3

from requests import get as requests_get
from config import *
from sqlite3 import connect
from time import time as time_now
import json

connection = connect(database)
cursor = connection.cursor()
group_token_last = 0

sorted_params = {
    "last_online": -1,  # days ago
    "city": [],  # array of strings
    "country": [],  # array of strings
    "messages": -1,  # 1 - if can send direct messages
    "friends": -1,  # 1 - if can send friend requests
    "groups": -1  # the number of groups the user belongs to
}


def show_sort_params():
    out = "Loaded params:\n"
    for i in sorted_params.keys():
        out += f"{i} = {sorted_params[i]}\n"
    return out


def request_to_vk(request, params=None):
    url = f"https://api.vk.com/method/{request}?v={version}&access_token={user_token}"
    if params:
        for i in params.keys():
            url += f"&{i}={params[i]}"
    return requests_get(url).json()


def add_group(args):
    # Checking is argument URL
    if not "://" in args[0]:
        return "usage: add_group {url}"

    # getting group_id
    if "vk.com/club" in args[0]:
        group_id = args[0].split("/club")[-1]
    else:
        group_id = args[0].split("/")[-1]
    res = request_to_vk("groups.getById", {"group_id": group_id})
    try:
        cursor.execute(f"INSERT INTO groups VALUES ({res['response'][0]['id']}, '{res['response'][0]['name']}', 0)")
    except sqlite3.Error:
        print(res['response'][0]['name'] + " already in database")
    connection.commit()
    return f"added '{res['response'][0]['name']}' in database"


def add_groups(args):
    response = ""
    if args[-1] == "-f":
        with open(args[-2], "r") as file:
            for i in file.readlines():
                response += add_group([i.strip()]) + "\n"
    else:
        for i in args:
            response += add_group([i]) + "\n"
    return response


def get_groups(args):
    rs = cursor.execute("SELECT * FROM groups")
    answer = ""
    for i in rs.fetchall():
        answer += f"Group: {i[1]} url: https://vk.com/club{i[0]} parsed: 0\n"
    if len(args) == 2:
        if args[1] == "-f":
            with open(args[0], "w+") as file:
                file.write(answer)
    return answer


def parse_groups(args):
    continue_from = 0
    if len(args) == 2:
        try:
            continue_from = int(args[0])
            rs = cursor.execute("SELECT group_id, group_name FROM groups").fetchall()
        except ValueError:
            return "argument error"
    if len(args) == 1:
        if "-all" in args:
            rs = cursor.execute("SELECT group_id, group_name FROM groups").fetchall()
        else:
            continue_from = int(args[0])
            rs = cursor.execute("SELECT group_id, group_name FROM groups WHERE parsed = 0").fetchall()
    else:
        rs = cursor.execute("SELECT group_id, group_name FROM groups WHERE parsed = 0").fetchall()
    if len(rs) == 0:
        return "Nothing to parse"
    print("Starting parse groups: " + ", \n".join([i[1] for i in rs]))
    for i in rs:
        req = request_to_vk("groups.getMembers", {"group_id": i[0]})
        repeats = int(req['response']['count'] / 1000) + 1
        if "error" in req:
            print("ERROR")
            continue
        count = continue_from
        for repeat in range(repeats):
            req = request_to_vk("groups.getMembers", {
                "fields": "can_write_private_message, can_send_friend_request, city, country, last_seen",
                "group_id": i[0], "offset": 1000 * repeat + continue_from})
            for j in req["response"]["items"]:
                count += 1
                if "deactivated" in j:
                    continue
                cursor.execute(f"DELETE FROM users WHERE user_id = {j['id']}")

                user_id = j['id']
                last_seen = "NULL"
                city = "NULL"
                country = "NULL"
                private_message = j['can_write_private_message']
                friend_request = j['can_send_friend_request']

                if "city" in j:
                    city = j['city']['title']
                if "country" in j:
                    country = j['country']['title']
                if "last_seen" in j:
                    last_seen = j['last_seen']['time']
                city = city.replace("`", "").replace("'", "").replace('"', "").replace("'", "")
                country = country.replace("`", "").replace("'", "").replace('"', "")
                cursor.execute(
                    f"INSERT INTO users VALUES({user_id}, {last_seen}, '{city}','{country}', {friend_request},"
                    f" {private_message})")
                try:
                    cursor.execute(f"INSERT INTO users_groups VALUES({user_id},{i[0]})")
                except sqlite3.Error:
                    pass
                print(f"Parsing {i[1]} group ({count}/{req['response']['count']})")
                connection.commit()
        continue_from = 0
        cursor.execute(f"UPDATE groups SET parsed = 1 WHERE group_id = {i[0]}")
    return "parsing finished"


def load_params(args=None):
    if args is None:
        args = ["sort_params.txt"]
    global sorted_params
    if len(args) == 0:
        args.append("sort_params.txt")
    try:
        with open(args[0], "r") as file:
            json_file = file.read().replace("\n", "").replace(" ", "")
            sorted_params = json.loads(json_file)
            print(show_sort_params())
            return f"params loaded from {args[0]}"
    except FileNotFoundError:
        return f"'{args[0]}' not found"
    except Exception:
        return "file reading error"


def get_users(args):
    params = []
    time_func_call = time_now()

    if sorted_params["last_online"] != -1:
        params.append(f"({int(time_now())} - user_last_seen < {sorted_params['last_online'] * 86400})")
    if sorted_params["city"]:
        _in = "'" + "', '".join(sorted_params["city"]) + "'"
        params.append(f"user_city IN ({_in})")
    if sorted_params["country"]:
        _in = "'" + "', '".join(sorted_params["country"]) + "'"
        params.append(f"user_country IN ({_in})")
    if sorted_params["friends"] != -1:
        params.append(f"user_friend_request = {sorted_params['friends']}")
    if sorted_params["messages"] != -1:
        params.append(f"user_private_message = {sorted_params['messages']}")

    result = ""
    sql_req = f"DROP TABLE IF EXISTS result_table;" + \
              "DROP TABLE IF EXISTS table_sorted_users_by_groups; " + \
              "CREATE TABLE table_sorted_users_by_groups AS SELECT user_id, COUNT(user_id) AS join_groups " + \
              "FROM users_groups GROUP BY users_groups.user_id ORDER BY join_groups;" + \
              "CREATE TABLE result_table AS SELECT * FROM users WHERE user_id in " + \
              f"(SELECT table_sorted_users_by_groups.user_id FROM table_sorted_users_by_groups WHERE join_groups >= {sorted_params['groups']});"
    cursor.executescript(sql_req)
    if params:
        sql_req = "SELECT * FROM result_table INNER JOIN table_sorted_users_by_groups USING(user_id) WHERE(" \
                  + ") AND (".join(params) + \
                  ")GROUP BY user_id ORDER BY join_groups"
    else:
        return "please enter parameters"
    rs = cursor.execute(sql_req)
    for i in rs.fetchall():
        place = "unknown"
        if i[3] != "NULL" and i[3] != None:
            if i[2] != "NULL" and i[2] != None:
                place = i[2] + "," + i[3]
            else:
                place = i[2]
        if i[1] == None:
            last_online = "unknown"
        else:
            last_online = (int(time_func_call) - i[1]) // 86400
        result += (f"https://vk.com/id{str(i[0]).ljust(10, ' ')}| {place.ljust(20, ' ')}"
                   f"| last online {str(last_online).ljust(5, ' ')} days ago| member in {i[6]} groups\n")
    if len(args) == 1:
        with open(args[0], "w+") as file:
            file.write(result)
    return result


if __name__ == "__main__":
    load_params()
    connection.commit()
    command, args = "", []
    while command != "exit":
        input_string = input()
        command, *args = input_string.split()
        try:
            print(globals()[command](args))
        except KeyError as err:
            if command != "exit":
                print("Unknown command")
    connection.close()
