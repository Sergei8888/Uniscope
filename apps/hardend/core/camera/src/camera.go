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
	"strings"
)

type Camera struct {
	name                 string
	cameraID             int
	maxHeight            int64
	maxWidth             int64
	isColorCam           bool
	bayerPattern         constants.ASI_BAYER_PATTERN
	supportedBins        []int                    // 1 means bin1 which is supported by every src, 2 means bin 2 etc...
	supportedVideoFormat []constants.ASI_IMG_TYPE // this array will content with the support output
	pixelSize            float64                  // the pixel size of the src, unit is um. such like 5.6um
	controlCaps          []*ControlCap
	ROI                  *ROI
	mechanicalShutter    bool
	st4Port              bool
	isCoolerCam          bool
	isUSB3Host           bool
	isUSB3Camera         bool
	elecPerADU           bool
	bitDepth             int // the actual ADC depth of image sensor
	isTriggerCam         bool
	buffer               []byte
}

func NewCamera(iCameraIndex int) (*Camera, error) {
	var camInfo C.struct__ASI_CAMERA_INFO
	status := C.ASIGetCameraProperty(&camInfo, C.int(iCameraIndex))
	if status != 0 {
		return nil, errors.New(fmt.Sprintf("get src property status: %d", int(status)))
	}

	cSupportedBins := camInfo.SupportedBins[:]
	supportedBins := make([]int, 0)
	for _, bin := range cSupportedBins {
		if int(bin) == 0 {
			break
		}
		supportedBins = append(supportedBins, int(bin))
	}

	cSupportedVideoFormat := camInfo.SupportedVideoFormat[:]
	supportedVideoFormat := make([]constants.ASI_IMG_TYPE, 0)
	for _, vf := range cSupportedVideoFormat {
		if constants.ASI_IMG_TYPE(vf) == constants.ASI_IMG_END {
			break
		}
		supportedVideoFormat = append(supportedVideoFormat, constants.ASI_IMG_TYPE(vf))
	}

	cam := &Camera{
		name:                 C.GoString(&camInfo.Name[0]),
		cameraID:             int(camInfo.CameraID),
		maxHeight:            int64(camInfo.MaxHeight),
		maxWidth:             int64(camInfo.MaxWidth),
		isColorCam:           int(camInfo.IsColorCam) != 0,
		bayerPattern:         constants.ASI_BAYER_PATTERN(camInfo.BayerPattern),
		supportedBins:        supportedBins,
		supportedVideoFormat: supportedVideoFormat,
		pixelSize:            float64(camInfo.PixelSize),
		mechanicalShutter:    int(camInfo.MechanicalShutter) != 0,
		st4Port:              int(camInfo.ST4Port) != 0,
		isCoolerCam:          int(camInfo.IsCoolerCam) != 0,
		isUSB3Host:           int(camInfo.IsUSB3Host) != 0,
		isUSB3Camera:         int(camInfo.IsUSB3Camera) != 0,
		elecPerADU:           int(camInfo.ElecPerADU) != 0,
		bitDepth:             int(camInfo.BitDepth),
		isTriggerCam:         int(camInfo.IsTriggerCam) != 0,
		buffer:               []byte{},
	}
	return cam, nil
}

func (c *Camera) FullInit(roi *ROI) error {
	if err := c.Open(); err != nil {
		return err
	}

	if err := c.Init(); err != nil {
		c.Close()
		return err
	}

	if err := c.InitControlCaps(); err != nil {
		c.Close()
		return err
	}

	if err := c.SetROIFormat(roi); err != nil {
		c.Close()
		return err
	}

	if err := c.SetStartPosition(roi); err != nil {
		c.Close()
		return err
	}
	return nil
}

func (c *Camera) Info() string {
	return fmt.Sprintf(
		`Name: %s
Camera ID: %d
Max Height: %d
Max Width: %d
Is Color Cam: %t
Bayer Pattern: %s
supportedBins: %v
supportedVideoFormat: %v
Pixel Size: %f
Control Caps: %v
Mechanical Shutter: %t
ST4Port: %t
Is Cooler Cam: %t
Is USB3 Host: %t
Is USB3 Camera: %t
Elec Per ADU: %t
Bit Depth: %d
Is Trigger Cam: %t
`,
		c.name, c.cameraID,
		c.maxHeight, c.maxWidth,
		c.isColorCam, c.bayerPattern.ToString(),
		c.supportedBins, c.supportedVideoFormat,
		c.pixelSize, strings.Join(c.GetListOfControlCaps(), ", "),
		c.mechanicalShutter, c.st4Port,
		c.isCoolerCam, c.isUSB3Host,
		c.isUSB3Camera, c.elecPerADU,
		c.bitDepth, c.isTriggerCam,
	)
}

func (c *Camera) Open() error {
	status := C.ASIOpenCamera(C.int(c.cameraID))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("'open' operation finished with status: %d", int(status)))
	}
	return nil
}

func (c *Camera) Init() error {
	status := C.ASIInitCamera(C.int(c.cameraID))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("'init' operation finished with status: %d", int(status)))
	}
	return nil
}

func (c *Camera) initBuffer() (buffSize int) {
	typeMultiplier := 1
	switch c.ROI.ImageType() {
	case constants.ASI_IMG_RAW8:
		typeMultiplier = 1
		break
	case constants.ASI_IMG_RAW16:
		typeMultiplier = 2
		break
	case constants.ASI_IMG_RGB24:
		typeMultiplier = 3
		break
	}
	buffSize = typeMultiplier * (int(c.ROI.Width()) - c.ROI.XStart()) * (int(c.ROI.Height()) - c.ROI.YStart())
	c.buffer = make([]byte, buffSize)
	return buffSize
}

func (c *Camera) Close() error {
	status := C.ASICloseCamera(C.int(c.cameraID))
	if int(status) != 0 {
		return errors.New(fmt.Sprintf("'close' operation finished with status: %d", int(status)))
	}
	return nil
}

func (c *Camera) Name() string {
	return c.name
}

func (c *Camera) CameraID() int {
	return c.cameraID
}

func (c *Camera) MaxHeight() int64 {
	return c.maxHeight
}

func (c *Camera) MaxWidth() int64 {
	return c.maxWidth
}

func (c *Camera) IsColorCam() bool {
	return c.isColorCam
}

func (c *Camera) BayerPattern() constants.ASI_BAYER_PATTERN {
	return c.bayerPattern
}

func (c *Camera) SupportedBins() []int {
	return c.supportedBins
}

func (c *Camera) SupportedVideoFormat() []constants.ASI_IMG_TYPE {
	return c.supportedVideoFormat
}

func (c *Camera) PixelSize() float64 {
	return c.pixelSize
}

func (c *Camera) ControlCaps() []*ControlCap {
	return c.controlCaps
}

func (c *Camera) MechanicalShutter() bool {
	return c.mechanicalShutter
}

func (c *Camera) St4Port() bool {
	return c.st4Port
}

func (c *Camera) IsCoolerCam() bool {
	return c.isCoolerCam
}

func (c *Camera) IsUSB3Host() bool {
	return c.isUSB3Host
}

func (c *Camera) IsUSB3Camera() bool {
	return c.isUSB3Camera
}

func (c *Camera) ElecPerADU() bool {
	return c.elecPerADU
}

func (c *Camera) BitDepth() int {
	return c.bitDepth
}

func (c *Camera) IsTriggerCam() bool {
	return c.isTriggerCam
}
