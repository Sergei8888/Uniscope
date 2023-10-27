package main

import (
	"bufio"
	"core/mount/configs"
	"core/mount/src/Coordinates/Earth"
	"core/mount/src/Coordinates/Sky"
	mount "core/mount/src/Mount"
	celestron "core/mount/src/Mount/Commander/Celestron"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	port, err := mount.ConfigurePort(configs.PORT_NAME, configs.MODE)
	if err != nil {
		log.Fatalln(err)
	}

	mount := celestron.NewMountCelestron(port)

	if err = mount.PrepareAfterTurnOn(); err != nil {
		log.Fatalln(err)
	}

	for true {
		fmt.Print("Enter text: ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.Trim(input, "\n")

		args := strings.Split(input, " ")

		switch args[0] {

		case "time":
			t, err := mount.GetTime()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(t)
			break
		case "pos":
			horizPosCS, err := mount.GetAzmAlt()
			if err != nil {
				log.Fatalln(err)
			}
			eqCS, err := mount.GetRaDec()
			if err != nil {
				log.Fatalln(err)
			}
			switch args[1] {
			case "horiz":
				fmt.Printf("AZM: %s, ALT: %s\n", horizPosCS.Azm().ToString(), horizPosCS.Alt().ToString())
				break
			case "eq":
				fmt.Printf("RA: %s, DEC: %s\n", eqCS.Ra().ToString(), eqCS.Dec().ToString())
				break
			}
			break
		case "goto":
			if len(args) < 8 {
				log.Println("Wrong arguments length")
				break
			}
			var argInt int
			var argInts []int
			for _, arg := range args[2:] {
				argInt, err = strconv.Atoi(arg)
				if err != nil {
					err = errors.New(fmt.Sprintf("Wrong argument: %v", err))
					break
				}
				if argInt < 0 || argInt > 360 {
					err = errors.New(fmt.Sprintf("Wrong argument: expected arg between 0 and 360, got %d", argInt))
				}
				argInts = append(argInts, argInt)
			}
			if err != nil {
				log.Println(err)
				break
			}
			switch args[1] {
			case "horiz":
				azm, err := Sky.NewAzimuth(argInts[0], argInts[1], argInts[2])
				if err != nil {
					log.Println(err)
					break
				}
				alt, err := Sky.NewAltitude(argInts[3], argInts[4], argInts[5])
				if err != nil {
					log.Println(err)
					break
				}
				if err = mount.GotoAzmAlt(Sky.NewHorizontalCS(azm, alt)); err != nil {
					log.Println(err)
				}
				break
			case "eq":
				ra, err := Sky.NewRightAscension(argInts[0], argInts[1], argInts[2])
				if err != nil {
					log.Println(err)
					break
				}
				dec, err := Sky.NewDeclination(argInts[3], argInts[4], argInts[5])
				if err != nil {
					log.Println(err)
					break
				}

				rqCS := Sky.NewEquatorialCS(ra, dec)
				if err = mount.GotoRaDec(rqCS); err != nil {
					log.Println(err)
				}
				break
			}
			inProgress, err := mount.IsGotoInProgress()
			for inProgress {
				time.Sleep(5 * time.Second)
				if err != nil {
					log.Println(err)
					break
				}
				log.Println("Goto is in progress...")
				inProgress, err = mount.IsGotoInProgress()
			}
			break
		case "loc":
			switch args[1] {
			case "get":
				earthCoord, err := mount.GetLocation()
				if err != nil {
					log.Println(err)
					break
				}
				fmt.Println(earthCoord.ToString())
				break
			case "set":
				if len(args) < 8 {
					log.Println("Wrong arguments length")
					break
				}
				earthCoord := Earth.NewLocation()
				err = earthCoord.SetByString(strings.Join(args[2:], " "))
				fmt.Println(earthCoord)
				if err != nil {
					log.Println(err)
					break
				}
				err = mount.SetLocation(earthCoord)
				if err != nil {
					log.Println(err)
					break
				}
			}
		case "slew":
			switch args[1] {
			case "fixed":
				if len(args) < 3 {
					log.Println("Wrong arguments length")
					break
				}
				rate, err := strconv.Atoi(args[2])
				if err != nil {
					log.Println(err)
					break
				}
				mount.FixedRateSlew(celestron.AZM_RA, true, byte(rate))
			}

		}

		if input == "0" {
			break
		}
	}

	os.Exit(0)
}
