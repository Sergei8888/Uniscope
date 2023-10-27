package mount

import (
	"errors"
	"fmt"
	"math"
)

type Rotation struct {
	degrees float64
}

func NewRotation(degrees float64) (*Rotation, error) {
	if degrees < 0 || degrees > 360 {
		return nil, errors.New(fmt.Sprintf("invalid degrees' values: expected between 0 and 360, got %f", degrees))
	}
	return &Rotation{degrees: degrees}, nil
}

// Decompose returns rotation presented by degrees, minutes and seconds
func (r Rotation) Decompose() (degrees, minutes, seconds int) {
	degreesF, minutesF := math.Modf(r.degrees)
	minutesF *= 60
	minutesF, secsF := math.Modf(minutesF)
	secsF *= 60
	return int(degreesF), int(minutesF), int(secsF)
}
