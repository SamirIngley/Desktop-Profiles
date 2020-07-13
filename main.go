package main

import (
	"fmt"
	"runtime"
	"os"
	"os/exec"
	"log"
	"io/ioutil"
	
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


// reads Profile file
func readFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(dat)
}


// open url in browser
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

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