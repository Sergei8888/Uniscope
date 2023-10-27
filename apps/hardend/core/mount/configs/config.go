package configs

import "go.bug.st/serial"

type Config struct {
	PortName string
}

const (
	PORT_NAME = "/dev/ttyUSB0"
)

var (
	MODE = &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
)
