package main

//#cgo CFLAGS: -I./include
//#cgo LDFLAGS: -L./lib/x64 -lASICamera2 -Wl,-rpath=./lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
	"core/camera/constants"
	"core/camera/src"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Number of connected cameras: ", int(C.ASIGetNumOfConnectedCameras()))
	cam, err := src.NewCamera(0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cam)

	if err = cam.Open(); err != nil {
		log.Fatalln(err)
	}

	if err = cam.Init(); err != nil {
		log.Fatalln(err)
	}

	numOfControls, err := cam.GetNumberOfControls()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Number of controls: ", numOfControls)

	if err = cam.InitControlCaps(); err != nil {
		log.Println(err)
	}

	fmt.Println(cam.Info())

	//namesOfControlCaps := cam.GetListOfControlCaps()
	//
	//for _, name := range namesOfControlCaps {
	//	fmt.Println(cam.GetControlCapByName(name).Info())
	//}

	//fmt.Println(cam.GetControlCapInfo("Gain"))
	//fmt.Println(cam.SetControlCapValueByName("Gain", 90, true))
	//fmt.Println(cam.GetControlCapInfo("Gain"))

	roi := src.NewROI(
		cam.MaxWidth(), cam.MaxHeight(),
		0, 0,
		1, constants.ASI_IMG_RGB24)

	if err = cam.SetROIFormat(roi); err != nil {
		log.Println(err)
	}

	if err = cam.SetStartPosition(roi); err != nil {
		log.Println(err)
	}
	ok, err := cam.ExecuteExposure(false)
	if !ok {
		log.Println(err)
	} else {
		if err = cam.GetDataAfterExp(); err != nil {
			log.Println(err)
		}

		if _, err = cam.SaveFrame("../frames"); err != nil {
			log.Println(err)
		}
	}

	err = cam.Close()
	if err != nil {
		log.Println(err)
	}
}
