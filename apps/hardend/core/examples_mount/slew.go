package main

import (
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

	mount.PrepareAfterTurnOn()

	err = mount.VariableRateSlew(celestron.AZM_RA, true, 32703)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(time.Second)
	mount.VariableRateSlew(celestron.AZM_RA, true, 0)
	//
	//loc, err := Earth.NewLocation("118 20 17 W 33 50 41 N")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = mount.SetLocation(loc)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//

	//fmt.Println(com_mis.Echo(mount, 'H'))
	//fmt.Println(com_mis.GetVersion(mount))
	//fmt.Println(com_mis.GetModel(mount))
	//
	//loc, err := Earth.NewLocation("118 20 17 W 33 50 41 N")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//mount.Earth = loc
	//
	//err = com_loc.SetLocation(mount, loc)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//location, err := com_loc.GetLocation(mount)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(location.ToString())

	//ra, dec, err := position.GetRaDec(mount)
	//fmt.Println(ra.ToString(), dec.ToString(), err)
	//azm, alt, err := commands.GetAzmAlt(mount)
	//fmt.Println(azm.ToString(), alt.ToString(), err)
	//ra, dec, err = commands.GetPreciseRaDec(mount)
	//fmt.Println(ra.ToString(), dec.ToString(), err)
	//azm, alt, err = commands.GetPreciseAzmAlt(mount)
	//fmt.Println(azm.ToString(), alt.ToString(), err)

	//commands2.GotoRaDec(mount, []byte{0, 48, 56, 56}, []byte{70, 70, 70, 48})
	//ra, dec, err := commands.GetRaDec(mount)
	//fmt.Println(ra.ToString(), dec.ToString(), err)

	//fmt.Println(commands3.GetMountingPointState(mount))
	//fmt.Println(mount.GetTime())
	//
	////ra, _ := Sky.NewRightAscension(10, 0, 0)
	//
	//timeNow := time.Now()
	//
	//err = mount.SetTime(timeNow)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(mount.GetTime())

	//fmt.Println(cmd_slewing.FixedRateSlew(mount, cmd_slewing.AZM_RA, cmd_slewing.POS_DIRECT_FIXED, 6))
	os.Exit(0)
}
