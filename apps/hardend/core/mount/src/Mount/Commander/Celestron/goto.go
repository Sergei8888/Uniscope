package celestron

import (
	"core/mount/src/Coordinates/Sky"
	"fmt"
	"io"
	"log"
	"time"
)

// GotoRaDec accepts hexadecimal numbers
// Example: if RA = AB36, argument for function
// is []byte{'A', 'B', '3', '6'}
func (m *MountCelestron) GotoRaDec(eqCS *Sky.EquatorialCS) error {
	raHex, err := eqCS.Ra().ToHex()
	if err != nil {
		return err
	}
	decHex, err := eqCS.Dec().ToHex()
	if err != nil {
		return err
	}
	return m.gotoCoordinates(raHex, decHex, 'R')
}

// GotoPreciseRaDec accepts hexadecimal numbers
// Example: if RA = AB365F, argument for function
// is []byte{'A', 'B', '3', '6', '5', 'F'}
func (m *MountCelestron) GotoPreciseRaDec(eqCS *Sky.EquatorialCS) error {
	raHex, err := eqCS.Ra().ToHexPrecise()
	if err != nil {
		return err
	}
	decHex, err := eqCS.Dec().ToHexPrecise()
	if err != nil {
		return err
	}

	return m.gotoCoordinates(raHex, decHex, 'r')
}

// GotoAzmAlt accepts hexadecimal numbers
// Example: if AZM = AB36, argument for function
// is []byte{'A', 'B', '3', '6'}
func (m *MountCelestron) GotoAzmAlt(horizCS *Sky.HorizontalCS) error {

	azmHex, err := horizCS.Azm().ToHex()
	if err != nil {
		return err
	}
	altHex, err := horizCS.Alt().ToHex()
	if err != nil {
		return err
	}

	return m.gotoCoordinates(azmHex, altHex, 'B')
}

// GotoPreciseAzmAlt accepts hexadecimal numbers
// Example: if AZM = AB365F, argument for function
// is []byte{'A', 'B', '3', '6', '5', 'F'}
func (m *MountCelestron) GotoPreciseAzmAlt(horizCS *Sky.HorizontalCS) error {
	azmHex, err := horizCS.Azm().ToHexPrecise()
	if err != nil {
		return err
	}
	altHex, err := horizCS.Alt().ToHexPrecise()
	if err != nil {
		return err
	}

	return m.gotoCoordinates(azmHex, altHex, 'b')
}

func (m *MountCelestron) gotoCoordinates(firstItem, secondItem []byte, _type byte) error {

	var msg []byte
	msg = append(msg, _type)
	msg = append(msg, firstItem...)
	msg = append(msg, ',')
	msg = append(msg, secondItem...)
	_, err := m.Communicator().Write(msg)
	if err != nil {
		return err
	}

	if ok, err := m.Communicator().GetApproval(); !ok {
		return err
	}

	return nil
}

func (m *MountCelestron) GotoLock(w io.Writer) {
	go func() {
		for true {
			fmt.Fprintln(w, "Goto is in progress...")
			time.Sleep(3 * time.Second)
		}
	}()
	for true {
		moving, err := m.IsGotoInProgress()
		if err != nil {
			log.Println(err)
			break
		}
		if !moving {
			break
		}
	}
}
