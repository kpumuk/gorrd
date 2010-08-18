package rrd
//
// This is go-bindings package for librrd
//

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "rrd.h"
*/
import "C"

import (
    "os"
    "unsafe"
    // "fmt"
    // "container/vector"
)

func Create(filename string, pdp_step, last_up int64, values []string) (err os.Error) {
    cfilename := C.CString(filename)
    defer C.free(unsafe.Pointer(cfilename))

    cvalues := makeCStringArray(values)
    defer freeCStringArray(cvalues)

    ret := C.rrd_create_r(cfilename, C.ulong(pdp_step), C.time_t(last_up),
        C.int(len(values)), getCStringArrayPointer(cvalues))

    if int(ret) != 0 { err = os.NewError(Error()) }
    return
}

func Update(filename, template string, values []string) (err os.Error) {
    cfilename := C.CString(filename)
    defer C.free(unsafe.Pointer(cfilename))

    ctemplate := C.CString(template)
    defer C.free(unsafe.Pointer(ctemplate))

    cvalues := makeCStringArray(values)
    defer freeCStringArray(cvalues)

    ret := C.rrd_update_r(cfilename, ctemplate,
        C.int(len(values)), getCStringArrayPointer(cvalues))

    if int(ret) != 0 { err = os.NewError(Error()) }
    return
}

func Error() string {
    return C.GoString(C.rrd_get_error())
}

func clearError() {
    C.rrd_clear_error()
}

func makeCStringArray(values []string) (cvalues []*C.char) {
    cvalues = make([]*C.char, len(values))
    for i := range values {
        cvalues[i] = C.CString(values[i])
    }
    return
}

func freeCStringArray(cvalues []*C.char) {
    for i := range cvalues {
        C.free(unsafe.Pointer(cvalues[i]))
    }
}

func getCStringArrayPointer(cvalues []*C.char) **C.char {
    return (**C.char)(unsafe.Pointer(&cvalues[0]))
}
