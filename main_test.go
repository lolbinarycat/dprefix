package dprefix

import (
	"os"
	"time"
	"testing"
)

func TestGetRaw(t *testing.T) {
	t.Log("starting")
	go func() {
		time.Sleep(time.Second * 10)
		// if we mess up un-grabbing the root window,
		// this is here to save us.
		t.Log("timeout exeeded, exiting")
		os.Exit(5)
		t.Fail()
		t.FailNow()
	} ()

	k, _ := GetRaw()
	t.Log(k)
	t.Log(k.String())
}

func TestGetString(t *testing.T) {
	t.Log(GetString())
}
