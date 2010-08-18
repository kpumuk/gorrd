include $(GOROOT)/src/Make.$(GOARCH)

TARG        = rrd
CGOFILES    = rrd.go
CGO_LDFLAGS = `pkg-config --libs librrd`

include $(GOROOT)/src/Make.pkg
