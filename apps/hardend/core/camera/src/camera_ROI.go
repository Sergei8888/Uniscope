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

func (c *Camera) SetROIFormat(roi *ROI) error {
	if roi.width < 0 || roi.width > c.maxWidth {
		return errors.New(fmt.Sprintf("width can be between 0 and %d, got: %d", c.maxWidth, roi.width))
	}
	if roi.height < 0 || roi.height > c.maxHeight {
		return errors.New(fmt.Sprintf("height can be between 0 and %d, got: %d", c.maxHeight, roi.height))
	}
	var isValidBin bool
	for _, bin := range c.supportedBins {
		if bin == roi.bin {
			isValidBin = true
			break
		}
	}

	if !isValidBin {
		return errors.New(fmt.Sprintf("Bin %d isn't valid", roi.bin))
	}

	if roi.imgType < constants.ASI_IMG_RAW8 || roi.imgType > constants.ASI_IMG_Y8 {
		return errors.New(fmt.Sprintf("Image type %d isn't valid", roi.imgType))
	}

	status := C.ASISetROIFormat(C.int(c.cameraID), C.int(roi.width), C.int(roi.height), C.int(roi.bin), C.int(roi.imgType))
	if status != 0 {
		return errors.New(fmt.Sprintf("setting of ROI format finished with status: %d", int(status)))
	}
	c.ROI = roi

	return nil
}

func (c *Camera) GetROIFormat() (*ROI, error) {
	var piWidth C.int
	var piHeight C.int
	var piBin C.int
	var pImgType C.int
	status := C.ASIGetROIFormat(C.int(c.cameraID), &piWidth, &piHeight, &piBin, &pImgType)
	if status != 0 {
		return nil, errors.New(fmt.Sprintf("setting of ROI format finished with status: %d", int(status)))
	}
	c.ROI.width = int64(piWidth)
	c.ROI.height = int64(piHeight)
	c.ROI.bin = int(piBin)
	c.ROI.imgType = constants.ASI_IMG_TYPE(pImgType)

	return c.ROI, nil
}

func (c *Camera) SetStartPosition(roi *ROI) error {
	status := C.ASISetStartPos(C.int(c.cameraID), C.int(roi.xStart), C.int(roi.yStart))
	if status != 0 {
		return errors.New(fmt.Sprintf("setting of start position finished with status: %d", int(status)))
	}
	return nil
}

func (c *Camera) GetStartPosition(roi *ROI) error {
	var piStartX C.int
	var piStartY C.int
	status := C.ASIGetStartPos(C.int(c.cameraID), &piStartX, &piStartY)
	if status != 0 {
		return errors.New(fmt.Sprintf("getting of start position finished with status: %d", int(status)))
	}
	c.ROI.xStart = int(piStartX)
	c.ROI.yStart = int(piStartY)
	return nil
}
