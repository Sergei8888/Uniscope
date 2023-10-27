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

type ControlCap struct {
	name            string //control type name，like "Gain", "Exposure"...
	description     string //control parameter description
	maxValue        int64  //maximum value
	minValue        int64  //minimum value
	defaultValue    int64  //default value
	currentValue    int64
	isAutoSupported bool //is auto adjust supported?
	isAutoTurnedOn  bool
	isWritable      bool                       //can be adjusted, for example sensor temperature can't be modified
	controlType     constants.ASI_CONTROL_TYPE //control type ID
}

func newControlCap(iControlIndex int, iCameraID int) *ControlCap {
	var camControlCap C.struct__ASI_CONTROL_CAPS
	C.ASIGetControlCaps((C.int)(iCameraID), (C.int)(iControlIndex), &camControlCap)

	controlType := constants.ASI_CONTROL_TYPE(camControlCap.ControlType)
	maxValue := int64(camControlCap.MaxValue)
	minValue := int64(camControlCap.MinValue)

	if controlType == constants.ASI_TEMPERATURE {
		maxValue *= 10
		minValue *= 10
	}

	return &ControlCap{
		name:            C.GoString(&camControlCap.Name[0]),
		description:     C.GoString(&camControlCap.Description[0]),
		maxValue:        maxValue,
		minValue:        minValue,
		defaultValue:    int64(camControlCap.DefaultValue),
		currentValue:    int64(camControlCap.DefaultValue),
		isAutoSupported: int(camControlCap.IsAutoSupported) != 0,
		isAutoTurnedOn:  false,
		isWritable:      int(camControlCap.IsWritable) != 0,
		controlType:     constants.ASI_CONTROL_TYPE(camControlCap.ControlType),
	}
}

func (c *ControlCap) Info() string {
	return fmt.Sprintf(
		`Name: %s
Description: %s
Max Value: %d
Min Value: %d
Default Value: %d
Current Value: %d
Is Auto Supported: %t
is Auto Turned On: %t
Is Writable: %t
Control Type Index: %d
`,
		c.name, c.description,
		c.maxValue, c.minValue,
		c.defaultValue, c.currentValue,
		c.isAutoSupported, c.isAutoTurnedOn,
		c.isWritable, c.controlType,
	)
}

func (c *ControlCap) setValue(iCameraID int, lValue int64) error {

	if !c.isWritable {
		return errors.New("value is not writable")
	}

	if lValue < c.minValue || lValue > c.maxValue {
		return errors.New(
			fmt.Sprintf("invalid value for %s, can be between %d and %d, got: %d",
				c.name, c.minValue, c.maxValue, lValue),
		)
	}

	status := C.ASISetControlValue(C.int(iCameraID), C.int(c.controlType), C.long(lValue), C.int(0))
	if status != 0 {
		return errors.New(fmt.Sprintf("setting of value finished with status: %d", int(status)))
	}

	c.currentValue = lValue

	return nil
}

func (c *ControlCap) changeAutoValue(iCameraID int, bAuto bool) error {

	if !c.isWritable {
		return errors.New("value is not writable")
	}

	if bAuto && !c.isAutoSupported {
		return errors.New("auto adjust is not supported")
	}

	var cBAuto C.int
	if bAuto {
		cBAuto = 1
	}

	status := C.ASISetControlValue(C.int(iCameraID), C.int(c.controlType), C.long(c.currentValue), cBAuto)
	if status != 0 {
		return errors.New(fmt.Sprintf("setting of value finished with status: %d", int(status)))
	}

	c.isAutoTurnedOn = bAuto

	return nil
}

func (c *ControlCap) getValue(iCameraID int) error {

	var plValue C.long
	var cbAuto C.int
	status := C.ASIGetControlValue(C.int(iCameraID), C.int(c.controlType), &plValue, &cbAuto)
	if status != 0 {
		return errors.New(fmt.Sprintf("getting of value finished with status: %d", int(status)))
	}

	var bAuto bool
	if cbAuto == 1 {
		bAuto = true
	}

	c.currentValue = int64(plValue)
	c.isAutoTurnedOn = bAuto

	return nil
}
