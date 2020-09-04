// dprefix is a utility designed for creating prefix commands for dwm.
package dprefix

import (
	"os"
	"log"
	
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

func init() {
	log.SetOutput(os.Stderr)
}

func GetRaw() (xevent.KeyPressEvent) {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	keyCh := NextKeyPressChan(X)
	//defer close(keyCh)
	return <-keyCh
}

// NextKeyPressChan grabs the next key press on the root window and sends it through a channel.
// Note that the returned channel is only valid for one key press
func NextKeyPressChan(X *xgbutil.XUtil) <-chan xevent.KeyPressEvent {
	keyChan := make(chan xevent.KeyPressEvent)
	log.Println("connecting event")
	xevent.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println("got event")
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
