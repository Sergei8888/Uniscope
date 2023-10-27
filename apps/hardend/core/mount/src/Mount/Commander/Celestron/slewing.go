package celestron

import (
	"errors"
)

type SlewParam uint8

const (
	SLEW_IDENT       = 0x50
	VARIABLE_RATE    = 0x03
	FIXED_RATE       = 0x02
	AZM_RA           = 0x10
	ALT_DEC          = 0x11
	POS_DIRECT_VAR   = 0x06
	NEG_DIRECT_VAR   = 0x07
	POS_DIRECT_FIXED = 0x24
	NEG_DIRECT_FIXED = 0x25
	TRACK_RATE_HIGH  = 0x08
	TRACK_RATE_LOW   = 0x01
)

// FixedRateSlew accepts byte between 0x00 and 0x08
func (m *MountCelestron) FixedRateSlew(coordinatesType SlewParam, isPosDirection bool, rate byte) error {
	if rate < 0 || rate > 9 {
		return errors.New("invalid rate: rate should be between 0 and 9")
	}
	if coordinatesType != AZM_RA && coordinatesType != ALT_DEC {
		return errors.New("invalid type of coordinates: type of coordinates should be AZM_RA or ALT_DEC")
	}
	direction := NEG_DIRECT_FIXED
	if isPosDirection {
		direction = POS_DIRECT_FIXED
	}

	_, err := m.Communicator().Write([]byte{SLEW_IDENT, FIXED_RATE, byte(coordinatesType), byte(direction), rate, 0x00, 0x00, 0x00})
	if err != nil {
		return err
	}

	if ok, err := m.Communicator().GetApproval(); !ok {
		return err
	}

	return nil
}

// VariableRateSlew takes arcseconds per second as speed
func (m *MountCelestron) VariableRateSlew(coordinatesType SlewParam, isPosDirection bool, speed int) error {

	if coordinatesType != AZM_RA && coordinatesType != ALT_DEC {
		return errors.New("invalid type of coordinates: type of coordinates should be AZM_RA or ALT_DEC")
	}
	direction := NEG_DIRECT_VAR
	if isPosDirection {
		direction = POS_DIRECT_VAR
	}

	trackRateHigh := (speed * 4) / 256
	trackRateLow := (speed * 4) % 256

	_, err := m.Communicator().Write([]byte{
		SLEW_IDENT,
		VARIABLE_RATE,
		byte(coordinatesType),
		byte(direction),
		byte(trackRateHigh),
		byte(trackRateLow),
		0x00, 0x00,
	})
	if err != nil {
		return err
	}

	if ok, err := m.Communicator().GetApproval(); !ok {
		return err
	}

	return nil
}
