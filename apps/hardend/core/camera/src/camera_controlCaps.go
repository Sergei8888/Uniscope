package src

//#cgo CFLAGS: -I../include
//#cgo LDFLAGS: -L../lib/x64 -lASICamera2 -Wl,-rpath=../lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
    "core/camera/constants"
    "errors"
    "fmt"
)

func (c *Camera) GetNumberOfControls() (int, error) {
    var numberOfControls C.int
    status := C.ASIGetNumOfControls(C.int(c.cameraID), &numberOfControls)
    if int(status) != 0 {
        return 0, errors.New(fmt.Sprintf("getting number of controls finished with status: %d", int(status)))
    }
    return int(numberOfControls), nil
}

func (c *Camera) InitControlCaps() error {
    controlsNum, err := c.GetNumberOfControls()
    if err != nil {
        return err
    }

    for i := 0; i < controlsNum; i++ {
        c.controlCaps = append(c.controlCaps, newControlCap(i, c.cameraID))
    }

    return nil
}

func (c *Camera) GetListOfControlCaps() (names []string) {
    for _, controlCap := range c.controlCaps {
        names = append(names, controlCap.name)
    }
    return
}

func (c *Camera) getControlCapByName(name string) *ControlCap {
    for _, controlCap := range c.controlCaps {
        if controlCap.name == name {
            return controlCap
        }
    }
    return nil
}

func (c *Camera) getControlCapByID(controlID constants.ASI_CONTROL_TYPE) *ControlCap {
    for _, controlCap := range c.controlCaps {
        if controlCap.controlType == controlID {
            return controlCap
        }
    }
    return nil
}

func (c *Camera) GetControlCapInfo(name string) string {
    controlCap := c.getControlCapByName(name)
    if controlCap != nil {
        return controlCap.Info()
    }
    return ""
}

func (c *Camera) GetControlCapValueByName(name string) (currValue int64, bAuto bool, err error) {
    controlCap := c.getControlCapByName(name)
    if err := controlCap.getValue(c.cameraID); err != nil {
        return 0, false, err
    }
    return controlCap.currentValue, controlCap.isAutoTurnedOn, nil
}

func (c *Camera) SetControlCapValueByName(name string, newValue int64, bAuto bool) error {
    controlCap := c.getControlCapByName(name)
    if err := controlCap.setValue(c.cameraID, newValue); err != nil {
        return err
    }
    if err := controlCap.changeAutoValue(c.cameraID, bAuto); err != nil {
        return err
    }
    return nil
}

func (c *Camera) GetControlCapValueByID(controlID constants.ASI_CONTROL_TYPE) (currValue int64, bAuto bool, err error) {
    controlCap := c.getControlCapByID(controlID)
    if err := controlCap.getValue(c.cameraID); err != nil {
        return 0, false, err
    }
    return controlCap.currentValue, controlCap.isAutoTurnedOn, nil
}

func (c *Camera) SetControlCapValueByID(controlID constants.ASI_CONTROL_TYPE, newValue int64) error {
    controlCap := c.getControlCapByID(controlID)
    if err := controlCap.setValue(c.cameraID, newValue); err != nil {
        return err
    }
    return nil
}

func (c *Camera) IsControlCapAutoAligned(controlID constants.ASI_CONTROL_TYPE, bAuto bool) error {
    controlCap := c.getControlCapByID(controlID)
    if err := controlCap.changeAutoValue(c.cameraID, bAuto); err != nil {
        return err
    }
    return nil
}
