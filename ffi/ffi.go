package main

import (
	"C"

	"github.com/itamadev/golang-ffi-python-demo/pkg/pkg"
)

//export print
func print(name *C.char) {
	pkg.Print(C.GoString(name))
}

func main() {}
