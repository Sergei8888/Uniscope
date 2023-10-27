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

func (c *Camera) GetCameraSupportMode() ([]constants.ASI_CAMERA_MODE, error) {
	var supportedModes C.struct__ASI_SUPPORTED_MODE
	status := C.ASIGetCameraSupportMode(C.int(c.cameraID), &supportedModes)
	if status != 0 {
		return nil, errors.New(fmt.Sprintf("getting src's support mode failed with status: %d", int(status)))
	}
	goSupportedModes := make([]constants.ASI_CAMERA_MODE, 0)
	for _, mode := range supportedModes.SupportedCameraMode {
		if mode == constants.ASI_MODE_END {
			break
		}
		goSupportedModes = append(goSupportedModes, constants.ASI_CAMERA_MODE(mode))
	}
	return goSupportedModes, nil
}

func (c *Camera) GetCameraMode() (constants.ASI_CAMERA_MODE, error) {
	var mode C.ASI_CAMERA_MODE
	status := C.ASIGetCameraMode(C.int(c.cameraID), &mode)
	if status != 0 {
		return 0, errors.New(fmt.Sprintf("getting src's support mode failed with status: %d", int(status)))
	}
	goMode := constants.ASI_CAMERA_MODE(mode)
	return goMode, nil
}

func (c *Camera) SetCameraMode(mode constants.ASI_CAMERA_MODE) error {
	status := C.ASISetCameraMode(C.int(c.cameraID), C.ASI_CAMERA_MODE(mode))
	if status != 0 {
		return errors.New(fmt.Sprintf("getting src's support mode failed with status: %d", int(status)))
	}
	return nil
}
