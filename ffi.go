package main

import (
	"C"

	lib "github.com/itamadev/golang-ffi-python-demo/lib"
)

func print(name *C.char) {
	// Print it
	lib.Print(C.GoString(name))
}
