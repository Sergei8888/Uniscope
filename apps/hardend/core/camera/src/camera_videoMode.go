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
	"time"
)

func (c *Camera) VideoCaptureByTime(t time.Duration, destDirectory string) (framesTotal int, dropped int, err error) {
	if err = c.StartVideoCapture(); err != nil {
		return 0, 0, err
	}

	expTime, _, err := c.GetControlCapValueByID(constants.ASI_EXPOSURE)
	if err != nil {
		return 0, 0, err
	}

	execEnd := false

	tmr := time.AfterFunc(t, func() {
		execEnd = true
	})
	defer tmr.Stop()

	for !execEnd {
		err = c.GetVideoData(int(expTime)*2 + 500)
		if err != nil {
			return
		}
		_, err = c.SaveFrame(destDirectory)
		if err != nil {
			return
		}
		framesTotal++
	}

	dropped, err = c.GetDroppedFramesCount()
	if err != nil {
		return
	}

	err = c.StopVideoCapture()
	if err != nil {
		return
	}
	return
}

func (c *Camera) StartVideoCapture() error {
	status := C.ASIStartVideoCapture(C.int(c.cameraID))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("start video capture finished with status: %d", int(status)))
	}
	return nil
}

func (c *Camera) StopVideoCapture() error {
	status := C.ASIStopVideoCapture(C.int(c.cameraID))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("stop video capture finished with status: %d", int(status)))
	}
	return nil
}

// iWaitMs value: exposure_time*2 + 500ms
func (c *Camera) GetVideoData(iWaitMs int) error {

	buffSize := c.initBuffer()
	status := C.ASIGetVideoData(C.int(c.cameraID), (*C.uchar)(&c.buffer[0]), C.long(buffSize), C.int(iWaitMs))
	if status != 0 {
		return errors.New(fmt.Sprintf("getting video data failed with status: %d", int(status)))
	}

	return nil
}

func (c *Camera) GetDroppedFramesCount() (n int, err error) {

	var piDropFrames C.int
	status := C.ASIGetDroppedFrames(C.int(c.cameraID), &piDropFrames)
	if status != 0 {
		return 0, errors.New(fmt.Sprintf("getting dropped frames number failed with status: %d", int(status)))
	}
	return int(piDropFrames), nil
}

// custom
func (c *Camera) AwaitVideoCaptureStop() {

}
