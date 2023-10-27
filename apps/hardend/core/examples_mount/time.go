package main

import (
    "core/mount/configs"
    mount "core/mount/src/Mount"
    celestron "core/mount/src/Mount/Commander/Celestron"
    "fmt"
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

    fmt.Println(mount.GetTime())

    timeNow := time.Now()

    err = mount.SetTime(timeNow)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(mount.GetTime())
    os.Exit(0)
}
