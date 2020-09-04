// dprefix is a utility designed for creating prefix commands for dwm.
package dprefix

import (
	//"os"
	"strings"
	
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)


// GetRaw grabs the keyboard on the root window until
// a key is pressed, then returns the raw event for that key.
// It also returns the X connection it openes,
// which is useful if you need to use functions such as
// keybind.LookupString.
func GetRaw() (xevent.KeyPressEvent, *xgbutil.XUtil) {
	X, err := xgbutil.NewConn()
	if err != nil {
		panic(err)
	}
	keybind.Initialize(X)
	keyCh := NextKeyPressChan(X)
	return <-keyCh, X
}

// GetKeyWithMods returns the textual representation of a key,
// such as "w" or " ", along with a slice of
// the modifiers pressed, in the format of "mod1", "mod2", etc.
func GetKeyWithMods() (key string, mods []string) {
	e, X := GetRaw()
	modStr := keybind.ModifierString(e.State)
	key = keybind.LookupString(X, e.State, e.Detail)
	mods = strings.Split(modStr,"-")
	return key, mods
}

// GetString returns the keys pressed in the form of "mod1-shift-a".
// The last charachter should always be the key pressed.
func GetString() string {
	e, X := GetRaw()
	modStr := keybind.ModifierString(e.State)
	key := keybind.LookupString(X, e.State, e.Detail)
	if len(modStr) == 0 {
		return key
	}
	return modStr + "-" + key
}

// GetEmacs returns the emacs representation of the keys pressed,
// with Alt mapped to Meta.
func GetEmacs() string {
	return "Not Implemented"
}

// NextKeyPressChan grabs the next key press on the root window and sends it through a channel.
// Note that the returned channel is only valid for one key press
func NextKeyPressChan(X *xgbutil.XUtil) <-chan xevent.KeyPressEvent {
	keyChan := make(chan xevent.KeyPressEvent)
	xevent.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			keyChan <- e
			// modStr := keybind.ModifierString(e.State)
			// keyStr := keybind.LookupString(X, e.State, e.Detail)
			// if len(modStr) > 0 {
			// 	log.Printf("Key: %s-%s\n", modStr, keyStr)
			// } else {
			// 	log.Println("Key:", keyStr)
			// }
			xevent.Quit(X)
		}).Connect(X,X.RootWin())
	keybind.GrabKeyboard(X, X.RootWin())
	go xevent.Main(X)
	return keyChan
}
