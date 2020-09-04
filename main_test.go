package dprefix

import (
	"os"
	"time"
	"testing"

	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgb/xproto"
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

func TestLookupString(t *testing.T) {
	X, err := xgbutil.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	keybind.Initialize(X)
	
	str := keybind.LookupString(X,0,xproto.Keycode(19))
	t.Log(str)
	//t.Log(GetString())
}

func TestGetString(t *testing.T) {
	t.Log(GetString())
}
