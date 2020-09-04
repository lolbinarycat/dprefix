package main

import "C"

import "github.com/lolbinarycat/dprefix"

//export dprefix_get_string
func dprefix_get_string() *C.char {
	return C.CString(dprefix.GetString())
}

func main() {}
