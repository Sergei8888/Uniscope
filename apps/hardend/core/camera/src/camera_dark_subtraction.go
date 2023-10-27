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

// TODO: test
func (c *Camera) EnableDarkSubtract(pcBMPPath string) error {
	status := C.ASIEnableDarkSubtract(C.int(c.cameraID), C.CString(pcBMPPath))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("dark subtract enable finished with status: %d", int(status)))
	}
	return nil
}

// TODO: test
func (c *Camera) DisableDarkSubtract() error {
	status := C.ASIDisableDarkSubtract(C.int(c.cameraID))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("dark subtract disable finished with status: %d", int(status)))
	}
	return nil
}
