package Celestron

import (
	"bytes"
	"core/mount/configs"
	"core/mount/src/Coordinates/Sky"
	mount "core/mount/src/Mount"
	celestron "core/mount/src/Mount/Commander/Celestron"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"testing"
)

//func TestPositionEquatorialCoordinates(t *testing.T) {
//	port, err := mount.ConfigurePort(configs.PORT_NAME, configs.MODE)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	mount := celestron.NewMountCelestron(port)
//
//	if err = mount.PrepareAfterTurnOn(); err != nil {
//		log.Fatalln(err)
//	}
//
//	eqCS, err := mount.GetRaDec()
//	assert.NoError(t, err)
//
//	eqCS.Dec().SetDegrees(10)
//
//	err = mount.GotoRaDec(eqCS)
//	assert.NoError(t, err)
//
//	buff := bytes.Buffer{}
//	mount.GotoLock(&buff)
//	got, err := mount.GetRaDec()
//	assert.NoError(t, err)
//
//	if eqCS.Ra().Degrees() != got.Ra().Degrees() {
//		t.Errorf("ra, degrees: expected %d, got %d", eqCS.Ra().Degrees(), got.Ra().Degrees())
//	}
//
//	if eqCS.Ra().Minutes() != got.Ra().Minutes() {
//		t.Errorf("ra, minutes: expected %d, got %d", eqCS.Ra().Minutes(), got.Ra().Minutes())
//	}
//
//	if eqCS.Ra().Seconds() != got.Ra().Seconds() {
//		t.Errorf("ra, seconds: expected %d, got %d", eqCS.Ra().Seconds(), got.Ra().Seconds())
//	}
//
//	if eqCS.Dec().Degrees() != got.Dec().Degrees() {
//		t.Errorf("dec, degrees: expected %d, got %d", eqCS.Dec().Degrees(), got.Dec().Degrees())
//	}
//
//	if eqCS.Dec().Minutes() != got.Dec().Minutes() {
//		t.Errorf("dec, minutes: expected %d, got %d", eqCS.Dec().Minutes(), got.Dec().Minutes())
//	}
//
//	if eqCS.Dec().Seconds() != got.Dec().Seconds() {
//		t.Errorf("dec, seconds: expected %d, got %d", eqCS.Dec().Seconds(), got.Dec().Seconds())
//	}
//}

func TestPositionHorizontalCoordinates(t *testing.T) {
	port, err := mount.ConfigurePort(configs.PORT_NAME, configs.MODE)
	if err != nil {
		log.Fatalln(err)
	}

	mount := celestron.NewMountCelestron(port)

	if err = mount.PrepareAfterTurnOn(); err != nil {
		log.Fatalln(err)
	}

	azm, _ := Sky.NewAzimuth(0, 0, 0)
	alt, _ := Sky.NewAltitude(0, 0, 0)

	mount.GotoAzmAlt(Sky.NewHorizontalCS(azm, alt))

	azm, _ = Sky.NewAzimuth(10, 20, 20)
	alt, _ = Sky.NewAltitude(10, 20, 20)
	horizCS := Sky.NewHorizontalCS(azm, alt)

	err = mount.GotoAzmAlt(horizCS)
	assert.NoError(t, err)

	buff := bytes.Buffer{}
	mount.GotoLock(&buff)

	got, err := mount.GetAzmAlt()
	assert.NoError(t, err)

	if horizCS.Alt().Degrees() != got.Alt().Degrees() {
		t.Errorf("Alt, degrees: expected %d, got %d", horizCS.Alt().Degrees(), got.Alt().Degrees())
	}

	if horizCS.Alt().Minutes() != got.Alt().Minutes() {
		t.Errorf("Alt, minutes: expected %d, got %d", horizCS.Alt().Minutes(), got.Alt().Minutes())
	}

	if horizCS.Alt().Seconds() != got.Alt().Seconds() {
		if math.Abs(float64(horizCS.Alt().Seconds()-got.Alt().Seconds())) > 19.8 {
			t.Errorf("Alt, seconds: expected %d, got %d", horizCS.Alt().Seconds(), got.Alt().Seconds())
		}
	}

	if horizCS.Azm().Degrees() != got.Azm().Degrees() {
		t.Errorf("Azm, degrees: expected %d, got %d", horizCS.Azm().Degrees(), got.Azm().Degrees())
	}

	if horizCS.Azm().Minutes() != got.Azm().Minutes() {
		t.Errorf("Azm, minutes: expected %d, got %d", horizCS.Azm().Minutes(), got.Azm().Minutes())
	}

	if horizCS.Azm().Seconds() != got.Azm().Seconds() {
		if math.Abs(float64(horizCS.Azm().Seconds()-got.Azm().Seconds())) > 19.8 {
			t.Errorf("Azm, seconds: expected %d, got %d", horizCS.Azm().Seconds(), got.Azm().Seconds())
		}
	}
}

func TestPrecisePositionHorizontalCoordinates(t *testing.T) {
	port, err := mount.ConfigurePort(configs.PORT_NAME, configs.MODE)
	if err != nil {
		log.Fatalln(err)
	}

	mount := celestron.NewMountCelestron(port)

	if err = mount.PrepareAfterTurnOn(); err != nil {
		log.Fatalln(err)
	}

	azm, _ := Sky.NewAzimuth(0, 0, 0)
	alt, _ := Sky.NewAltitude(0, 0, 0)

	mount.GotoPreciseAzmAlt(Sky.NewHorizontalCS(azm, alt))

	azm, _ = Sky.NewAzimuth(15, 20, 20)
	alt, _ = Sky.NewAltitude(15, 20, 20)
	horizCS := Sky.NewHorizontalCS(azm, alt)

	err = mount.GotoPreciseAzmAlt(horizCS)
	assert.NoError(t, err)

	buff := bytes.Buffer{}
	mount.GotoLock(&buff)

	got, err := mount.GetPreciseAzmAlt()
	assert.NoError(t, err)

	if horizCS.Alt().Degrees() != got.Alt().Degrees() {
		t.Errorf("Alt, degrees: expected %d, got %d", horizCS.Alt().Degrees(), got.Alt().Degrees())
	}

	if horizCS.Alt().Minutes() != got.Alt().Minutes() {
		t.Errorf("Alt, minutes: expected %d, got %d", horizCS.Alt().Minutes(), got.Alt().Minutes())
	}

	if horizCS.Alt().Seconds() != got.Alt().Seconds() {
		if math.Abs(float64(horizCS.Alt().Seconds()-got.Alt().Seconds())) > 0.08 {
			t.Errorf("Alt, seconds: expected %d, got %d", horizCS.Alt().Seconds(), got.Alt().Seconds())
		}
	}

	if horizCS.Azm().Degrees() != got.Azm().Degrees() {
		t.Errorf("Azm, degrees: expected %d, got %d", horizCS.Azm().Degrees(), got.Azm().Degrees())
	}

	if horizCS.Azm().Minutes() != got.Azm().Minutes() {
		t.Errorf("Azm, minutes: expected %d, got %d", horizCS.Azm().Minutes(), got.Azm().Minutes())
	}

	if horizCS.Azm().Seconds() != got.Azm().Seconds() {
		if math.Abs(float64(horizCS.Azm().Seconds()-got.Azm().Seconds())) > 0.08 {
			t.Errorf("Azm, seconds: expected %d, got %d", horizCS.Azm().Seconds(), got.Azm().Seconds())
		}
	}
}
