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
	for _, file := range extList {
		if file == nameFile {
			// fmt.Printf("Found " + extList[i] + "\n")
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

	// THIS FUNCTION CREATES YOUR APP DIRECTORY FILE WHICH IS NEEDED FOR THIS PROGRAM TO RUN
	// **** YOUR rootToApps SHOULD BE THE PATH TO YOUR APPLICATIONS FOLDER *****
	// The commented out portion below can help you get the path to your applications folder.
	// If you know the path you can change rootToApps to the path here, and in the
	// getApplications function

	var rootToAppsSYSTEM = ("/Volumes/Macintosh HD/System/Applications")
	var rootToAppsMACHD = ("/Volumes/Macintosh HD/Applications")
	var rootToAppsUSER = ("/Volumes/Macintosh HD/Users/samiringle/Applications")

	var appList string
	ext := ".app"

	// GET ALL APPS

	// loop through array doesn't work so manual
	// paths := [rootToAppsSYSTEM, rootToAppsMACHD, rootToAppsUSER]
	// fmt.Println("PATH LEN: ", paths.length)
	// for (i = 0; i < paths.length; i++) {

	filepath.Walk(rootToAppsSYSTEM, func(getRoot string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(getRoot) == ext {
			app := fileInfo.Name()
			name := app[0 : len(app)-len(ext)]
			appList += string("0" + name + "\n")

			// fmt.Printf(appnames)
			// fmt.Print(name)

		}
		return nil
	})

	filepath.Walk(rootToAppsMACHD, func(getRoot string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(getRoot) == ext {
			app := fileInfo.Name()
			name := app[0 : len(app)-len(ext)]
			appList += string("1" + name + "\n")

			// fmt.Printf(appnames)
			fmt.Print(name)

		}
		return nil
	})
	filepath.Walk(rootToAppsUSER, func(getRoot string, fileInfo os.FileInfo, _ error) error {
		if filepath.Ext(getRoot) == ext {
			app := fileInfo.Name()
			name := app[0 : len(app)-len(ext)]
			appList += string("2" + name + "\n")

			// fmt.Printf(appnames)
			fmt.Print(name)

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

	// THIS IS THE LOCATION FOR MY APPLICATIONS - THIS ACTUAL PATH IS NEEDED TO OPEN THE FILE
	rootToAppsSYSTEM := "/Volumes/Macintosh HD/System/Applications/"
	rootToAppsMACHD := "/Volumes/Macintosh HD/Applications/"
	rootToAppsUSER := "/Volumes/Macintosh HD/Users/{user-name}/Applications/"

	// paths := [rootToAppsSYSTEM, rootToAppsMACHD, rootToAppsUSER]
	// fmt.Printf("root to apps: ", string(rootToApps)+"\n")
	ext := ".app"

	// appDir, err := ioutil.ReadFile("appDir.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Split input appnames on comma.
	result := strings.Split(appnames, ",")

	apps := readFile("appDir.txt")
	scanner := bufio.NewScanner(strings.NewReader(apps))
	fmt.Println("OPENING APPS")

	// result2 := strings.Split(string(appDir), ",")

	fmt.Print("APPNAMES: ", appnames, "\n")

	// NEED TO OPTIMIZE FOR EFFICIENT SEARCHING
	for scanner.Scan() {
		// LOOP - does anything in app dir = any of the app names in profile ??
		for _, item := range result {
			// fmt.Print(item, "\n")

			// fmt.Print(item2, "\n")
			fmt.Print(item, " ... ", string(scanner.Text()), "\n")

			if string(item) == string(scanner.Text()[1:]) {
				rootID := string(scanner.Text()[:1])
				// rootID := app[0]
				fmt.Println("ROOTID: ", rootID)
				// fmt.Print(" WE HAVE A WINNER !!!!!!! \n")
				fmt.Print("Opening " + string(item) + "\n")

				if rootID == "0" {
					rootToApp := rootToAppsSYSTEM + item + ext
					fmt.Printf(rootToApp)
					fmt.Print("Opening " + string(item) + "\n")
					err := open.Run(string(rootToApp))
					if err != nil {
						fmt.Print(err)
					}
				}
				if rootID == "1" {
					rootToApp := rootToAppsMACHD + item + ext
					fmt.Printf(rootToApp)
					fmt.Print("Opening " + string(item) + "\n")
					err := open.Run(string(rootToApp))
					if err != nil {
						fmt.Print(err)
					}
				}
				if rootID == "2" {
					rootToApp := rootToAppsUSER + item + ext
					fmt.Printf(rootToApp)
					fmt.Print("Opening " + string(item) + "\n")
					err := open.Run(string(rootToApp))
					if err != nil {
						fmt.Print(err)
					}
				}

			}
		}
	}

	// fmt.Printf("appList: ", appList)
	return
}

func writeToFile(file string, content string, ext string) {
	// CREATING A NEW PROFILE or ADDING TO EXISTING PROFILE
	// NOT USED IN MAIN,

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
	add := flag.String("add", "no", "creates new profile or adds <app:app-name> or <url:url-address> to existing profile")
	del := flag.String("del", "no", "deletes profile if 'profile' is typed, otherwise deletes app or url entered")
	flag.Parse()

	// FLAG STATUSES
	// fmt.Println("open: ", *o)
	// fmt.Println("add: ", *add)
	// fmt.Println("del: ", *del)

	ext := ".txt"

	// seekProfile(ext, *profile)
	// openBrowser("https://www.google.com")

	// IF PROFILE EXISTS
	if seekProfile(ext, *pf) {
		fmt.Println("Accessing file...")

		// OPEN IT
		if *o == "yes" && *add == "no" && *del == "no" {
			fmt.Println("Opening " + *pf)

			// access file, open them
			data := readFile("profiles/" + *pf + ext)
			// fmt.Printf(data)

			// scanner reads file line by line
			scanner := bufio.NewScanner(strings.NewReader(data))
			var applications string

			// CREATE APP DIRECTORY
			if !(checkIfAppDir()) {
				fmt.Println("CREATING APP DIRECTORY")
				fmt.Println("This only happens the first time you run the program")
				fmt.Println("This may take a minute...")
				fmt.Println("Note: ")
				fmt.Println("** If you encounter an error with this step, or your apps won't load, you'll need to specify the path to your Applications folder")
				fmt.Println("** It's easy - instructions can be found in the README at https://www.github.com/SamirIngley/DesktopProfiles")
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
						fmt.Print("Opening browser ", line[4:], "\n")
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
