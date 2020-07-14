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

func testingOpen() {
	// WORKS !!!
	err := open.Run("/Volumes/Macintosh HD/Applications/Atom.app")
	fmt.Print(err)
}

func seekProfile(ext string, name string) bool {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	extList := []string{}
	path = path + "/profiles"
	// fmt.Print(path)

	// finds all txt files in current directory
	filepath.Walk(path, func(path string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(path) == ext {
			extList = append(extList, fileInfo.Name())
		}
		return nil
	})

	// makes the name a .txt for matching
	var nameFile string
	nameFile = name + ext

	// fmt.Println("seeking profile ", nameFile, "...")
	// fmt.Println("list of txt files: ", extList)

	// looks for our file
	for i, file := range extList {
		if file == nameFile {
			fmt.Printf("Found " + extList[i] + "\n")
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

func checkIfAppDir() bool {
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

func createAppDir() {

	rootToApps := "/Volumes/Macintosh HD/Applications"
	var appList string
	ext := ".app"

	// GET ALL APPS
	filepath.Walk(rootToApps, func(getRoot string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(getRoot) == ext {
			app := fileInfo.Name()
			name := app[0 : len(app)-len(ext)]
			appList += name + "\n"
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

	// appDir, err := ioutil.ReadFile("appDir.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Split on comma.
	result := strings.Split(appnames, ",")

	apps := readFile("appDir.txt")
	scanner := bufio.NewScanner(strings.NewReader(apps))

	// result2 := strings.Split(string(appDir), ",")

	// fmt.Print("APPNAMES: ", appnames, "\n")

	// NEED TO OPTIMIZE FOR EFFICIENT SEARCHING
	for scanner.Scan() {
		// LOOP - does anything in app dir = any of the app names in profile ??
		for _, item := range result {
			// fmt.Print(item, "\n")

			// fmt.Print(item2, "\n")
			// fmt.Print(item, scanner.Text()+"\n")

			if string(item) == string(scanner.Text()) {
				// fmt.Print(" WE HAVE A WINNER !!!!!!! \n")
				fmt.Print("Opening " + string(item) + "\n")

				rootToApp := rootToApps + item + ext
				// fmt.Printf(rootToApp)
				err := open.Run(string(rootToApp))
				if err != nil {
					fmt.Print(err)
				}

			}
		}
	}

	// fmt.Printf("appList: ", appList)
	return
}

func writeToFile(file string, content string, ext string) {
	fileLoc := "profiles/" + file + ext
	// data := readFile("profiles/" + *profile + ext)
	addMe := "\n" + content + "\n"
	// fmt.Print("ADDING file: ", file, addMe)

	// OPEN AND ADD TO FILE
	currentFile, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	if _, err := currentFile.WriteString(addMe); err != nil {
		log.Println(err)
	}
	defer currentFile.Close()
	fmt.Println("Added ", content, " to ", file)
}

func main() {

	// figure out if adding to a profile or opening a profile
	// if profile exists, will be added to profile
	// if profile dne, will be created and added to
	pf := flag.String("pf", "profile-name", "profile name")
	o := flag.String("o", "yes", "open this profile")
	add := flag.String("add", "no", "creates or adds <app:app-name> or <url:url-address> to profile")
	del := flag.String("del", "no", "deletes <app:app-name> or <url:url-address> from profile")
	flag.Parse()

	fmt.Println("Desktop Profiles: " + *pf)
	fmt.Println("open: ", *o)
	fmt.Println("add: ", *add)
	fmt.Println("del: ", *del)

	ext := ".txt"

	// seekProfile(ext, *profile)
	// openBrowser("https://www.google.com")

	// IF PROFILE EXISTS
	if seekProfile(ext, *pf) {
		fmt.Println("Accessing file...")

		// OPEN IT
		if *o == "yes" {

			// access file, open them
			data := readFile("profiles/" + *pf + ext)
			// fmt.Printf(data)

			// scanner reads file line by line
			scanner := bufio.NewScanner(strings.NewReader(data))
			var applications string

			// CREATE APP DIRECTORY if needed
			if !(checkIfAppDir()) {
				fmt.Print("CREATING APP DIRECTORY (this may take a minute)...")
				createAppDir()
			}

			// GO THROUGH PROFILE, separate urls and apps, THEN OPEN THEM
			for scanner.Scan() {
				// fmt.Println(scanner.Text())
				line := scanner.Text()
				site := "url:"
				app := "app:"
				// print(line, "\n")
				if line != "" {
					lineID := line[:4]
					if lineID == site {
						// OPEN BROWSER WEBSITE
						fmt.Print("Opening browser ", line, "\n")
						openBrowser(line[4:])
					} else if lineID == app {
						// GET DESKTOP APPLICATIONS
						// fmt.Printf(lineID)
						// getting all the existing apps and sorting through at one time easier than pulling the app names every time
						lineComma := line[4:] + ","
						applications += lineComma
					}
				}
			}
			// OPEN APPLICATIONS
			getApplications(strings.TrimSuffix(applications, ","))

		}

	}

	// ADDS TO PROFILE, CREATES PROFILE IF NEEDED
	if !(*add == "no") {
		fileLoc := "profiles/" + *pf + ext
		// data := readFile("profiles/" + *profile + ext)
		addMe := "\n" + *add + "\n"
		// fmt.Print("ADDING file: ", file, addMe)

		// OPEN AND ADD TO FILE
		currentFile, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		if _, err := currentFile.WriteString(addMe); err != nil {
			log.Println(err)
		}
		defer currentFile.Close()
		fmt.Println("Added ", *add, " to ", *pf)

	}

	if !(*del == "no" || *del == "profile") {
		fmt.Print("not working yet")
		// fileLoc := "profiles/" + *profile + ext

		// fileData, err := ioutil.ReadFile(fileLoc)
		// if err != nil {
		// 	log.Fatal(err)
		// }

	}

	if *del == "profile" {
		fileLoc := "profiles/" + *pf + ext

		var yn string
		fmt.Print("Are you sure you want to delete ", *pf, " profile? [y/n]..")
		fmt.Scanln(&yn)
		if yn == "y" || yn == "yes" {
			var err = os.Remove(fileLoc)
			if err != nil {
				fmt.Println("Error deleting file.")
				log.Fatal(err)
			}

			fmt.Println("Profile Deleted")

		}
	}

}
