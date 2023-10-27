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

func (c *Camera) SendSoftTrigger(bStart bool) error {
	var cbStart C.int
	if bStart {
		cbStart = 1
	}
	status := C.ASISendSoftTrigger(C.int(c.cameraID), cbStart)
	if status != 0 {
		return errors.New(fmt.Sprintf("sending soft trigger failed with status: %d", int(status)))
	}
	return nil
}
