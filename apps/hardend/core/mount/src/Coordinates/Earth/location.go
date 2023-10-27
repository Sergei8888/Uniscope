package Earth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Location struct {
	latitude  *Latitude
	longitude *Longitude
}

func NewLocation() *Location {
	return &Location{
		latitude:  &Latitude{},
		longitude: &Longitude{},
	}
}

// SetByString takes argument like: 118 20 17 W 33 50 41 N
func (l *Location) SetByString(locStr string) error {
	locStr = strings.Trim(locStr, " \t\n")
	coords := strings.Split(locStr, " ")
	fmt.Println(len(coords), coords)
	if len(coords) != 8 {
		return errors.New(fmt.Sprintf("invalid arguments number: expected 8, got %d", len(coords)))
	}
	if coords[3] != "W" && coords[3] != "E" {
		return errors.New("invalid longitude's position")
	}
	if coords[7] != "S" && coords[7] != "N" {
		return errors.New("invalid latitude's position")
	}

	fromWest := coords[3] == "W"
	fromSouth := coords[7] == "S"

	d, m, s, err := coordinatesStrToInt(coords[0], coords[1], coords[2])
	if err != nil {
		return err
	}
	latitude, err := NewLatitude(d, m, s, fromWest)
	if err != nil {
		return err
	}

	d, m, s, err = coordinatesStrToInt(coords[4], coords[5], coords[6])
	if err != nil {
		return err
	}
	fmt.Println("latitude:", latitude)
	longitude, err := NewLongitude(d, m, s, fromSouth)
	if err != nil {
		return err
	}

	fmt.Println("longitude:", longitude)
	l.latitude = latitude
	l.longitude = longitude

	return nil
}

func coordinatesStrToInt(dStr, mStr, sStr string) (d, m, s int, err error) {
	fmt.Println(dStr, mStr, sStr)
	d, err = strconv.Atoi(dStr)
	if err != nil {
		return 0, 0, 0, err
	}
	m, err = strconv.Atoi(mStr)
	if err != nil {
		return 0, 0, 0, err
	}
	s, err = strconv.Atoi(sStr)
	if err != nil {
		return 0, 0, 0, err
	}
	return d, m, s, nil
}

func (l *Location) Latitude() *Latitude {
	return l.latitude
}

func (l *Location) Longitude() *Longitude {
	return l.longitude
}

func (l *Location) ToString() string {

	return fmt.Sprintf("%d°%d’%d” %s, %d°%d’%d” %s",
		l.Longitude().Degrees(),
		l.Longitude().Minutes(),
		l.Longitude().Seconds(),
		l.Longitude().EastWestToString(),
		l.Latitude().Degrees(),
		l.Latitude().Minutes(),
		l.Latitude().Seconds(),
		l.Latitude().SouthNorthToString(),
	)
}

func (l *Location) SetLatitude(latitude *Latitude) {
	l.latitude = latitude
}

func (l *Location) SetLongitude(longitude *Longitude) {
	l.longitude = longitude
}
