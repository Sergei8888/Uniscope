package utils

//#cgo CFLAGS: -I./../include
//#cgo LDFLAGS: -L./../lib/x64 -lASICamera2 -Wl,-rpath=./../lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ReadNumberOfCamera(w io.Writer, r io.Reader) (camIndex int, err error) {
	numOfConnCams := int(C.ASIGetNumOfConnectedCameras())
	if numOfConnCams == 0 {
		return 0, errors.New("no cameras connected")
	}

	if numOfConnCams == 1 {
		return 0, nil
	}
	_, err = fmt.Fprintln(w, fmt.Sprintf("Choose number of src from 0 to %d:", numOfConnCams-1))
	if err != nil {
		return 0, err
	}
	reader := bufio.NewReader(r)
	camIndexInput, _ := reader.ReadString('\n')

	camIndexInput = strings.Trim(camIndexInput, "\r\n")

	camIndex, err = strconv.Atoi(camIndexInput)
	if err != nil {
		return 0, err
	}
	if camIndex > numOfConnCams-1 {
		fmt.Fprintln(w, fmt.Sprintf("Number %d is larger than %d", camIndex, numOfConnCams-1))
	}

	return camIndex, nil
}
