// This file is not needed for the Python demo, but is used to show a native Go execution.

package main

import (
	"github.com/itamadev/golang-ffi-python-demo/pkg/pkg"
)

func main() {
	pkg.Print("From Go!")
}
