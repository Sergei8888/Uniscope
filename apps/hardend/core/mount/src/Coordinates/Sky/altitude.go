package Sky

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Altitude struct {
	degrees int
	minutes int
	seconds int
}

func NewAltitude(degrees int, minutes int, seconds int) (*Altitude, error) {
	//if degrees < 0 || degrees > 90 {
	//	return nil, errors.New(fmt.Sprintf("altitude: degrees should be from 0 to 90, got: %d", degrees))
	//}
	if degrees < 0 || degrees > 360 {
		return nil, errors.New(fmt.Sprintf("altitude: degrees should be from 0 to 360, got: %d", degrees))
	}

	if minutes < 0 || minutes > 60 {
		return nil, errors.New(fmt.Sprintf("altitude: minutes should be from 0 to 60, got: %d", minutes))
	}
	if seconds < 0 || seconds > 60 {
		return nil, errors.New(fmt.Sprintf("altitude: seconds should be from 0 to 60, got: %d", seconds))
	}
	return &Altitude{
			degrees: degrees,
			minutes: minutes,
			seconds: seconds,
		},
		nil
}

func (alt *Altitude) Degrees() int {
	return alt.degrees
}

func (alt *Altitude) Minutes() int {
	return alt.minutes
}

func (alt *Altitude) Seconds() int {
	return alt.seconds
}

func (alt *Altitude) SetDegrees(degrees int) {
	alt.degrees = degrees
}

func (alt *Altitude) SetMinutes(minutes int) {
	alt.minutes = minutes
}

func (alt *Altitude) SetSeconds(seconds int) {
	alt.seconds = seconds
}

func (alt *Altitude) ToString() string {
	return fmt.Sprintf("%d°%d'%d\"", alt.degrees, alt.minutes, alt.seconds)
}

func (alt *Altitude) ConvToArcSeconds() int {
	return alt.degrees*3600 + alt.minutes*60 + alt.seconds
}

func (alt *Altitude) ToHex() ([]byte, error) {
	degreesFloat := (float64(alt.degrees) + float64(alt.minutes)/60 +
		float64(alt.minutes)/3600) / 360 * 65536

	degreesInt := int64(math.Round(degreesFloat))

	hex := []byte(strings.ToUpper(strconv.FormatInt(degreesInt, 16)))

	for len(hex) < 4 {
		hex = append([]byte{'0'}, hex...)
	}

	return hex, nil
}

func (alt *Altitude) ToHexPrecise() ([]byte, error) {
	degreesFloat := (float64(alt.degrees) + float64(alt.minutes)/60 +
		float64(alt.minutes)/3600) / 360 * 16777216

	degreesInt := int64(math.Round(degreesFloat))

	hex := []byte(strings.ToUpper(strconv.FormatInt(degreesInt, 16)))

	for len(hex) < 6 {
		hex = append([]byte{'0'}, hex...)
	}

	return append(hex, '0', '0'), nil
}
