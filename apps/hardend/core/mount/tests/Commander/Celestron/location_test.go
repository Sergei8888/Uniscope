package Celestron

//import (
//	"core/mount/configs"
//	"core/mount/src/Coordinates/Earth"
//	mount "core/mount/src/Mount"
//	celestron "core/mount/src/Mount/Commander/Celestron"
//	"github.com/stretchr/testify/assert"
//	"log"
//	"testing"
//)
//
//func TestLocation(t *testing.T) {
//	earthCoordinates := Earth.NewLocation()
//	err := earthCoordinates.SetByString("10 34 13 E 56 12 43 S")
//
//	assert.NoError(t, err)
//
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
//	err = mount.SetLocation(earthCoordinates)
//
//	assert.NoError(t, err)
//
//	got, err := mount.GetLocation()
//
//	if earthCoordinates.Latitude().Degrees() != got.Latitude().Degrees() {
//		t.Errorf("latitude, degrees: expected %d, got %d", earthCoordinates.Latitude().Degrees(), got.Latitude().Degrees())
//	}
//
//	if earthCoordinates.Latitude().Minutes() != got.Latitude().Minutes() {
//		t.Errorf("latitude, minutes: expected %d, got %d", earthCoordinates.Latitude().Minutes(), got.Latitude().Minutes())
//	}
//
//	if earthCoordinates.Latitude().Seconds() != got.Latitude().Seconds() {
//		t.Errorf("latitude, seconds: expected %d, got %d", earthCoordinates.Latitude().Seconds(), got.Latitude().Seconds())
//	}
//
//	if earthCoordinates.Latitude().SouthNorthPosition() != got.Latitude().SouthNorthPosition() {
//		t.Errorf("latitude, SN: expected %d, got %d", earthCoordinates.Latitude().SouthNorthPosition(), got.Latitude().SouthNorthPosition())
//	}
//
//	if earthCoordinates.Longitude().Degrees() != got.Longitude().Degrees() {
//		t.Errorf("longitude, degrees: expected %d, got %d", earthCoordinates.Longitude().Degrees(), got.Longitude().Degrees())
//	}
//
//	if earthCoordinates.Longitude().Minutes() != got.Longitude().Minutes() {
//		t.Errorf("longitude, minutes: expected %d, got %d", earthCoordinates.Longitude().Minutes(), got.Longitude().Minutes())
//	}
//
//	if earthCoordinates.Longitude().Seconds() != got.Longitude().Seconds() {
//		t.Errorf("longitude, seconds: expected %d, got %d", earthCoordinates.Longitude().Seconds(), got.Longitude().Seconds())
//	}
//
//	if earthCoordinates.Longitude().EastWestPosition() != got.Longitude().EastWestPosition() {
//		t.Errorf("latitude, EW: expected %d, got %d", earthCoordinates.Longitude().EastWestPosition(), got.Longitude().EastWestPosition())
//	}
//}
