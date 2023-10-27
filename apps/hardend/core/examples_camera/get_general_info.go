package main

//#cgo CFLAGS: -I./../camera/include
//#cgo LDFLAGS: -L./../camera/lib/x64 -lASICamera2 -Wl,-rpath=./../camera/lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
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

	fmt.Println(src.GetProductIDs())
	fmt.Println(src.GetSDKVersion())

	if cam.IsUSB3Camera() {
		fmt.Println(cam.GetID())
		fmt.Println(cam.SetID([8]byte{'h', 'e', 'l', 'l', 'o'}))
		fmt.Println(cam.GetID())
	}

	supportedModes, err := cam.GetCameraSupportMode()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Supported modes: ")
	for _, mode := range supportedModes {
		fmt.Println(mode.ToString())
	}

	mode, err := cam.GetCameraMode()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Current mode: ")
	fmt.Println(mode.ToString())
	err = cam.Close()
	if err != nil {
		log.Println(err)
	}
}
