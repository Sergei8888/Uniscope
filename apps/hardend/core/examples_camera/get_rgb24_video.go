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
	"fmt"
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

	if err = cam.Open(); err != nil {
		log.Fatalln(err)
	}

	if err = cam.Init(); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	if err = cam.InitControlCaps(); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	roi := src.NewROI(
		cam.MaxWidth(), cam.MaxHeight(),
		0, 0,
		1, constants.ASI_IMG_RGB24,
	)

	if err = cam.SetROIFormat(roi); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	if err = cam.SetStartPosition(roi); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	expLen, _, err := cam.GetControlCapValueByID(constants.ASI_EXPOSURE)
	if err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	if err = cam.SetControlCapValueByID(constants.ASI_EXPOSURE, expLen/2, false); err != nil {
		cam.Close()
		log.Fatalln(err)
	}
	waitMs := int(expLen)*2 + 500
	fmt.Println("WaitMS: ", waitMs)
	//go func() {
	//	fmt.Println("capturing")
	//	if err = cam.GetVideoData(waitMs); err != nil {
	//		cam.Close()
	//		log.Fatalln(err)
	//	}
	//}()

	if err = cam.StartVideoCapture(); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("capturing")
	for i := 0; i < 10; i++ {
		if err = cam.GetVideoData(waitMs); err != nil {
			cam.Close()
			log.Fatalln(err)
		}
		if filename, err := cam.SaveFrame("../video_test/"); err != nil {
			log.Println(err)
		} else {
			log.Println(fmt.Sprintf("Frame saved, name: %s", filename))
		}
		time.Sleep(time.Duration(waitMs) * time.Microsecond)
		droppedFrames, err := cam.GetDroppedFramesCount()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Dropped frames: ", droppedFrames)
		}
	}

	//reader := bufio.NewReader(os.Stdin)
	//text, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println(err)
	//}
	//ch <- text
	//close(ch)

	err = cam.Close()
	if err != nil {
		log.Println(err)
	}
}
