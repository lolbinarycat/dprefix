package main

import (
	"os"

	"github.com/lolbinarycat/dprefix"
)

func main() {
	str := dprefix.GetString()
	os.Stdout.Write([]byte(str))
}
