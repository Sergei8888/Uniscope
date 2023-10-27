package main

//#cgo CFLAGS: -I../camera/include
//#cgo LDFLAGS: -L../camera/lib/x64 -lASICamera2 -Wl,-rpath=../camera/lib/x64
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
    "time"
)

func main() {
    timeStart := time.Now()
    camInput, err := utils.ReadNumberOfCamera(os.Stdout, os.Stdin)
    if err != nil {
        log.Fatalln(err)
    }
    diff := time.Now()
    fmt.Println("camInput:", camInput, diff.Sub(timeStart))

    cam, err := src.NewCamera(camInput)
    if err != nil {
        log.Fatalln(fmt.Sprintf("src wasn't created: %v", err))
    }
    diff = time.Now()
    fmt.Println("camera:", cam, diff.Sub(timeStart))

    if err = cam.Open(); err != nil {
        log.Fatalln(err)
    }
    diff = time.Now()
    fmt.Println("camera is opened", diff.Sub(timeStart))

    if err = cam.Init(); err != nil {
        cam.Close()
        log.Fatalln(err)
    }
    diff = time.Now()
    fmt.Println("camera is initialized", diff.Sub(timeStart))

    if err = cam.InitControlCaps(); err != nil {
        cam.Close()
        log.Fatalln(err)
    }
    diff = time.Now()
    fmt.Println("control caps are initialized", diff.Sub(timeStart))

    fmt.Println("\nControl caps info:")
    for _, controlCap := range cam.ControlCaps() {
        fmt.Println(controlCap.Info())
    }

    formats := []constants.ASI_IMG_TYPE{
        constants.ASI_IMG_RAW8,
        constants.ASI_IMG_RAW16,
        constants.ASI_IMG_RGB24,
        constants.ASI_IMG_Y8,
    }

    expLen, _, err := cam.GetControlCapValueByID(constants.ASI_EXPOSURE)
    if err != nil {
        cam.Close()
        log.Fatalln(err)
    }
    diff = time.Now()
    fmt.Println("exposure is got", diff.Sub(timeStart))

    expLen = expLen / 4

    //if err = cam.IsControlCapAutoAligned(constants.ASI_EXPOSURE, true); err != nil {
    //    cam.Close()
    //    log.Fatalln(err)
    //}
    if err = cam.SetControlCapValueByID(constants.ASI_EXPOSURE, 3000); err != nil {
        cam.Close()
        log.Fatalln(err)
    }

    diff = time.Now()
    fmt.Println("exposure is set", diff.Sub(timeStart))

    if err = cam.IsControlCapAutoAligned(constants.ASI_GAIN, true); err != nil {
        cam.Close()
        log.Fatalln(err)
    }
    diff = time.Now()
    fmt.Println("gain is set", diff.Sub(timeStart))

    if err = cam.IsControlCapAutoAligned(constants.ASI_WB_R, true); err != nil {
        cam.Close()
        log.Fatalln(err)
    }
    diff = time.Now()
    fmt.Println("WB_R is set", diff.Sub(timeStart))
    timeStart = time.Now()

    if err = cam.IsControlCapAutoAligned(constants.ASI_WB_B, true); err != nil {
        cam.Close()
        log.Fatalln(err)
    }
    diff = time.Now()
    fmt.Println("WB_B is set", diff.Sub(timeStart))
    timeStart = time.Now()

    if err = cam.IsControlCapAutoAligned(constants.ASI_BANDWIDTHOVERLOAD, true); err != nil {
        cam.Close()
        log.Fatalln(err)
    }
    diff = time.Now()
    fmt.Println("BandWidth is set", diff.Sub(timeStart))
    timeStart = time.Now()

    fmt.Println("\ncamera supports modes")
    fmt.Println(cam.GetCameraSupportMode())

    fmt.Println("\ncamera's mode")
    camMode, err := cam.GetCameraMode()
    if err != nil {
        log.Println(err)
    }
    fmt.Println(camMode)
    err = cam.SetCameraMode(camMode)
    if err != nil {
        log.Println(err)
    }
    diff = time.Now()
    fmt.Println("\ncamera's mode is set", diff.Sub(timeStart))
    fmt.Println()

    cam.SendSoftTrigger(true)

    for _, format := range formats {
        roi := src.NewROI(
            cam.MaxWidth(), cam.MaxHeight(),
            0, 0,
            1, format,
        )
        if err = cam.SetROIFormat(roi); err != nil {
            cam.Close()
            log.Fatalln(err)
        }
        diff = time.Now()
        fmt.Println("ROI format set", diff.Sub(timeStart))
        fmt.Println(cam.ROI)
        timeStart = time.Now()

        if err = cam.SetStartPosition(roi); err != nil {
            cam.Close()
            log.Fatalln(err)
        }
        diff = time.Now()
        fmt.Println("start position set", diff.Sub(timeStart))
        timeStart = time.Now()

        fmt.Println(cam.Info())
        diff = time.Now()
        fmt.Println("cam info printed", diff.Sub(timeStart))
        timeStart = time.Now()

        ok, err := cam.ExecuteExposure(false)
        if !ok {
            log.Println(err)
        } else {
            if err = cam.GetDataAfterExp(); err != nil {
                log.Println(err)
            }

            diff = time.Now()
            fmt.Println("shot without saving:", diff.Sub(timeStart))

            if filename, err := cam.SaveFrame("../frames"); err != nil {
                log.Println(err)
            } else {
                log.Println(fmt.Sprintf("Frame saved, name: %s", filename))
            }
        }
        diff = time.Now()
        fmt.Println("exposure executed, data saved", diff.Sub(timeStart))
        timeStart = time.Now()

        err = cam.StartVideoCapture()
        if err != nil {
            cam.Close()
            log.Fatalln(err)
        }
        diff = time.Now()
        fmt.Println("\nstarted video capture", diff.Sub(timeStart))
        timeStart = time.Now()
        for i := 0; i < 10; i++ {
            timeStartLoop := time.Now()
            err = cam.GetVideoData(int(expLen) + 500)
            if err != nil {
                cam.Close()
                log.Fatalln(err)
            }

            diff = time.Now()
            fmt.Println("video shot without saving:", diff.Sub(timeStartLoop))

            if filename, err := cam.SaveFrame("../frames"); err != nil {
                log.Println(err)
            } else {
                log.Println(fmt.Sprintf("Frame saved, name: %s", filename))
            }
        }
        err = cam.StopVideoCapture()
        if err != nil {
            cam.Close()
            log.Fatalln(err)
        }
        diff = time.Now()
        fmt.Println("stopped video capture", diff.Sub(timeStart))
        timeStart = time.Now()

    }

    if err = cam.Close(); err != nil {
        log.Println("camera didn't close")
    }

}
