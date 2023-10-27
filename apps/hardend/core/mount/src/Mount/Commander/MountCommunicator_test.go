package Commander

import (
	"bytes"
	celestron "core/mount/src/Mount/Commander/Celestron"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.bug.st/serial"
	"io"
	"testing"
	"time"
)

type TestPort struct {
	data bytes.Buffer
}

func (tp *TestPort) SetMode(mode *serial.Mode) error {
	return nil
}
func (tp *TestPort) Read(p []byte) (n int, err error) {
	n, err = tp.data.Read(p)
	if err == io.EOF {
		err = nil
	}
	return
}
func (tp *TestPort) Write(p []byte) (n int, err error) {
	return tp.data.Write(p)
}
func (tp *TestPort) ResetInputBuffer() error {
	tp.data.Reset()
	return nil
}
func (tp *TestPort) ResetOutputBuffer() error {
	tp.data.Reset()
	return nil
}
func (tp *TestPort) SetDTR(dtr bool) error {
	return nil
}
func (tp *TestPort) SetRTS(rts bool) error {
	return nil
}
func (tp *TestPort) GetModemStatusBits() (*serial.ModemStatusBits, error) {
	return &serial.ModemStatusBits{}, nil
}
func (tp *TestPort) Close() error {
	return nil
}

func (tp *TestPort) SetReadTimeout(t time.Duration) error {
	return nil
}

func (tp *TestPort) Break(time.Duration) error {
	return nil
}

func TestGetLocation(t *testing.T) {
	cases := []struct {
		location       []byte
		expectedOutput string
	}{
		{[]byte{33, 50, 41, 0, 118, 20, 17, 1, '#'}, "118°20’17” W, 33°50’41” N"},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			l := bytes.Buffer{}
			l.Write(tc.location)
			port := &TestPort{data: l}
			m := celestron.NewMountCelestron(port)
			got, err := m.GetLocation()

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, got)
		})
	}
}

func TestSetLocation(t *testing.T) {
	cases := []struct {
		location       []byte
		expectedOutput string
	}{
		{[]byte{33, 50, 41, 0, 118, 20, 17, 1, '#'}, "118°20’17” W, 33°50’41” N"},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			l := bytes.Buffer{}
			l.Write(tc.location)
			port := &TestPort{data: l}
			m := celestron.NewMountCelestron(port)
			got, err := m.GetLocation()

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, got)
		})
	}
}
