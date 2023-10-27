package Sky

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Declination struct {
	degrees int
	minutes int
	seconds int
}

func NewDeclination(degrees int, minutes int, seconds int) (*Declination, error) {
	//if degrees < 0 || degrees > 90 {
	//	return nil, errors.New(fmt.Sprintf("declination: degrees should be from 0 to 90, got: %d", degrees))
	//}
	if degrees < 0 || degrees > 360 {
		return nil, errors.New(fmt.Sprintf("declination: degrees should be from 0 to 360, got: %d", degrees))
	}
	if minutes < 0 || minutes > 60 {
		return nil, errors.New(fmt.Sprintf("declination: minutes should be from 0 to 60, got: %d", minutes))
	}
	if seconds < 0 || seconds > 60 {
		return nil, errors.New(fmt.Sprintf("declination: seconds should be from 0 to 60, got: %d", seconds))
	}
	return &Declination{
			degrees: degrees,
			minutes: minutes,
			seconds: seconds,
		},
		nil
}

func (dec *Declination) Degrees() int {
	return dec.degrees
}

func (dec *Declination) Minutes() int {
	return dec.minutes
}

func (dec *Declination) Seconds() int {
	return dec.seconds
}

func (dec *Declination) SetDegrees(degrees int) {
	dec.degrees = degrees
}

func (dec *Declination) SetMinutes(minutes int) {
	dec.minutes = minutes
}

func (dec *Declination) SetSeconds(seconds int) {
	dec.seconds = seconds
}

func (dec *Declination) ToString() string {
	return fmt.Sprintf("%d°%d'%d\"", dec.degrees, dec.minutes, dec.seconds)
}

func (dec *Declination) ConvToArcSeconds() int {
	return dec.degrees*3600 + dec.minutes*60 + dec.seconds
}

func (dec *Declination) ToHex() ([]byte, error) {
	degreesFloat := (float64(dec.degrees) + float64(dec.minutes)/60 +
		float64(dec.minutes)/3600) / 360 * 65536

	degreesInt := int64(math.Round(degreesFloat))

	hex := []byte(strings.ToUpper(strconv.FormatInt(degreesInt, 16)))

	for len(hex) < 4 {
		hex = append([]byte{'0'}, hex...)
	}

	fmt.Println(hex, string(hex))

	return hex, nil
}

func (dec *Declination) ToHexPrecise() ([]byte, error) {
	degreesFloat := (float64(dec.degrees) + float64(dec.minutes)/60 +
		float64(dec.minutes)/3600) / 360 * 16777216

	degreesInt := int64(math.Round(degreesFloat))

	hex := []byte(strings.ToUpper(strconv.FormatInt(degreesInt, 16)))

	for len(hex) < 6 {
		hex = append([]byte{'0'}, hex...)
	}

	return append(hex, '0', '0'), nil
}
