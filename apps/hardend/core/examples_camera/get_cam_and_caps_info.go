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
	"strings"
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

	filename := strings.ReplaceAll(fmt.Sprintf("../build/info_%s.txt", cam.Name()), " ", "_")
	outputFile, err := os.OpenFile(filename,
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}

	_, err = outputFile.WriteString(cam.Info() + "\n")
	if err != nil {
		log.Println("Write src info error: ", err)
	}

	capsList := cam.GetListOfControlCaps()

	for _, capName := range capsList {
		_, err = outputFile.WriteString(cam.GetControlCapInfo(capName) + "\n")
		if err != nil {
			log.Println("Write caps info error: ", err)
		}
	}

	err = outputFile.Close()
	if err != nil {
		log.Println(err)
	}

	err = cam.Close()
	if err != nil {
		log.Println(err)
	}
}
