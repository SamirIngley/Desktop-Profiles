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
)

type Config struct {
	Profile struct {
		name  string `json: name`
		files string `json: files`
	}
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

func getApplications(appnames string) (string, string) {

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

	rootToApps := "/Applications/"
	// fmt.Printf("root to apps: ", string(rootToApps)+"\n")
	var appList string
	ext := ".app"

	filepath.Walk(rootToApps, func(getRoot string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(getRoot) == ext {
			app := fileInfo.Name()
			name := app[0 : len(app)-len(ext)]
			appList += name + ","

			// fmt.Printf(appnames)
			// fmt.Printf(name)
		}
		return nil
	})

	// Split on comma.
	result := strings.Split(appList, ",")
	result2 := strings.Split(appnames, ",")
	fmt.Printf("APPNAMES: ", appnames)

	for _, item := range result {
		for _, item2 := range result2 {
			if string(item) == string(item2) {
				fmt.Printf(" WE HAVE A WINNER !!!!!!! ")
				fmt.Printf(string(item) + "\n")
				// rootToApp := rootToApps + item + ".app"
				rootToApp2 := rootToApps + item
				var bingo, err = os.OpenFile(rootToApp2, os.O_RDWR, 0644)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer bingo.Close()
			}
		}
	}

	// fmt.Printf("appList: ", appList)
	return appList, rootToApps
}

// func openApplications(appnames string) {
// 	appList, root := getApplications()

// 	for _, item := range appnames {
// 		for _, item2 := range appList {
// 			// fmt.Printf(string(item), "  ", item2)
// 			if string(item) == string(item2) {
// 				os.Open(root)
// 			}
// 		}
// 	}
// }

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

			// loop through lines, determine if it's a website or an application
			for scanner.Scan() {
				// fmt.Println(scanner.Text())
				line := scanner.Text()
				lineID := line[:3]
				site := "htt"
				app := "app"
				// print(line, "\n")
				if lineID == site {
					// openBrowser(line)
				} else if lineID == app {
					// fmt.Printf(lineID)
					// getting all the existing apps and sorting through at one time easier than pulling the app names every time
					lineComma := line[3:] + ","
					applications += lineComma
				}

			}
			// openApplications()
			fmt.Sprint(applications)
			getApplications(strings.TrimSuffix(applications, ","))

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
