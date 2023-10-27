package mount

import (
	"errors"
	"fmt"
)

type Motor struct {
	raAzm        float64
	decAlt       float64
	isHorizontal bool
}

func NewMotor(raAzm, decAlt float64) (*Motor, error) {
	if raAzm < 0 || raAzm > 360 {
		return nil, errors.New(fmt.Sprintf("invalid RA: expected between 0 and 360, got %f", raAzm))
	}
	if decAlt < 0 || decAlt > 360 {
		return nil, errors.New(fmt.Sprintf("invalid AZM: expected between 0 and 360, got %f", decAlt))
	}
	return &Motor{raAzm: raAzm, decAlt: decAlt}, nil
}

func (h Motor) RaAzm() float64 {
	return h.raAzm
}

func (h Motor) DecAlt() float64 {
	return h.decAlt
}
