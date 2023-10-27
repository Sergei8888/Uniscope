package Earth

import "errors"

type Longitude struct {
	degrees byte
	minutes byte
	seconds byte
	ew      byte // East or West
}

func NewLongitude(degrees, minutes, seconds int, isWestPosition bool) (*Longitude, error) {
	if !IsLongitudeDegreesValid(degrees) {
		return nil, errors.New("invalid degrees")
	}
	if !IsLongitudeMinutesValid(minutes) {
		return nil, errors.New("invalid minutes")
	}
	if !IsLongitudeSecondsValid(seconds) {
		return nil, errors.New("invalid seconds")
	}

	var ew byte
	if isWestPosition {
		ew = 0x01
	}

	return &Longitude{
		degrees: byte(degrees),
		minutes: byte(minutes),
		seconds: byte(seconds),
		ew:      ew,
	}, nil
}

func (l *Longitude) EastWestPosition() byte {
	return l.ew
}

func (l *Longitude) EastWestToString() string {
	if l.ew == 0 {
		return "E"
	}
	return "W"
}

func (l *Longitude) Minutes() byte {
	return l.minutes
}

func (l *Longitude) Seconds() byte {
	return l.seconds
}

func (l *Longitude) Degrees() byte {
	return l.degrees
}

func IsLongitudeDegreesValid(degrees int) bool {
	return degrees >= 0 && degrees <= 180
}

func IsLongitudeMinutesValid(minutes int) bool {
	return minutes >= 0 && minutes <= 60
}

func IsLongitudeSecondsValid(seconds int) bool {
	return seconds >= 0 && seconds <= 60
}
