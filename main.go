package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// type profile struct {
// 	name string

// }

// take in user arguments, store as JSON or .txt
func takeInput() {
	// figure out if adding to a profile or opening a profile
	fmt.Println("Desktop Profiles...")
	name := flag.String("name", "profile-name", "profile name")
	open := flag.Bool("open", true, "open this profile")
	edit := flag.Bool("edit", false, "do you want to modify a file")
	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*open)
	fmt.Println(*edit)

}

func createProfile() {

}

func openProfile() {

}


// open file with specified path
func openFile() {
	file, err := os.Open("proposal.md") // For read access.
	bytes, _ := ioutil.ReadAll(file)
	stringBody := string(bytes)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(file)
	fmt.Print(stringBody)
}

// open app on desktop

// func main() {
// 	// openBrowser("https://www.google.com")
// 	// openBrowser("https://www.hackernews.com")

// }
