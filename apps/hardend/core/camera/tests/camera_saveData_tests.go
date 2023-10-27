package tests

import (
	"core/camera/constants"
	"core/camera/src"
	"core/camera/utils"
	"fmt"
	"io"
	"os"
	"time"
)

func SaveRGB24Benchmark(w io.Writer, frames int) {
	saveBenchmark(w, frames, constants.ASI_IMG_RGB24)
}

func SaveRAW16Benchmark(w io.Writer, frames int) {
	saveBenchmark(w, frames, constants.ASI_IMG_RAW16)
}

func SaveRAW8Benchmark(w io.Writer, frames int) {
	saveBenchmark(w, frames, constants.ASI_IMG_RAW8)
}

func saveBenchmark(w io.Writer, frames int, imgType constants.ASI_IMG_TYPE) {
	camInput, err := utils.ReadNumberOfCamera(os.Stdout, os.Stdin)

	cam, err := src.NewCamera(camInput)
	if err != nil {
		fmt.Errorf("src wasn't created: %v", err)
	}

	err = cam.FullInit(src.NewFullMatrixROI(cam, imgType))
	if err != nil {
		fmt.Errorf("src wasn't initialised: %v", err)
	}

	cam.SetControlCapValueByID(constants.ASI_EXPOSURE, 1000, false)

	timeStart := time.Now()
	for i := 0; i < frames; i++ {
		_, err = cam.MakeCapture(false, "../frames_rgb24_bench/")
		if err != nil {
			fmt.Errorf("making shapshot stopped with status: %v", err)
		}
	}
	timeFinish := time.Now()
	timeDiff := timeFinish.Sub(timeStart)

	fmt.Fprintf(w, "Image type: %v\n", imgType.ToString())
	fmt.Fprintf(w, "Time difference: %v\n", timeDiff)
	fmt.Fprintf(w, "Total frames: %v\n", frames)
	fmt.Fprintf(w, "Frames per second: %v\n", float64(frames)/timeDiff.Seconds())
	fmt.Fprintln(w)

	err = cam.Close()
	if err != nil {
		fmt.Errorf("closing src finished with status: %v", err)
	}

	os.RemoveAll("../frames_rgb24_bench/")
	os.MkdirAll("../frames_rgb24_bench/", 0777)
}
