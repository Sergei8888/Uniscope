package src

//#cgo CFLAGS: -I../include
//#cgo LDFLAGS: -L../lib/x64 -lASICamera2 -Wl,-rpath=../lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
	"core/camera/constants"
	"errors"
	"fmt"
)

// TODO: test
func (c *Camera) PulseGuideOn(direction constants.ASI_GUIDE_DIRECTION) error {
	if !c.St4Port() {
		return errors.New("src doesn't support ST4 port")
	}

	if direction < 0 || direction > 4 {
		return errors.New("invalid direction")
	}
	status := C.ASIPulseGuideOn(C.int(c.cameraID), C.int(direction))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("PulseGuideOn finished with status: %d", int(status)))
	}

	return nil
}

// TODO: test
func (c *Camera) PulseGuideOff(direction constants.ASI_GUIDE_DIRECTION) error {
	if !c.St4Port() {
		return errors.New("src doesn't support ST4 port")
	}

	if direction < 0 || direction > 4 {
		return errors.New("invalid direction")
	}
	status := C.ASIPulseGuideOff(C.int(c.cameraID), C.int(direction))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("PulseGuideOff finished with status: %d", int(status)))
	}

	return nil
}
