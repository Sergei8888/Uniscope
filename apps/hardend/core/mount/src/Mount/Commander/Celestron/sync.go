package celestron

import (
	Sky "core/mount/src/Coordinates/Sky"
)

func (m *MountCelestron) SyncRaDec(ra *Sky.RightAscension, dec *Sky.Declination) error {

	raHex, err := ra.ToHex()
	if err != nil {
		return err
	}

	decHex, err := dec.ToHex()
	if err != nil {
		return err
	}
	msg := []byte{'S'}
	msg = append(msg, raHex...)
	msg = append(msg, ',')
	msg = append(msg, decHex...)
	_, err = m.Communicator().Write(msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *MountCelestron) SyncPreciseRaDec(ra *Sky.RightAscension, dec *Sky.Declination) error {

	raHex, err := ra.ToHexPrecise()
	if err != nil {
		return err
	}

	decHex, err := dec.ToHexPrecise()
	if err != nil {
		return err
	}
	msg := []byte{'s'}
	msg = append(msg, raHex...)
	msg = append(msg, ',')
	msg = append(msg, decHex...)
	_, err = m.Communicator().Write(msg)
	if err != nil {
		return err
	}
	return nil
}
