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

func (c *Camera) MakeCapture(isDark bool, destDirectory string) (filename string, err error) {
	ok, err := c.ExecuteExposure(isDark)
	if !ok {
		return "", err
	}
	if err = c.GetDataAfterExp(); err != nil {
		return "", err
	}

	if filename, err = c.SaveFrame(destDirectory); err != nil {
		return "", err
	}
	return
}

func (c *Camera) ExecuteExposure(isDark bool) (isSucceed bool, err error) {
	var isDarkInt int
	if isDark {
		isDarkInt = 1
	}

	status := C.ASIStartExposure(C.int(c.cameraID), C.int(isDarkInt))
	if status != 0 {
		return false, errors.New(fmt.Sprintf("start of exposure failed with status: %d", int(status)))
	}
	var expStatus C.ASI_EXPOSURE_STATUS
	for int(expStatus) < constants.ASI_EXP_SUCCESS {
		status = C.ASIGetExpStatus(C.int(c.cameraID), &expStatus)
	}

	if expStatus == constants.ASI_EXP_FAILED {
		return false, errors.New("exposure failed")
	}
	return true, nil
}

func (c *Camera) GetDataAfterExp() error {

	buffSize := c.initBuffer()
	status := C.ASIGetDataAfterExp(C.int(c.cameraID), (*C.uchar)(&c.buffer[0]), C.long(buffSize))
	if status != 0 {
		return errors.New(fmt.Sprintf("getting data after exposure failed with status: %d", int(status)))
	}
	return nil
}

func (c *Camera) StopExposure() error {
	status := C.ASIStopExposure(C.int(c.cameraID))
	if status != 0 {
		return errors.New(fmt.Sprintf("getting data after exposure failed with status: %d", int(status)))
	}
	return nil
}
