package src

//#cgo CFLAGS: -I../include
//#cgo LDFLAGS: -L../lib/x64 -lASICamera2 -Wl,-rpath=../lib/x64
//#include <stdio.h>
//#include "ASICamera2.h"
import "C"

// get the product ID of each supported src, at first set pPIDs as 0 and get length and then
// malloc a buffer to contain the PIDs

func GetProductIDs() (pIDs []int) {
	var probe C.int
	length := int(C.ASIGetProductIDs(&probe))
	cpIDs := make([]C.int, length)
	C.ASIGetProductIDs(&cpIDs[0])
	pIDs = append(pIDs, make([]int, length)...)
	for i := 0; i < length; i++ {
		pIDs[i] = int(cpIDs[i])
	}
	return pIDs
}

func GetSDKVersion() string {
	ver := C.ASIGetSDKVersion()
	return C.GoString(ver)
}
