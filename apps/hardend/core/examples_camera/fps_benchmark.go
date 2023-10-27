package main

import (
	"core/camera/constants"
	"core/camera/src"
	"core/camera/tests"
	"core/camera/utils"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filename := fmt.Sprintf("../test_info/test_%v", time.Now().UnixNano())
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	camInput, err := utils.ReadNumberOfCamera(os.Stdout, os.Stdin)

	cam, err := src.NewCamera(camInput)
	if err != nil {
		fmt.Errorf("src wasn't created: %v", err)
	}

	err = cam.FullInit(src.NewFullMatrixROI(cam, constants.ASI_IMG_RGB24))
	if err != nil {
		fmt.Errorf("src wasn't initialised: %v", err)
	}

	cam.SetControlCapValueByID(constants.ASI_EXPOSURE, 1000, false)

	fmt.Println(cam.VideoCaptureByTime(5*time.Second, "../video_test/"))

	tests.SaveRGB24Benchmark(file, 20)
	tests.SaveRAW16Benchmark(file, 20)
	tests.SaveRAW8Benchmark(file, 20)
	file.Close()

	err = cam.Close()
	if err != nil {
		log.Println(err)
	}
}
