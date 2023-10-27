package src

import (
	"core/camera/constants"
)

type ROI struct { // Region of interest
	width   int64
	height  int64
	xStart  int
	yStart  int
	bin     int
	imgType constants.ASI_IMG_TYPE
}

func NewROI(width, height int64, xStart, yStart, bin int, imgType constants.ASI_IMG_TYPE) *ROI {
	return &ROI{
		width:   width,
		height:  height,
		xStart:  xStart,
		yStart:  yStart,
		bin:     bin,
		imgType: imgType,
	}
}

func NewFullMatrixROI(cam *Camera, imgType constants.ASI_IMG_TYPE) *ROI {
	return NewROI(cam.MaxWidth(), cam.MaxHeight(), 0, 0, 1, imgType)
}

func (R ROI) XStart() int {
	return R.xStart
}

func (R ROI) YStart() int {
	return R.yStart
}

func (R ROI) Width() int64 {
	return R.width
}

func (R ROI) Height() int64 {
	return R.height
}

func (R ROI) Bin() int {
	return R.bin
}

func (R ROI) ImageType() constants.ASI_IMG_TYPE {
	return R.imgType
}
