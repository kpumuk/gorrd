// This is go-bindings package for librrd
package rrd

// #cgo LDFLAGS: -lrrd_th
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include "rrd.h"
import "C"

import (
	"errors"
	"unsafe"
)

// The Create function lets you set up new Round Robin Database (RRD) files.
// The file is created at its final, full size and filled with *UNKNOWN* data.
//
//      filename::
//          The name of the RRD you want to create. RRD files should end with the
//          extension .rrd. However, it accept any filename.
//      step::
//          Specifies the base interval in seconds with which data will be
//          fed into the RRD.
//      start_time::
//          Specifies the time in seconds since 1970-01-01 UTC when the first
//          value should be added to the RRD. It will not accept any data timed
//          before or at the time specified.
//      values::
//          A list of strings identifying datasources (in format "DS:ds-name:DST:dst arguments")
//          and round robin archives - RRA (in format "RRA:CF:cf arguments").
//          There should be at least one DS and RRA.
//
// See http://oss.oetiker.ch/rrdtool/doc/rrdcreate.en.html for detauls.
//
func Create(filename string, step, start_time int64, values []string) (err error) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	cvalues := makeCStringArray(values)
	defer freeCStringArray(cvalues)

	clearError()
	ret := C.rrd_create_r(cfilename, C.ulong(step), C.time_t(start_time),
		C.int(len(values)), getCStringArrayPointer(cvalues))

	if int(ret) != 0 {
		err = errors.New(error_())
	}
	return
}

// The Update function feeds new data values into an RRD. The data is time aligned
// (interpolated) according to the properties of the RRD to which the data is written.
//
//      filename::
//          The name of the RRD you want to create. RRD files should end with the
//          extension .rrd. However, it accept any filename.
//      template::
//          The template switch allows you to specify which data sources you are going
//          to update and in which order. If the data sources specified in the
//          template are not available in the RRD file, the update process will
//          abort with an error. Format: "ds-name[:ds-name]..."
//      values::
//          A list of strings identifying values to be updated with corresponding
//          timestamps.
//
// See http://oss.oetiker.ch/rrdtool/doc/rrdupdate.en.html for detauls.
//
func Update(filename, template string, values []string) (err error) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	ctemplate := C.CString(template)
	defer C.free(unsafe.Pointer(ctemplate))

	cvalues := makeCStringArray(values)
	defer freeCStringArray(cvalues)

	clearError()
	ret := C.rrd_update_r(cfilename, ctemplate,
		C.int(len(values)), getCStringArrayPointer(cvalues))

	if int(ret) != 0 {
		err = errors.New(error_())
	}
	return
}

//----- Helper methods ---------------------------------------------------------

func error_() string {
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
