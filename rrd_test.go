package rrd

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestCreateDs(t *testing.T) {
	cleanup()

	values := []string{
		"DS:ok:GAUGE:600:0:U",
		"RRA:AVERAGE:0.5:1:25920",
	}
	err := Create("test.rrd", 5, time.Now().Sub(10), values)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCreateError(t *testing.T) {
	cleanup()

	values := []string{
		"DS:ok:GAUGE:600:0:U",
	}
	err := Create("test.rrd", 5, time.Now().Sub(10), values)
	if err == nil {
		t.Fatalf("Expected error: you must define at least one Round Robin Archive")
	}
}

func TestUpdate(t *testing.T) {
	cleanup()

	values := []string{
		"DS:ok:GAUGE:600:0:U",
		"RRA:AVERAGE:0.5:1:25920",
	}
	err := Create("test.rrd", 15, time.Now().Sub(10), values)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	err = Update("test.rrd", "ok", []string{fmt.Sprintf("%d:%d", time.Now(), 15)})
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
}

func TestUpdateInvalidDs(t *testing.T) {
	cleanup()

	values := []string{
		"DS:ok:GAUGE:600:0:U",
		"RRA:AVERAGE:0.5:1:25920",
	}
	err := Create("test.rrd", 15, time.Now().Sub(10), values)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	err = Update("test.rrd", "fail", []string{fmt.Sprintf("%d:%d", time.Now(), 15)})
	if err == nil {
		t.Fatalf("Expexted error: unknown DS name 'fail'", err)
	}
}

func cleanup() {
	os.Remove("test.rrd")
	clearError()
}
