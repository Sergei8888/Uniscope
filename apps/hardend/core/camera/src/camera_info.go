package src

//#cgo CFLAGS: -I../include
//#cgo LDFLAGS: -L../lib/x64 -lASICamera2 -Wl,-rpath=../lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
	"errors"
	"fmt"
)

// get src id stored in flash, only available for USB3.0 src
func (c *Camera) GetID() (string, error) {
	var pID C.struct__ASI_ID
	status := C.ASIGetID(C.int(c.cameraID), &pID)
	if int(status) != 0 {
		return "", errors.New(fmt.Sprintf("GetID func finished with status: %d", int(status)))
	}
	goID := make([]byte, 8)
	for i := 0; i < 8; i++ {
		goID[i] = byte(pID.id[i])
	}
	return string(goID), nil
}

// write src id to flash, only available for USB3.0 src
func (c *Camera) SetID(ID [8]byte) error {
	var cID [8]C.uchar
	for i := 0; i < 8; i++ {
		cID[i] = C.uchar(ID[i])
	}
	status := C.ASISetID(C.int(c.cameraID), C.struct__ASI_ID{cID})
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("GetID func finished with status: %d", int(status)))
	}
	return nil
}
