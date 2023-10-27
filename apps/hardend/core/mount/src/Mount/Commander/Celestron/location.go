package celestron

import (
	"core/mount/src/Coordinates/Earth"
	"errors"
	"fmt"
)

func (m *MountCelestron) GetLocation() (*Earth.Location, error) {
	if _, err := m.Communicator().Write([]byte{'w'}); err != nil {
		return nil, err
	}
	answer, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return nil, err
	}
	answer = answer[1:]

	//118°20’17” W, 33°50’41” N

	latitude, err := Earth.NewLatitude(int(answer[0]), int(answer[1]), int(answer[2]), answer[3] == 0x01)
	if err != nil {
		return nil, err
	}
	longitude, err := Earth.NewLongitude(int(answer[4]), int(answer[5]), int(answer[6]), answer[7] == 0x01)
	if err != nil {
		return nil, err
	}

	earthLoc := Earth.NewLocation()

	earthLoc.SetLatitude(latitude)
	earthLoc.SetLongitude(longitude)

	return earthLoc, nil
}

// SetLocation sends location to mounting,
// if location == nil sends location already set
// in mounting
func (m *MountCelestron) SetLocation(location *Earth.Location) error {

	if location == nil {
		return errors.New("location: nil pointer")
	}

	m.earthCoordinates = location

	fmt.Println("here", []byte{
		'W',
		m.earthCoordinates.Latitude().Degrees(),
		m.earthCoordinates.Latitude().Minutes(),
		m.earthCoordinates.Latitude().Seconds(),
		m.earthCoordinates.Latitude().SouthNorthPosition(),
		m.earthCoordinates.Longitude().Degrees(),
		m.earthCoordinates.Longitude().Minutes(),
		m.earthCoordinates.Longitude().Seconds(),
		m.earthCoordinates.Longitude().EastWestPosition(),
	})

	_, err := m.Communicator().Write([]byte{
		'W',
		m.earthCoordinates.Latitude().Degrees(),
		m.earthCoordinates.Latitude().Minutes(),
		m.earthCoordinates.Latitude().Seconds(),
		m.earthCoordinates.Latitude().SouthNorthPosition(),
		m.earthCoordinates.Longitude().Degrees(),
		m.earthCoordinates.Longitude().Minutes(),
		m.earthCoordinates.Longitude().Seconds(),
		m.earthCoordinates.Longitude().EastWestPosition(),
	})
	if err != nil {
		return err
	}

	if ok, err := m.Communicator().GetApproval(); !ok {
		return err
	}
	m.earthCoordinates = location
	return nil
}
