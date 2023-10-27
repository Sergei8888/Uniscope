package Sky

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RightAscension struct {
	degrees int
	minutes int
	seconds int
}

func NewRightAscension(degrees int, minutes int, seconds int) (*RightAscension, error) {
	if degrees < 0 || degrees > 360 {
		return nil, errors.New(fmt.Sprintf("right ascension: degrees should be from 0 to 360, got: %d", degrees))
	}
	if minutes < 0 || minutes > 60 {
		return nil, errors.New(fmt.Sprintf("right ascension: minutes should be from 0 to 60, got: %d", minutes))
	}
	if seconds < 0 || seconds > 60 {
		return nil, errors.New(fmt.Sprintf("right ascension: seconds should be from 0 to 60, got: %d", seconds))
	}
	return &RightAscension{
			degrees: degrees,
			minutes: minutes,
			seconds: seconds,
		},
		nil
}

func (ra *RightAscension) Degrees() int {
	return ra.degrees
}

func (ra *RightAscension) Minutes() int {
	return ra.minutes
}

func (ra *RightAscension) Seconds() int {
	return ra.seconds
}

func (ra *RightAscension) SetDegrees(degrees int) {
	ra.degrees = degrees
}

func (ra *RightAscension) SetMinutes(minutes int) {
	ra.minutes = minutes
}

func (ra *RightAscension) SetSeconds(seconds int) {
	ra.seconds = seconds
}

func (ra *RightAscension) ToString() string {
	return fmt.Sprintf("%d°%d'%d\"", ra.degrees, ra.minutes, ra.seconds)
}

func (ra *RightAscension) ConvToArcSeconds() int {
	return ra.degrees*3600 + ra.minutes*60 + ra.seconds
}

func (ra *RightAscension) ToHex() ([]byte, error) {
	degreesFloat := (float64(ra.degrees) + float64(ra.minutes)/60 +
		float64(ra.minutes)/3600) / 360 * 65536

	degreesInt := int64(math.Round(degreesFloat))

	hex := []byte(strings.ToUpper(strconv.FormatInt(degreesInt, 16)))

	for len(hex) < 4 {
		hex = append([]byte{'0'}, hex...)
	}

	fmt.Println(hex, string(hex))

	return hex, nil
}

func (ra *RightAscension) ToHexPrecise() ([]byte, error) {
	degreesFloat := (float64(ra.degrees) + float64(ra.minutes)/60 +
		float64(ra.minutes)/3600) / 360 * 16777216

	degreesInt := int64(math.Round(degreesFloat))

	hex := []byte(strings.ToUpper(strconv.FormatInt(degreesInt, 16)))

	for len(hex) < 6 {
		hex = append([]byte{'0'}, hex...)
	}

	return append(hex, '0', '0'), nil
}
