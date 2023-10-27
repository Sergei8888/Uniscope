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
		1, constants.ASI_IMG_RAW8,
	)

	if err = cam.SetROIFormat(roi); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	if err = cam.SetStartPosition(roi); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	if err = cam.SetControlCapValueByID(constants.ASI_GAIN, 100, true); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	fmt.Println(cam.GetControlCapValueByID(constants.ASI_GAIN))

	if err = cam.SetControlCapValueByID(constants.ASI_EXPOSURE, 20000, true); err != nil {
		cam.Close()
		log.Fatalln(err)
	}

	fmt.Println(cam.GetControlCapValueByID(constants.ASI_EXPOSURE))

	ok, err := cam.ExecuteExposure(false)
	if !ok {
		log.Println(err)
	} else {
		if err = cam.GetDataAfterExp(); err != nil {
			log.Println(err)
		}

		if filename, err := cam.SaveFrame("../frames"); err != nil {
			log.Println(err)
		} else {
			log.Println(fmt.Sprintf("Frame saved, name: %s", filename))
		}
	}

	err = cam.Close()
	if err != nil {
		log.Println(err)
	}
}
