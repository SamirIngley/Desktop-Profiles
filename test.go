package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

type Config struct {
	Profile struct {
		name  string `json: name`
		files string `json: files`
	}
}

func testingOpen() {
	route := "/Volumes/Macintosh HD/Applications/Calculator.app"
	err := open.Run(route)
	fmt.Print(err)
}

func seekProfile(ext string, name string) bool {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	extList := []string{}

	// finds all txt files in current directory
	filepath.Walk(path, func(path string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(path) == ext {
			extList = append(extList, fileInfo.Name())
		}
		return nil
	})

	// make the name a .txt for matching
	var nameFile string
	nameFile = name + ext

	fmt.Println("seek profile ", nameFile)
	fmt.Println("list of txt files: ", extList)

	// looks for our file
	for i, file := range extList {
		if file == nameFile {
			fmt.Printf("Exists: ")
			fmt.Printf(extList[i] + "\n")
			return true
		}
	}

	return false
}

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

func checkIfAppListing() bool {
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

func createAppListing() {

	fmt.Print("CREATING APP DIRECTORY")

	rootToApps := "/Volumes/Macintosh HD/Applications"
	var appList string
	ext := ".app"

	// GET ALL APPS
	filepath.Walk(rootToApps, func(getRoot string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(getRoot) == ext {
			app := fileInfo.Name()
			name := app[0 : len(app)-len(ext)]
			appList += name + ","
			// fmt.Printf(appnames)
			// fmt.Print(name)

		}
		return nil
	})

	// SAVE TO CSV
	newFile, err := os.Create("appDir.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	// io.Copy(newFile, strings.NewReader(appList))

	n, err := newFile.WriteString(appList)
	if err != nil {
		log.Fatal("whoops", err)
	}

	fmt.Println("bytes written: ", n)

}

func getApplications(appnames string) {

	// GETTING THE USER OR ANY PART OF FILEPATH
	// get current filepath
	// _, b, _, _ := runtime.Caller(0)
	// d := path.Join(path.Dir(b))
	// fp := filepath.Dir(d)
	// // fmt.Printf(fp, "\n")

	// var slashes string
	// var getRoot string
	// slash := "/"

	// // get the root directory
	// for _, item := range fp {
	// 	if len(slashes) < 3 {
	// 		// find root by counting between first two slashes
	// 		if string(item) == slash {
	// 			slashes += slash
	// 		}
	// 		getRoot += string(item)
	// 	}
	// }

	rootToApps := "/Volumes/Macintosh HD/Applications/"
	// fmt.Printf("root to apps: ", string(rootToApps)+"\n")
	ext := ".app"

	appDir, err := ioutil.ReadFile("appDir.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Split on comma.
	result := strings.Split(appnames, ",")
	result2 := strings.Split(string(appDir), ",")

	fmt.Print("APPNAMES: ", appnames)

	// OPTIMIZE FOR FASTER SEARCHING
	for _, item := range result {
		// fmt.Print(item, "\n")
		for _, item2 := range result2 {
			// fmt.Print(item2, "\n")

			if string(item) == string(item2) {
				fmt.Print(" WE HAVE A WINNER !!!!!!! ")
				fmt.Print(string(item) + "\n")
				rootToApp := rootToApps + item + ext
				fmt.Printf(rootToApp)
				err := open.Run(string(rootToApp))
				fmt.Print(err)
			}
		}
	}

	// fmt.Printf("appList: ", appList)
	return
}

func main() {

	// figure out if adding to a profile or opening a profile
	profile := flag.String("profile", "profile-name", "profile name")
	open := flag.Bool("open", true, "open this profile")
	edit := flag.Bool("edit", false, "do you want to modify a file")
	flag.Parse()

	fmt.Println("Desktop Profiles: " + *profile)
	fmt.Println("open: ", *open)
	fmt.Println("edit: ", *edit)

	ext := ".txt"

	// seekProfile(ext, *profile)
	// openBrowser("https://www.google.com")

	if seekProfile(ext, *profile) {

		fmt.Println("Accessing file...")
		// open or edit
		if *open {
			// access file, open them
			data := readFile(*profile + ext)
			// fmt.Printf(data)

			// scanner reads file line by line
			scanner := bufio.NewScanner(strings.NewReader(data))
			var applications string

			// CREATE APP DIRECTORY
			if checkIfAppListing() {
				fmt.Print("DIR EXISTS")
			} else {
				createAppListing()
				fmt.Print("MAKE THE DIR")
			}

			// loop through lines, determine if it's a website or an application
			for scanner.Scan() {
				// fmt.Println(scanner.Text())
				line := scanner.Text()
				lineID := line[:3]
				site := "htt"
				app := "app"
				// print(line, "\n")
				if lineID == site {
					// FOR BROWSER WEBSITE
					fmt.Printf("Opening browser: ", line)
					// openBrowser(line)
				} else if lineID == app {
					// FOR DESKTOP APPLICATIONS
					// fmt.Printf(lineID)
					// getting all the existing apps and sorting through at one time easier than pulling the app names every time
					lineComma := line[3:] + ","
					applications += lineComma
				}

			}

			// openApplications()
			// fmt.Sprint(applications)
			// getApplications(strings.TrimSuffix(applications, ","))
			testingOpen()

		}
	}
	// 	if *edit {
	// 		// read contents, write to it
	// 	}

	// } else {
	// 	// create
	// 	var create string
	// 	fmt.Println("This profile DNE, would you like to make one? (y/n)")
	// 	fmt.Scan(&create)

	// 	if create == "y" {
	// 		// createFile()
	// 		// editFile()
	// 	}
	// }

}
