package Earth

import "errors"

type Latitude struct {
	degrees byte
	minutes byte
	seconds byte
	sn      byte // South or North
}

func NewLatitude(degrees, minutes, seconds int, isSouthPosition bool) (*Latitude, error) {
	if !IsLatitudeDegreesValid(degrees) {
		return nil, errors.New("invalid degrees")
	}
	if !IsLatitudeMinutesValid(minutes) {
		return nil, errors.New("invalid minutes")
	}
	if !IsLatitudeSecondsValid(seconds) {
		return nil, errors.New("invalid seconds")
	}

	var sn byte
	if isSouthPosition {
		sn = 0x01
	}

	return &Latitude{
		degrees: byte(degrees),
		minutes: byte(minutes),
		seconds: byte(seconds),
		sn:      sn,
	}, nil
}

func (l *Latitude) SouthNorthPosition() byte {
	return l.sn
}

func (l *Latitude) SouthNorthToString() string {
	if l.sn == 0 {
		return "N"
	}
	return "S"
}

func (l *Latitude) Degrees() byte {
	return l.degrees
}

func (l *Latitude) Minutes() byte {
	return l.minutes
}

func (l *Latitude) Seconds() byte {
	return l.seconds
}

func IsLatitudeDegreesValid(degrees int) bool {
	return degrees >= 0 && degrees <= 90
}

func IsLatitudeMinutesValid(minutes int) bool {
	return minutes >= 0 && minutes <= 60
}

func IsLatitudeSecondsValid(seconds int) bool {
	return seconds >= 0 && seconds <= 60
}
