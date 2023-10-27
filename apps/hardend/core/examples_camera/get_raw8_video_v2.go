package main

//#cgo CFLAGS: -I./../include
//#cgo LDFLAGS: -L./../lib/x64 -lASICamera2 -Wl,-rpath=./../lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
	"core/camera/constants"
	"core/camera/src"
	"core/camera/utils"
	"log"
	"os"
	"time"
)

func main() {

	camInput, err := utils.ReadNumberOfCamera(os.Stdout, os.Stdin)

	cam, err := src.NewCamera(camInput)
	if err != nil {
		log.Fatalln(err)
	}

	roi := src.NewROI(
		cam.MaxWidth(), cam.MaxHeight(),
		0, 0,
		1, constants.ASI_IMG_RGB24,
	)

	err = cam.FullInit(roi)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(5 * time.Second)
	log.Println("Staring capture")

	frames, dropped, err := cam.VideoCaptureByTime(20*time.Second, "../video_test_rgb24/")
	if err != nil {
		log.Println(err)
	}
	log.Println("Total frames", frames)
	log.Println("Dropped frames:", dropped)

	err = cam.Close()
	if err != nil {
		log.Println(err)
	}
}
