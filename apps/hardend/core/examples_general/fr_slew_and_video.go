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

// (not tested)
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

	roi := src.NewROI(
		cam.MaxWidth(), cam.MaxHeight(),
		0, 0,
		1, camConst.ASI_IMG_RGB24,
	)

	err = cam.FullInit(roi)
	if err != nil {
		log.Fatalln(err)
	}

	err = mount.FixedRateSlew(celestron.AZM_RA, false, 0x09)
	if err != nil {
		mount.FixedRateSlew(celestron.AZM_RA, false, 0x00)
		log.Println(err)
	}
	frames, dropped, err := cam.VideoCaptureByTime(5*time.Second, "../frames_example/")
	if err != nil {
		log.Println(err)
	}
	log.Println("Saved frames:", frames)
	log.Println("Dropped frames:", dropped)

	err = mount.FixedRateSlew(celestron.AZM_RA, true, 0x09)
	if err != nil {
		mount.FixedRateSlew(celestron.AZM_RA, true, 0x00)
		log.Println(err)
	}
	time.Sleep(5 * time.Second)
	if err := mount.FixedRateSlew(celestron.AZM_RA, true, 0x00); err != nil {
		log.Println(err)
	}

	log.Println("Done")
}
