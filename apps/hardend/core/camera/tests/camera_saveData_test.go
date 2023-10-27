package tests

import (
	"core/camera/constants"
	"core/camera/src"
	"core/camera/utils"
	"os"
	"testing"
)

func BenchmarkSaveRGB24(b *testing.B) {
	camInput, err := utils.ReadNumberOfCamera(os.Stdout, os.Stdin)

	cam, err := src.NewCamera(camInput)
	if err != nil {
		b.Errorf("src wasn't created: %v", err)
	}

	err = cam.FullInit(src.NewFullMatrixROI(cam, constants.ASI_IMG_RGB24))
	if err != nil {
		b.Errorf("src wasn't initialised: %v", err)
	}

	cam.SetControlCapValueByID(constants.ASI_EXPOSURE, 100, false)

	for i := 0; i < b.N; i++ {
		_, err = cam.MakeCapture(false, "../frames_rgb24_bench")
		if err != nil {
			b.Errorf("making shapshot stopped with status: %v", err)
		}
	}
	err = cam.Close()
	if err != nil {
		b.Errorf("closing src finished with status: %v", err)
	}
}
