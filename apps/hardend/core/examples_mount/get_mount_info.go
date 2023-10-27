package main

import (
	"core/mount/configs"
	mount "core/mount/src/Mount"
	selestron "core/mount/src/Mount/Commander/Celestron"
	"fmt"
	"log"
	"os"
)

func main() {

	port, err := mount.ConfigurePort(configs.PORT_NAME, configs.MODE)
	if err != nil {
		log.Fatalln(err)
	}
	mount := selestron.NewMountCelestron(port)

	err = mount.PrepareAfterTurnOn()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Echo: ")
	fmt.Println(mount.Echo('h'))
	fmt.Print("GetDeviceVersion: ")
	fmt.Println(mount.GetDeviceVersion(true))
	fmt.Print("GetModel: ")
	fmt.Println(mount.GetModel())
	//fmt.Print("GetMountPointState: ")
	//fmt.Println(mount.GetMountPointState())
	fmt.Print("GetVersion: ")
	fmt.Println(mount.GetVersion())
	fmt.Print("IsAlignmentComplete: ")
	fmt.Println(mount.IsAlignmentComplete())
	fmt.Print("IsGotoInProgress: ")
	fmt.Println(mount.IsGotoInProgress())
	fmt.Println(mount)
	os.Exit(0)
}
