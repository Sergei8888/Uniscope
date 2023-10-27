package celestron

import (
	"errors"
	"fmt"
)

const (
	TRACKING_OFF        = 0
	ALT_AZ_TRACKING     = 1
	EQUATORIAL_TRACKING = 2
	PEC_MODE            = 3
)

func (m *MountCelestron) GetTrackingMode() error {
	_, err := m.Communicator().Write([]byte{'t'})
	if err != nil {
		return err
	}

	answer, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return err
	}

	err = m.SetTrackingMode(int(answer[0]))
	if err != nil {
		return err
	}

	return nil
}

// SetTrackingMode send new mode to mounting.
// To use mode already set in mounting,
// set value of mode = -1
func (m *MountCelestron) SetTrackingMode(mode int) error {
	if mode < 0 || mode > 4 {
		return errors.New(
			fmt.Sprintf("tracking mode should be from 0 to 4, got: %d", mode),
		)
	}
	_, err := m.Communicator().Write([]byte{'T', byte(mode)})
	if err != nil {
		return err
	}

	if err = m.SetTrackingMode(mode); err != nil {
		return err
	}

	return nil
}
