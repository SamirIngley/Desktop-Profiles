package main

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
)

func testingOpen() {
	// WORKS !!!
	err := open.Run("/Volumes/Macintosh HD/Applications/Atom.app")
	fmt.Print(err)
}
