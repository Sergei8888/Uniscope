package Sky

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Azimuth struct {
	degrees int
	minutes int
	seconds int
}

func NewAzimuth(degrees int, minutes int, seconds int) (*Azimuth, error) {
	if degrees < 0 || degrees > 360 {
		return nil, errors.New(fmt.Sprintf("azimuth: degrees should be from 0 to 360, got: %d", degrees))
	}
	if minutes < 0 || minutes > 60 {
		return nil, errors.New(fmt.Sprintf("azimuth: minutes should be from 0 to 60, got: %d", minutes))
	}
	if seconds < 0 || seconds > 60 {
		return nil, errors.New(fmt.Sprintf("azimuth: seconds should be from 0 to 60, got: %d", seconds))
	}
	return &Azimuth{
			degrees: degrees,
			minutes: minutes,
			seconds: seconds,
		},
		nil
}

func (azm *Azimuth) Degrees() int {
	return azm.degrees
}

func (azm *Azimuth) Minutes() int {
	return azm.minutes
}

func (azm *Azimuth) Seconds() int {
	return azm.seconds
}

func (azm *Azimuth) SetDegrees(degrees int) {
	azm.degrees = degrees
}

func (azm *Azimuth) SetMinutes(minutes int) {
	azm.minutes = minutes
}

func (azm *Azimuth) SetSeconds(seconds int) {
	azm.seconds = seconds
}

func (azm *Azimuth) ToString() string {
	return fmt.Sprintf("%d°%d'%d\"", azm.degrees, azm.minutes, azm.seconds)
}

func (azm *Azimuth) ConvToArcSeconds() int {
	return azm.degrees*3600 + azm.minutes*60 + azm.seconds
}

func (azm *Azimuth) ToHex() ([]byte, error) {
	degreesFloat := (float64(azm.degrees)*65536 + float64(azm.minutes)*65536/60 +
		float64(azm.minutes)*65536/3600) / 360

	degreesInt := int64(math.Round(degreesFloat))

	hex := []byte(strings.ToUpper(strconv.FormatInt(degreesInt, 16)))

	for len(hex) < 4 {
		hex = append([]byte{'0'}, hex...)
	}

	return hex, nil
}

func (azm *Azimuth) ToHexPrecise() ([]byte, error) {
	degreesFloat := (float64(azm.degrees) + float64(azm.minutes)/60 +
		float64(azm.minutes)/3600) / 360 * 16777216

	degreesInt := int64(math.Round(degreesFloat))

	hex := []byte(strings.ToUpper(strconv.FormatInt(degreesInt, 16)))

	for len(hex) < 6 {
		hex = append([]byte{'0'}, hex...)
	}

	return append(hex, '0', '0'), nil
}
