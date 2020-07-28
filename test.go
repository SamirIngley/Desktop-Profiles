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

// test opening with os exec, so we can get PID to close as well
func newOpen() {
	path := 
	cmd := exec.Command(path)
		err := cmd.Start()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(30 * time.Second):      // Kills the process after 30 seconds
		if err := cmd.Process.Kill(); err != nil {
		log.Fatal("failed to kill: ", err)
		}
		<-done // allow goroutine to exit
		log.Println("process killed")
		indexInit()
		case err := <-done:
		if err!=nil{
			log.Printf("process done with error = %v", err)
		}
	}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Waiting for command to finish...")
	//timer() // The time goes by...
		err = cmd.Wait()
	}
}

