// dprefix is a utility designed for creating prefix commands for dwm.
package dprefix

import (
	//"os"
	"strings"
	"time"
	
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)


// GetRaw grabs the keyboard on the root window until
// a key is pressed, then returns the raw event for that key.
// It also returns the X connection it openes,
// which is useful if you need to use functions such as
// keybind.LookupString.
func GetRaw() (xevent.KeyPressEvent, *xgbutil.XUtil, bool) {
	X, err := xgbutil.NewConn()
	if err != nil {
		panic(err)
	}
	keybind.Initialize(X)
	keyCh := NextKeyPressChan(X,true,time.Second*10)
	key, ok := <-keyCh
	return key, X, ok
}

// GetKeyWithMods returns the textual representation of a key,
// such as "w" or " ", along with a slice of
// the modifiers pressed, in the format of "mod1", "mod2", etc.
func GetKeyWithMods() (key string, mods []string) {
	e, X, _ := GetRaw()
	// TODO: handle error
	modStr := keybind.ModifierString(e.State)
	key = keybind.LookupString(X, e.State, e.Detail)
	mods = strings.Split(modStr,"-")
	return key, mods
}

// GetString returns the keys pressed in the form of "mod1-shift-a".
// The last charachter should always be the key pressed.
func GetString() (string, bool) {
	e, X, ok := GetRaw()
	if !ok {return "",false}
	modStr := keybind.ModifierString(e.State)
	key := keybind.LookupString(X, e.State, e.Detail)
	if len(modStr) == 0 {
		return key, true
	}
	return modStr + "-" + key, true
}

// NextKeyPressChan grabs the next key press on the root window and sends it through a channel.
// If ignoreMods is true, events caused by pressing a modifier key will
// be ignored (modifier data for other events will be unchanged)
// A timeout of less than 1 implies no timeout.
// Note that the returned channel is only valid for one key press.
func NextKeyPressChan(X *xgbutil.XUtil, ignoreMods bool, timeout time.Duration) <-chan xevent.KeyPressEvent {
	keyChan := make(chan xevent.KeyPressEvent)
	if timeout > 0 {
		go func () {
			time.Sleep(timeout)
			keybind.UngrabKeyboard(X)
			xevent.Quit(X)
			close(keyChan)
		} ()
	}
	
	xevent.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			if ignoreMods && keybind.ModGet(X,e.Detail) != 0 {
				// ModGet returns 0 if the given
				// keycode isn't a modifier
				return
			}
			keyChan <- e
			// modStr := keybind.ModifierString(e.State)
			// keyStr := keybind.LookupString(X, e.State, e.Detail)
			// if len(modStr) > 0 {
			// 	log.Printf("Key: %s-%s\n", modStr, keyStr)
			// } else {
			// 	log.Println("Key:", keyStr)
			// }
			keybind.UngrabKeyboard(X)
			xevent.Quit(X)
		}).Connect(X,X.RootWin())
	keybind.GrabKeyboard(X, X.RootWin())
	go xevent.Main(X)
	return keyChan
}
