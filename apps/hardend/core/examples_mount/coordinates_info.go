package main

import (
	"core/mount/configs"
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

	horizCS, err := mount.GetPreciseAzmAlt()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(horizCS.Azm().ToString(), horizCS.Alt().ToString())

	eqCS, err := mount.GetPreciseRaDec()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(eqCS.Ra().ToString(), eqCS.Dec().ToString())

	os.Exit(0)
}
