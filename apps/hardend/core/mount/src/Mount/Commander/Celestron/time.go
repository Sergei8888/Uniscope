package celestron

import (
	"fmt"
	"time"
)

func (m *MountCelestron) GetTime() (string, error) {
	_, err := m.Communicator().Write([]byte{'h'})
	if err != nil {
		return "", err
	}
	a, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d:%d:%d %d/%d/%d GMT %d Daylight saving time enabled: %t",
		a[1], a[2], a[3], a[5], a[4], a[6], a[7], a[8] == '1'), nil
}

func (m *MountCelestron) SetTime(datetime time.Time) error {

	_, zoneOffset := datetime.Zone()
	if zoneOffset < 0 {
		zoneOffset = 256 - zoneOffset
	}

	zoneOffset /= 3600
	daylightSaving := 0
	if datetime.IsDST() {
		daylightSaving = 1
	}

	_, err := m.Communicator().Write([]byte{
		'H',
		byte(datetime.Hour()),
		byte(datetime.Minute()),
		byte(datetime.Second()),
		byte(datetime.Month()),
		byte(datetime.Day()),
		byte(datetime.Year() % 100),
		byte(zoneOffset),
		byte(daylightSaving),
	})

	if err != nil {
		return err
	}
	if ok, err := m.Communicator().GetApproval(); !ok {
		return err
	}
	return nil
}
