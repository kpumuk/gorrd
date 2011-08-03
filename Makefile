include $(GOROOT)/src/Make.inc

TARG        = rrd
CGOFILES    = rrd.go

format:
	find . -type f -name '*.go' -exec gofmt -w {} ';'

include $(GOROOT)/src/Make.pkg
