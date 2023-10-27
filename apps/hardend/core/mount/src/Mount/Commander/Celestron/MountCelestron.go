package celestron

import (
	"core/mount/src/Coordinates/Earth"
	"core/mount/src/Coordinates/Sky"
	"core/mount/src/Mount/Commander"
	"errors"
	"io"
	"time"
)

type MountCelestron struct {
	communicator     *Commander.MountCommunicator
	earthCoordinates *Earth.Location
	skyCoordinates   Sky.SkyCoordinates
	trackingMode     int
}

func NewMountCelestron(dataBus io.ReadWriter) *MountCelestron {
	return &MountCelestron{
		communicator:     Commander.NewMountCommunicator(dataBus),
		earthCoordinates: Earth.NewLocation(),
	}
}

func (m *MountCelestron) PrepareAfterTurnOn() error {

	time.Sleep(700 * time.Millisecond)

	res := make(chan error, 1)
	timeout := make(chan time.Time, 1)

	time.AfterFunc(5*time.Second, func() {
		close(timeout)
	})

	go func() {
		res <- func() error {
			echo, err := m.Echo('h')
			if err != nil {
				return err
			}

			for echo != "h" {
				echo, err = m.Echo('h')
				if err != nil {
					return err
				}
			}
			return nil
		}()
	}()

	var err error

	select {
	case <-timeout:
		return errors.New("timeout")
	case err = <-res:
		return err
	}
}

func (m *MountCelestron) Communicator() *Commander.MountCommunicator {
	return m.communicator
}
