package main

import (
	"io/ioutil"
	"os"
	"time"
	"log"

	"github.com/lolbinarycat/dprefix"
)

func init() {
	log.SetOutput(ioutil.Discard)
}


func main() {
	go func() {
		time.Sleep(time.Second * 5)
		os.Exit(5)
	} ()
	str := dprefix.GetString()
	os.Stdout.Write([]byte(str))
}
