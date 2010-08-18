# gorrd

Simple Go interface for librrd library.

## Installation

1. Make sure you have the a working Go environment. See the [install instructions](http://golang.org/doc/install.html). gorrd will always compile on the `release` tag.
2. git clone git://github.com/kpumuk/gorrd.git
3. cd gorrd && make install && make test

Please note: `goinstall github.com/kpumuk/gorrd` does not work right now because of goinstall problems with cgo.

## Example

    package main
    
    import (
        "log"
        "fmt"
        "rrd"
        "time"
    )
    
    func main() {
        err := rrd.Create("test.rrd", 10, time.Seconds() - 10, []string {
            "DS:ok:GAUGE:600:0:U",
            "RRA:AVERAGE:0.5:1:25920",
        })
        if err != nil { log.Exitf("Error: %s", err) }
        
        err = rrd.Update("test.rrd", "ok", []string {
            fmt.Sprintf("%d:%d", time.Seconds(), 15),
        })
        if err != nil { log.Exitf("Error: %s", err) }
        
        log.Stdout("Everything is OK")
    }

To run the application, put the code in a file called rrdtest.go and run:

    8g rrdtest.go && 8l -o rrdtest rrdtest.8 && ./rrdtest
