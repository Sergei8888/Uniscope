package celestron

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (m *MountCelestron) CancelGOTO() error {
	_, err := m.Communicator().Write([]byte{'M'})
	if err != nil {
		return err
	}
	if ok, err := m.Communicator().GetApproval(); !ok {
		return err
	}
	return nil
}

func (m *MountCelestron) Echo(echoByte byte) (string, error) {
	_, err := m.Communicator().Write([]byte{'K', echoByte})
	if err != nil {
		return "", err
	}
	echo, err := m.Communicator().ReadAnswerString()
	if err != nil {
		return "", err
	}
	return echo[1:], nil
}

func (m *MountCelestron) GetDeviceVersion(isHorizontalMotor bool) (string, error) {
	/*
		“P” &
		chr(1) &
		chr(dev) &
		chr(254) &
		chr(0) &
		chr(0) &
		chr(0) &
		chr(2)
	*/
	dev := 16
	if !isHorizontalMotor {
		dev++
	}
	_, err := m.Communicator().Write([]byte{'P', 1, byte(dev), 254, 0, 0, 0, 2})
	if err != nil {
		return "", err
	}

	ver, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return "", err
	}
	ver = ver[1:]
	verStr := make([]string, 0)
	for _, b := range ver {
		verStr = append(verStr, strconv.Itoa(int(b)))
	}

	return strings.Join(verStr, "."), nil
}

func (m *MountCelestron) GetModel() (string, error) {
	_, err := m.Communicator().Write([]byte{'m'})
	if err != nil {
		return "", err
	}
	//TODO: read bytes and convert to model
	echo, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return "", err
	}
	if len(echo) == 0 {
		return "", errors.New("got no model")
	}

	switch echo[0] {
	case 0x00:
		return "EQ6 GOTO Series", nil
	case 0x01:
		return "HEQ5 GOTO Series", nil
	case 0x02:
		return "EQ5 GOTO Series", nil
	case 0x03:
		return "EQ3 GOTO Series", nil
	case 0x04:
		return "EQ8 GOTO Series", nil
	case 0x05:
		return "AZ-EQ6 GOTO Series", nil
	case 0x06:
		return "AZ-EQ5 GOTO Series", nil
	}

	return "bad response", nil
}

// GetMountPointState doesn't work
func (m *MountCelestron) GetMountPointState() (string, error) {
	_, err := m.Communicator().Write([]byte{'p'})
	if err != nil {
		return "", err
	}
	fmt.Println(m.Communicator().ReadAnswerBytes())

	return "", err
}

func (m *MountCelestron) GetVersion() (string, error) {
	if _, err := m.Communicator().Write([]byte{'V'}); err != nil {
		return "", err
	}
	answer, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return "", err
	}
	answer = answer[1:]

	version := ""
	for _, b := range answer {
		version += strconv.Itoa(int(b)) + "."
	}

	return strings.Trim(version, "."), nil
}

func (m *MountCelestron) IsAlignmentComplete() (bool, error) {
	_, err := m.Communicator().Write([]byte{'L'})
	if err != nil {
		return false, err
	}
	answer, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return false, err
	}
	return answer[1] == '1', nil
}

func (m *MountCelestron) IsGotoInProgress() (bool, error) {
	_, err := m.Communicator().Write([]byte{'L'})
	if err != nil {
		return false, err
	}
	answer, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return false, err
	}

	return answer[1] == 49, nil
}
