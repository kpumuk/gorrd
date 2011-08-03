package main

import (
	"log"
	"fmt"
	"rrd"
	"time"
)

func main() {
	err := rrd.Create("test.rrd", 10, time.Seconds()-10, []string{
		"DS:ok:GAUGE:600:0:U",
		"RRA:AVERAGE:0.5:1:25920",
	})
	if err != nil {
		log.Exitf("Error: %s", err)
	}

	err = rrd.Update("test.rrd", "ok", []string{
		fmt.Sprintf("%d:%d", time.Seconds(), 15),
	})
	if err != nil {
		log.Exitf("Error: %s", err)
	}

	log.Stdout("Everything is OK")
}
