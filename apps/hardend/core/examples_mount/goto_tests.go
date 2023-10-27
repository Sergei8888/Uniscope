package main

import (
	"os"
)

func main() {

	//port, err := mount.ConfigurePort(configs.PORT_NAME, configs.MODE)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//mount := selestron.NewMountCelestron(port)
	//
	//if err = mount.PrepareAfterTurnOn(); err != nil {
	//	log.Fatalln(err)
	//}
	//
	//azm, err := Sky.NewAzimuth(10, 0, 0)
	//if err != nil {
	//	log.Println(err)
	//}
	//alt, err := Sky.NewAltitude(20, 0, 0)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//azmHex, err := azm.ToHex()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//altHex, err := alt.ToHex()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//azm, alt, err = mount.GetPreciseAzmAlt()
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println(azm.ToString(), alt.ToString())
	//fmt.Println(azm.ToString(), alt.ToString())
	//
	//if err = mount.GotoAzmAlt(azmHex, altHex); err != nil {
	//	log.Println(err)
	//}
	//
	//for true {
	//	azm, alt, err = mount.GetPreciseAzmAlt()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	fmt.Println(azm.ToString(), alt.ToString())
	//	fmt.Println(azm.ToString(), alt.ToString())
	//	if isGoto, err := mount.IsGotoInProgress(); !isGoto || err != nil {
	//		break
	//	}
	//	time.Sleep(500 * time.Millisecond)
	//}

	os.Exit(0)
}
