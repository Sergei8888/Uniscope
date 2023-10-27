package main

//#cgo CFLAGS: -I./../camera/include
//#cgo LDFLAGS: -L./../camera/lib/x64 -lASICamera2 -Wl,-rpath=./../camera/lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
	"core/camera/constants"
	"core/camera/src"
	camUtils "core/camera/utils"
	"core/mount/configs"
	"core/mount/src/Coordinates/Sky"
	mount "core/mount/src/Mount"
	celestron "core/mount/src/Mount/Commander/Celestron"
	"fmt"
	"log"
	"os"
)

func main() {
	port, err := mount.ConfigurePort(configs.PORT_NAME, configs.MODE)
	if err != nil {
		log.Fatalln(err)
	}

	mount := celestron.NewMountCelestron(port)

	if err = mount.PrepareAfterTurnOn(); err != nil {
		log.Fatalln(err)
	}

	camIndex, err := camUtils.ReadNumberOfCamera(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	cam, err := src.NewCamera(camIndex)
	if err != nil {
		log.Fatalln(err)
	}

	err = cam.FullInit(src.NewFullMatrixROI(cam, constants.ASI_IMG_RAW8))
	if err != nil {
		log.Fatalln(err)
	}

	err = cam.SetControlCapValueByID(constants.ASI_EXPOSURE, 50000, false)
	if err != nil {
		log.Println(err)
	}

	destDir := "../frames_example/"

	fileName, err := cam.MakeCapture(false, destDir)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Frame saved as:", fileName)
	}

	azm, err := Sky.NewAzimuth(20, 0, 0)
	alt, err := Sky.NewAltitude(0, 0, 0)
	p1 := Sky.NewHorizontalCS(azm, alt)
	if err != nil {
		log.Fatalln(err)
	}

	azm, err = Sky.NewAzimuth(10, 0, 0)
	alt, err = Sky.NewAltitude(10, 0, 0)
	p2 := Sky.NewHorizontalCS(azm, alt)
	if err != nil {
		log.Fatalln(err)
	}

	azm, err = Sky.NewAzimuth(0, 0, 0)
	alt, err = Sky.NewAltitude(0, 0, 0)
	p3 := Sky.NewHorizontalCS(azm, alt)
	if err != nil {
		log.Fatalln(err)
	}

	points := []*Sky.HorizontalCS{p1, p2, p3}
	imgTypes := []constants.ASI_IMG_TYPE{constants.ASI_IMG_RAW16, constants.ASI_IMG_RGB24, constants.ASI_IMG_Y8}
	for i, point := range points {
		err = mount.GotoAzmAlt(point)
		if err != nil {
			log.Println(err)
		}

		mount.GotoLock(os.Stdout)

		err = cam.SetROIFormat(src.NewFullMatrixROI(cam, imgTypes[i]))
		if err != nil {
			log.Println(err)
		}

		fileName, err = cam.MakeCapture(false, destDir)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Frame saved as:", fileName)
		}
	}
	err = cam.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Done")
}
