package Commander

import (
	"errors"
	"fmt"
	"io"
	"time"
)

type MountCommunicator struct {
	dataBus io.ReadWriter
}

func NewMountCommunicator(wr io.ReadWriter) *MountCommunicator {
	return &MountCommunicator{dataBus: wr}
}

//func (m *MountCommunicator) Port() serial.Port {
//	return m.dataBus
//}

//func (m *Mount) TrackingMode() int {
//	return m.trackingMode
//}
//
//func (m *Mount) SetTrackingMode(trackingMode int) error {
//	if trackingMode < 0 || trackingMode > 4 {
//		return errors.New(
//			fmt.Sprintf("tracking mode should be from 0 to 4, got: %d", trackingMode),
//		)
//	}
//	m.trackingMode = trackingMode
//	return nil
//}

//func (m *Mount) TrackingModeString() string {
//	switch m.trackingMode {
//	case 0:
//		return "Tracking off"
//	case 1:
//		return "Alt/Az tracking"
//	case 2:
//		return "Equatorial tracking"
//	case 3:
//		return "PEC mode (Sidereal + PEC)"
//	}
//	return "Undefined mode"
//}

//func (m *Mount) SetPort() error {
//	port, err := serial.Open(configs.PORT_NAME, mount.MODE)
//	if err != nil {
//		return err
//	}
//
//	err = port.SetMode(mount.MODE)
//	if err != nil {
//		return err
//	}
//
//	m.port = port
//	return nil
//}

func (m *MountCommunicator) Write(msg []byte) (n int, err error) {
	return m.dataBus.Write(msg)
}

func (m *MountCommunicator) ReadAnswerBytes() ([]byte, error) {
	buff := make([]byte, 64)

	var answer []byte
	var err error
	n := 1

	for true {
		if n == 0 {
			break
		}
		if buff[n-1] == '#' {
			answer = append(answer, buff[:n-1]...)
			break
		}
		answer = append(answer, buff[:n]...)
		n, err = m.dataBus.Read(buff)
		if err != nil {
			return []byte{}, err
		}
	}
	return answer, nil
}

func (m *MountCommunicator) ReadAnswerString() (string, error) {
	answerBytes, err := m.ReadAnswerBytes()
	return string(answerBytes), err
}

func (m *MountCommunicator) GetApproval() (bool, error) {
	buff := make([]byte, 4)
	n, err := m.dataBus.Read(buff)
	if err != nil {
		return false, err
	}
	if n != 1 {
		return false, errors.New(fmt.Sprintf("invalid number of read characters: %d", n))
	}
	if buff[0] != '#' {
		return false, errors.New(fmt.Sprintf("read invalid characters: %b", buff[0]))
	}
	return true, nil
}

func (m *MountCommunicator) TurnOn() ([]byte, error) {
	buff := make([]byte, 128)

	p := make(chan []byte)

	go func() error {
		if _, err := m.dataBus.Read(buff); err != nil {
			return err
		}
		p <- buff
		close(p)
		return nil
	}()

	select {
	case msg := <-p:
		return msg, nil
	case <-time.After(3 * time.Second):
		return nil, nil
	}
}
