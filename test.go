package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/skratchdot/open-golang/open"
)

func testingOpenApp() {
	// test opening an app
	err := open.Run("/Volumes/Macintosh HD/Applications/Atom.app")
	fmt.Print(err)
}

func testingOpenURL() {
	// testing url open
	err := open.Run("https://www.google.com")
	fmt.Print(err)
}

func testIfAppDir() bool {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	exists := false
	extList := []string{}

	filepath.Walk(path, func(path string, fileInfo os.FileInfo, _ error) error {
		if fileInfo.Name() == "appDir.txt" {
			extList = append(extList, fileInfo.Name())
			exists = true
		}
		return nil
	})
	return exists
}
