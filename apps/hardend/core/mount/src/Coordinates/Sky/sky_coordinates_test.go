package Sky

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAltitude_ToHex(t *testing.T) {
	cases := []struct {
		alt      *Altitude
		expected []byte
	}{
		{&Altitude{0, 0, 0}, []byte{48, 48, 48, 48}},
		{&Altitude{180, 0, 0}, []byte{56, 48, 48, 48}},
		//TODO: make tests for seconds and minutes
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			got, err := tc.alt.ToHex()
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestAltitude_ToHexPrecise(t *testing.T) {
	cases := []struct {
		alt      *Altitude
		expected []byte
	}{
		{&Altitude{0, 0, 0}, []byte{'0', '0', '0', '0', '0', '0', '0', '0'}},
		{&Altitude{180, 0, 0}, []byte{56, 48, 48, 48, 48, 48, 48, 48}},
		//{&Altitude{4, 8, 4}, []byte{48, 50, 70, 48, 65, 57, 48, 48}},
		//{&Altitude{9, 52, 57}, []byte{48, 55, 48, 55, 49, 49, 48, 48}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			got, err := tc.alt.ToHexPrecise()
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}
