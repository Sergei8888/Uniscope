package main

//#cgo CFLAGS: -I./../camera/include
//#cgo LDFLAGS: -L./../camera/lib/x64 -lASICamera2 -Wl,-rpath=./../camera/lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
	camConst "core/camera/constants"
	"core/camera/src"
	camUtils "core/camera/utils"
	"core/mount/configs"
	mount "core/mount/src/Mount"
	celestron "core/mount/src/Mount/Commander/Celestron"
	"log"
	"os"
	"time"
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

	startPoint, err := mount.GetAzmAlt()
	if err != nil {
		log.Println(err)
	}

	camIndex, err := camUtils.ReadNumberOfCamera(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	cam, err := src.NewCamera(camIndex)
	if err != nil {
		log.Fatalln(err)
	}

	err = cam.FullInit(src.NewFullMatrixROI(cam, camConst.ASI_IMG_RAW8))
	if err != nil {
		log.Fatalln(err)
	}

	err = mount.VariableRateSlew(celestron.AZM_RA, true, 3600)
	if err != nil {
		mount.VariableRateSlew(celestron.AZM_RA, true, 0)
		log.Println(err)
	}
	frames, dropped, err := cam.VideoCaptureByTime(5*time.Second, "../frames_example/")
	if err != nil {
		log.Println(err)
	}
	log.Println("Saved frames:", frames)
	log.Println("Dropped frames:", dropped)

	err = mount.GotoAzmAlt(startPoint)
	if err != nil {
		log.Println(err)
	}

	mount.GotoLock(os.Stdout)

	log.Println("Done")
}
