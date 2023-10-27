package mount

import (
	"errors"
	"fmt"
	"go.bug.st/serial"
)

func ConfigurePort(portName string, portMode *serial.Mode) (serial.Port, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}

	var isPortValid bool

	for _, port := range ports {
		if port == portName {
			isPortValid = true
			break
		}
	}

	if !isPortValid {
		return nil, errors.New(fmt.Sprintf("port wasn't found, available ports: %v", ports))
	}

	port, err := serial.Open(portName, portMode)
	if err != nil {
		return nil, err
	}
	err = port.SetMode(portMode)
	if err != nil {
		return nil, err
	}

	return port, nil
}
