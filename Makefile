include $(GOROOT)/src/Make.inc

TARG        = rrd
CGOFILES    = rrd.go
CGO_LDFLAGS = `pkg-config --libs librrd`

include $(GOROOT)/src/Make.pkg
