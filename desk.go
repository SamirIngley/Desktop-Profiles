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
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/skratchdot/open-golang/open"
)

// GLOBAL VAR FOR LOCATION OF PROFILES AND APPDIR
var dir string
var profpath string
var apps string

func seekProfile(ext string, name string) bool {

	extList := []string{}
	// apps := readFile("appDir.txt")
	// var apps string
	// scan through appDir
	// scanner := bufio.NewScanner(strings.NewReader(apps))

	// var profpath string
	// i := 0
	// for scanner.Scan() {
	// 	if i == 0 {
	// 		profpath = scanner.Text()
	// 		// fmt.Println("profpath ", profpath)
	// 		i++
	// 	} else if i == 1 {
	// 		apps = readFile(scanner.Text())
	// 		i++
	// 	} else {
	// 		break
	// 	}
	// }

	// finds all txt files in current directory
	filepath.Walk(profpath, func(profpath string, fileInfo os.FileInfo, _ error) error {
		// fmt.Println(fileInfo.Name(), filepath.Ext(profpath))
		if filepath.Ext(profpath) == ext {
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
	exists := false
	extList := []string{}

	// if app dir exists, return true
	if len(dir) > 0 && dir != "reset" {
		filepath.Walk(dir, func(dir string, fileInfo os.FileInfo, _ error) error {
			// fmt.Println(fileInfo.Name())
			if fileInfo.Name() == "appDir.txt" {
				extList = append(extList, fileInfo.Name())
				exists = true
			}
			return nil
		})
	} else {
		createAppDir()
	}

	return exists
}

func createAppDir() {

	// THIS FUNCTION CREATES YOUR APP DIRECTORY FILE WHICH IS NEEDED FOR THIS PROGRAM TO RUN
	// **** YOUR rootToApps SHOULD BE THE PATH TO YOUR APPLICATIONS FOLDER *****
	// The commented out portion below can help you get the path to your applications folder.
	// If you know the path you can change rootToApps to the path here, and in the
	// getApplications function
	fmt.Println(" ")
	fmt.Println("""
	________               ______  _____                        ________                _____________ ______              
	___  __ \_____ ___________  /____  /_______ ________        ___  __ \______________ ___  __/___(_)___  /_____ ________
	__  / / /_  _ \__  ___/__  //_/_  __/_  __ \___  __ \       __  /_/ /__  ___/_  __ \__  /_  __  / __  / _  _ \__  ___/
	_  /_/ / /  __/_(__  ) _  ,<   / /_  / /_/ /__  /_/ /       _  ____/ _  /    / /_/ /_  __/  _  /  _  /  /  __/_(__  ) 
	/_____/  \___/ /____/  /_/|_|  \__/  \____/ _  .___/        /_/      /_/     \____/ /_/     /_/   /_/   \___/ /____/  
												/_/                                                                       
	""")
	fmt.Println(" ")

	fmt.Println("We know the default locations of the Apps on your Mac, this step gives you the chance to modify those defaults if you've changed them,")
	fmt.Println("added more folders, or have moved them to different locations. Follow these instructions to create your App Directory (appDir.txt) ")
	fmt.Println("If you haven't made any changes - follow the instructions and don't add anything else. If you're not sure, you can always delete the")
	fmt.Println("appDir.txt file and another will be created for you next time you run the program.")
	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println(">>>>>>>>>> INSTRUCTIONS ~~ ~~ ")
	fmt.Println(" ")

	fmt.Println("**** Below are the 3 default paths to locations of app directories on a Mac *****")
	fmt.Println(" ")

	fmt.Println("1.  Copy & Paste the paths below (/Volumes...)")
	fmt.Println(" ")

	fmt.Println("2.  Change {USER-NAME} to your computer User name and hit Enter")
	fmt.Println(" ")

	fmt.Println("3.  These should be enough locations for most Macs. If you have other locations, feel free to add other paths as well")
	fmt.Println(" ")

	fmt.Println("4.  When you've finished adding paths, type 'done' on a new line and hit Enter")
	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println("*  If later you find out you have more/different locations, just delete the appDir.txt file and re-run this program - you will get this same setup guide again.")
	fmt.Println("*  These paths are locations to your: Mac system apps, Mac Hard Disk apps, Mac User apps")
	fmt.Println("----------------------------")
	fmt.Println("/Volumes/Macintosh HD/System/Applications")
	fmt.Println("/Volumes/Macintosh HD/Applications")
	fmt.Println("/Volumes/Macintosh HD/Users/{USER-NAME}/Applications")
	fmt.Println("----------------------------")

	var dirsArray []string
	var appList string

	ext := ".app"

	// SET GLOBAL PATHS IN ENVIRONMENT !! ----------------------------------------------------------
	// fmt.Println("setting env")

	// GET DIR PATH
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// CREATE Hardoated PATHS
	Hdir := path
	Hprofpath := Hdir + "/profiles"
	Happs := Hdir + "/appDir.txt"

	// WRITING TO ENV FILE
	stuff := "DIR=" + Hdir + "\n" + "PROFPATH=" + Hprofpath + "\n" + "APPS=" + Happs
	env, err1 := godotenv.Unmarshal(stuff)
	err1 = godotenv.Write(env, ".env")

	if err1 != nil {
		fmt.Println(err1)
	}
	// ASSIGN ENV VARS TO LOCAL VARS
	dir = os.Getenv("DIR")
	profpath = os.Getenv("PROFPATH")
	apps = os.Getenv("APPS")
	// ADDING PROFILE PATH AT LINE 0 and APP DIR PATH AT LINE 1
	// get permanent path to profiles (append -> first line of appList, next few lines will the the paths)
	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dir = path
	// profpath = dir + "/profiles"
	// apps = dir + "/appDir.txt"

	// fmt.Println("create app dir DIR ", dir)
	// fmt.Println("PROF ", profpath)
	// fmt.Println("APPS ", apps)

	// appList += profpath
	// appList += apps
	// fmt.Print("PROF PATH", profpath)

	// ADDS DIRS FROM os.STDIN to DIRS ARRAY and APPLIST
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "done" {
			fmt.Println("Creating your App Directory... this may take a minute, please wait")
			break
		} else {
			dirsArray = append(dirsArray, line)
			appList += "\n" + line
		}
	}

	// fmt.Println("Dirs Array: ", dirsArray)

	// STORE LOCATIONS IN FIRST FEW LINES OF APPDIR
	// INDEX OF THESE LOCATIONS - THAT NUMBER IS SELECTED PLACED IN FRONT OF THE APP SO WE KNOW WHERE TO FIND IT
	// THE 0 LOCATION IS RESERVED FOR THE PATH TO THE PROFILES, 1 LOCATION IS RESERVED FOR APP DIRECTORY
	i := 0
	for _, item := range dirsArray {
		// for each directory, walk the path and grab all the apps there, assign them your number
		path := item
		// fmt.Println("ITEM ITEM ITEM: ", item, i)
		filepath.Walk(path, func(getRoot string, fileInfo os.FileInfo, _ error) error {
			if filepath.Ext(getRoot) == ext {
				app := fileInfo.Name()
				name := app[0 : len(app)-len(ext)]
				iStr := strconv.Itoa(i)
				strFormat := string("\n" + iStr + name)
				// fmt.Println("strformat", strFormat)
				appList += strFormat

				// fmt.Printf(appnames)
				// fmt.Print(name)

			}
			return nil
		})
		i++
	}

	// SAVE TO CSV
	newFile, err := os.Create("appDir.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	// fmt.Println(appList)
	// io.Copy(newFile, strings.NewReader(appList))

	n, err := newFile.WriteString(appList)
	if err != nil {
		log.Fatal("whoops", err)
	}

	fmt.Println("App Directory created >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("bytes written: ", n)
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("SUCCESS: You're ready to begin creating profiles! It's easy!")
	fmt.Println("For a detailed guide, visit: https://github.com/SamirIngley/Desktop-Profiles")
	fmt.Println("type 'go run desk.go -help' for a list of options")
	fmt.Println(" ")
	fmt.Println("Important Info about using the app: ")
	fmt.Println("- If you encounter an error with this step and your apps won't load, find the correct paths to your Applications folders, delete the appDir.txt & .env files, re-run this program, and specify those App directories as well as or instead of the default ones supplied")
	fmt.Println("- Apps are case sensitive, apps must be typed EXACTLY as shown on your pc / in the app directory")
	fmt.Println("- If you're having trouble specifying an app, look for it in the appDir.txt file (which will be created after this step) and ignore the number in front of it when typing it in")
	fmt.Println("- If you add more new apps to your pc, delete the appDir.txt and .env file and new ones will be created for you next time you run the program.")
	fmt.Println(" ")
}

func getApplications(appnames string) {

	// GO THROUGH THE APP LIST ONCE, COMPARE THE APPNAMES TO EACH ITEM IN THE APPLIST
	// Once you find the app, grab the number in front of it, and go in order through the routes
	// at the top of appdir to find the correct route

	// Split input appnames on comma.
	appSplit := strings.Split(appnames, ",")
	var appsWithID []string
	ext := ".app"
	// apps := readFile("appDir.txt")
	// fmt.Println(apps)

	// scan through appDir
	scanner := bufio.NewScanner(strings.NewReader(readFile(apps)))

	// FINDS THE APPS TO OPEN IN THE APPDIR ->  ADDS THEIR IDS TO APPSWITHID
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		for _, item := range appSplit {

			// fmt.Print(item, " ... ", string(scanner.Text()), "\n")
			// if item from appSplit
			if scanner.Text() != "" {
				if string(item) == string(scanner.Text()[1:]) {
					// rootID := string(scanner.Text()[:1])
					// rootID := app[0]
					// fmt.Println("ROOTID: ", rootID)
					// fmt.Print(" WE HAVE A WINNER !!!!!!! \n")
					// fmt.Print("\n" + "Opening " + string(item) + "\n")
					appsWithID = append(appsWithID, scanner.Text()) // FINDS THE APP WITH ITS ID
				}
			}
		}
	}

	// TAKES ALL APPSWITHIDS, LOADS THEM FROM THEIR ID
	// WE COMPARE THE APP ID WITH THE CORRESPONDING PATH NUMBERS
	for _, item := range appsWithID {
		// WE START WITH 0 (but there will never be a 0 id app bc that's reserved for profiles path)
		i := 0
		rootint := item[:1]
		rootID, _ := strconv.Atoi(rootint)
		itemAsStr := string(item[1:])
		scanner := bufio.NewScanner(strings.NewReader(readFile(apps)))
		// for each app, scan through the appdir to the appropriate line
		for scanner.Scan() {
			line := scanner.Text()
			// fmt.Println(line)
			// if the line number matches the rootID of the app
			if line != "" {
				// either we have a match, or the counter goes up by 1 and we move to the next line in the appDir
				// we check our counter with the app id root id. 1 is the first path, 2 is the second.
				// We intentionally skip 0 - that's reserved for profiles
				if i == rootID { // COUNTER == APP ID
					// we're at the right line so open the app with this line
					rootToApp := line + "/" + itemAsStr + ext
					// fmt.Println("ROOT TO APP: ", rootToApp)
					// fmt.Printf(rootToApp)
					// fmt.Print("Opening " + string(item) + "\n")
					err := open.Run(string(rootToApp))
					if err != nil {
						fmt.Print(err)
					}
					break

					// } else if i < len(appSplit) {
					// 	i++
				} else {
					i++
				}
			}
		}

	}
}

func writeToFile(file string, content string, ext string) {
	// // CREATING A NEW PROFILE or ADDING TO EXISTING PROFILE
	// // NOT USED IN MAIN,

	// fileLoc := "profiles/" + file + ext
	// // data := readFile("profiles/" + *profile + ext)
	// addMe := "\n" + content + "\n"
	// // fmt.Print("ADDING file: ", file, addMe)

	// // OPEN AND ADD TO FILE
	// currentFile, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }
	// if _, err := currentFile.WriteString(addMe); err != nil {
	// 	log.Println(err)
	// }
	// defer currentFile.Close()
	// fmt.Println("Added ", content, " to ", file)
}

func main() {

	// figure out if adding to a profile or opening a profile
	// if profile exists, will be added to profile
	// if profile dne, will be created and added to
	yo := flag.String("help", "GO HERE", "https://github.com/SamirIngley/Desktop-Profiles")
	pf := flag.String("pf", "profile-name", "name of the profile")
	l := flag.String("l", "no", "list contents of the profile")
	o := flag.String("o", "yes", "open this profile")
	add := flag.String("add", "no", "specify 'app' or 'url' to create new profile or add to an existing profile")
	del := flag.String("del", "no", "deletes profile if 'profile' or name of profile is typed, otherwise deletes app or url entered")
	flag.Parse()

	// FLAG STATUSES
	// fmt.Println("list: ", *l)
	// fmt.Println("open: ", *o)
	// fmt.Println("add: ", *add)
	// fmt.Println("del: ", *del)

	ext := ".txt"

	// seekProfile(ext, *profile)
	// openBrowser("https://www.google.com")
	if *yo == "me" {
		fmt.Println("Here's the github: https://github.com/SamirIngley/Desktop-Profiles")
	}

	// fmt.Println("Checking if appdir")
	// fmt.Println("primaryDIR ", os.Getenv("DIR"))
	// fmt.Println("PROF ", os.Getenv("PROFPATH"))
	// fmt.Println("APPS ", os.Getenv("APPS"))

	// CREATE APP DIRECTORY --------------------------------------------------------- appDir text File
	e := godotenv.Load() // looks for .env file
	if e != nil {        // if dne, create app dir -> creates .env
		// fmt.Print("creating env")
		createAppDir()
		// fmt.Println("1")
	} else { // if it does exist, assign locals from env
		dir = os.Getenv("DIR")
		profpath = os.Getenv("PROFPATH")
		apps = os.Getenv("APPS")
		// fmt.Println("2")
		if !(checkIfAppDir()) { // if dotenv exists, but app dir doesn't
			createAppDir()
		}
	}

	dir = os.Getenv("DIR")
	profpath = os.Getenv("PROFPATH")
	apps = os.Getenv("APPS")
	// fmt.Println("yeetDIR ", dir)
	// fmt.Println("PROF ", profpath)
	// fmt.Println("APPS ", apps)

	// SHOW ALL AVAILABLE PROFILES
	if *pf == "profile-name" && *o == "yes" && *l == "no" && *add == "no" && *del == "no" {
		// path, err := os.Getwd()
		// if err != nil {
		// log.Fatal(err)
		// }

		extList := []string{}
		// path = path + "/profiles"
		// fmt.Print("PATH ", path)

		// apps := readFile("appDir.txt")

		// scan through appDir
		// scanner := bufio.NewScanner(strings.NewReader(apps))

		// var profpath string
		// i := 0
		// for scanner.Scan() {
		// 	if i == 0 {
		// 		profpath = scanner.Text()
		// 		i++
		// 	} else {
		// 		break
		// 	}
		// }
		// fmt.Println("prof path in show available profiles", profpath)
		// fmt.Println(dir)
		// fmt.Println(profpath)
		// fmt.Println(apps)

		// finds all txt files in directory
		filepath.Walk(profpath, func(profpath string, fileInfo os.FileInfo, _ error) error {
			// fmt.Println(profpath)
			if filepath.Ext(profpath) == ext {
				extList = append(extList, fileInfo.Name())
			}
			return nil
		})

		fmt.Println("Profiles: ")
		fmt.Println(extList)

		return
	}

	// CHECKS IF THE PROFILE EXISTS AND OPENS IT---------------------------------- if no other flags are called
	if seekProfile(ext, *pf) {
		// fmt.Println("Accessing file...")

		// LISTS CURRENT PROFILE CONTENTS----------------------------------------------------------
		if *l != "no" {
			// access file, open them
			// apps := readFile("appDir.txt")
			// scan through appDir
			// scanner := bufio.NewScanner(strings.NewReader(apps))
			// var profpathplus string
			// i := 0
			// for scanner.Scan() {
			// 	if i == 0 {
			// 		profpathplus = scanner.Text() + "/" + *pf + ext
			// 		i++
			// 	} else {
			// 		break
			// 	}
			// }
			profpathplus := profpath + "/" + *pf + ext
			data := readFile(profpathplus)
			fmt.Println(data)
		}

		// OPENS PROFILE -------------------------------------------------------------------------------------
		if *o == "yes" && *l == "no" && *add == "no" && *del == "no" {
			fmt.Println("Opening " + *pf)

			// access file, open them
			// apps := readFile("appDir.txt")
			// scan through appDir
			// scanner := bufio.NewScanner(strings.NewReader(apps))
			// var profpathplus string
			// i := 0
			// for scanner.Scan() {
			// 	if i == 0 {
			// 		profpathplus = scanner.Text() + "/" + *pf + ext
			// 		i++
			// 	} else {
			// 		break
			// 	}
			// }
			// fmt.Println("OPENNING")
			// fmt.Println("DIR ", dir)
			// fmt.Println("PROF ", profpath)
			// fmt.Println("APPS ", apps)

			profpathplus := profpath + "/" + *pf + ext

			// access file, open them
			data := readFile(profpathplus)

			// scanner reads file line by line
			scanner := bufio.NewScanner(strings.NewReader(data))
			var applications string

			//GO THROUGH PROFILE --------------------------------------- SEPARATE APPS AND URLS then OPEN EACH
			for scanner.Scan() {
				// fmt.Println(scanner.Text())
				line := scanner.Text()
				site := "url:"
				app := "app:"
				// print(line, "\n")
				if line != "" {
					// we need the line ID, so - nothing less than 5 characters allowed in the file
					lineID := line[:4]
					if lineID == site {
						// OPEN WEBSITES ------------------------------------------------------------------
						// fmt.Print("Opening browser ", line[4:], "\n")
						lineNoURLTag := line[4:]
						if lineNoURLTag[:4] == "http" {
							openBrowser(lineNoURLTag)
						} else {
							openBrowser("https://www." + lineNoURLTag)
						}
					} else if lineID == app {
						// GET DESKTOP APPLICATIONS
						// getting all the existing apps and sorting through at one time easier than pulling the app names every time

						// fmt.Printf(lineID)
						lineComma := line[4:] + ","
						applications += lineComma
					}
				}
			}
			// OPEN APPLICATIONS--------------------------------------------------------------------------
			getApplications(strings.TrimSuffix(applications, ","))

		}

	}

	// CREATES PROFILE IF NEEDED or ADDS TO PROFILE ------------------------------------------------------

	if *add == "app" {

		// access file, open them
		// apps := readFile("appDir.txt")
		// scan through appDir
		// scanner := bufio.NewScanner(strings.NewReader(apps))
		// var fileLoc string
		// i := 0
		// for scanner.Scan() {
		// 	if i == 0 {
		// 		fileLoc = scanner.Text() + "/" + *pf + ext
		// 		i++
		// 	} else {
		// 		break
		// 	}
		// }
		fileLoc := profpath + "/" + *pf + ext

		// data := readFile("profiles/" + *profile + ext)
		// addMe := "\n" + *add + "\n"
		// fmt.Print("ADDING file: ", file, addMe)
		if seekProfile(ext, *pf) == false {
			fmt.Println("Creating new profile", *pf)
		}

		fmt.Println("Enter apps one by one, type 'done' when finished: ")

		var appsToCreate []string

		// OPEN AND ADD TO FILE
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "done" {
				break
			} else {
				appsToCreate = append(appsToCreate, line)
			}
		}
		// fmt.Println("Apps to create: ", appsToCreate)
		for _, item := range appsToCreate {
			addMe := "\n" + "app:" + item
			// fmt.Printf("Input was: %q\n", line)

			// Here we APPEND to file or CREATE a new file
			currentFile, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				log.Println(err)
			}
			if _, err := currentFile.WriteString(addMe); err != nil {
				log.Println(err)
			}

			defer currentFile.Close()
			// fmt.Println(item)
		}
		fmt.Println("Added to ", *pf)

	} else if *add == "url" {
		// access file, open them
		// apps := readFile("appDir.txt")
		// scan through appDir
		// scanner := bufio.NewScanner(strings.NewReader(apps))
		// var fileLoc string
		// i := 0
		// for scanner.Scan() {
		// 	if i == 0 {
		// 		fileLoc = scanner.Text() + "/" + *pf + ext
		// 		i++
		// 	} else {
		// 		break
		// 	}
		// }
		// data := readFile("profiles/" + *profile + ext)
		// addMe := "\n" + *add + "\n"
		// fmt.Print("ADDING file: ", file, addMe)
		fileLoc := profpath + "/" + *pf + ext

		if seekProfile(ext, *pf) == false {
			fmt.Println("Creating new profile ", *pf)
		}

		fmt.Println("Enter urls one by one, type 'done' when finished: ")

		var urlsToCreate []string

		// OPEN AND ADD TO FILE
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "done" {
				break
			} else {
				urlsToCreate = append(urlsToCreate, line)
			}
		}
		// fmt.Println("Urls to create: ", urlsToCreate)
		for _, item := range urlsToCreate {
			addMe := "\n" + "url:" + item
			// fmt.Printf("Input was: %q\n", line)

			// Here we APPEND to file or CREATE a new file
			currentFile, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				log.Println(err)
			}
			if _, err := currentFile.WriteString(addMe); err != nil {
				log.Println(err)
			}

			defer currentFile.Close()
			// fmt.Println(item)
		}
		fmt.Println("Added to ", *pf)

	} else if *add != "no" {
		fmt.Println("Try again - specify what you want to add '-add app' or '-add url'")
	}

	// DELETES SPECIFIC LINES (APP OR URL) FROM PROFILE --------------------------------------------------
	if *del == "app" {
		// fmt.Print("not working yet\n")

		fmt.Println("Enter apps one by one, type 'done' when finished: ")

		// OPEN AND ADD TO FILE

		// read name of app/url to be deleted
		var line2 string
		scanner2 := bufio.NewScanner(os.Stdin)
		// i := 0
		var appsToDel []string
		for scanner2.Scan() {

			line2 = scanner2.Text()
			if line2 == "done" {
				break
			} else {
				appsToDel = append(appsToDel, line2)
			}
		}

		// fmt.Println(appsToDel)
		// access file, open them
		// apps := readFile("appDir.txt")
		// scan through appDir
		// scanner := bufio.NewScanner(strings.NewReader(apps))
		// var profpathplus string
		// i := 0
		// for scanner.Scan() {
		// 	if i == 0 {
		// 		profpathplus = scanner.Text() + "/" + *pf + ext
		// 		i++
		// 	} else {
		// 		break
		// 	}
		// }
		profpathplus := profpath + "/" + *pf + ext

		if seekProfile(ext, *pf) {
			// access file, open it

			// fmt.Printf(data)
			// FOR EACH APP IN APPS TO DELETE, DO A SCANNER ON THE FILE AND COMPARE EACH LINE to the app
			for _, item := range appsToDel {
				found := false
				data := readFile(profpathplus)
				scanner := bufio.NewScanner(strings.NewReader(data))
				for scanner.Scan() {
					// fmt.Println(scanner.Text())
					line := scanner.Text()
					delMe := "app:" + item
					if line == delMe {
						// DELETE THE DATA aka REPLACE IT WITH NOTHING ------------------------------
						newContents := strings.Replace(data, string(delMe), "", -1)
						// fmt.Println("NEW: \n", newContents)
						err := ioutil.WriteFile(profpathplus, []byte(newContents), 0)
						if err != nil {
							panic(err)
						} else {
							found = true
							fmt.Println("Deleted ", item, " from ", *pf)
						}
					}
				}
				// loop through all the lines and if found still false, let us know
				if found == false {
					fmt.Println("Could not find ", item, " in ", *pf)
				}
			}

		} else {
			fmt.Println("Typo? We couldn't find that profile")
		}

	} else if *del == "url" {
		// fmt.Print("not working yet\n")
		fmt.Println("Enter urls one by one, type 'done' when finished: ")

		// OPEN AND ADD TO FILE
		var line2 string
		scanner2 := bufio.NewScanner(os.Stdin)
		// i := 0
		var urlsToDel []string
		for scanner2.Scan() {

			line2 = scanner2.Text()
			if line2 == "done" {
				break
			} else {
				urlsToDel = append(urlsToDel, line2)
			}
		}

		// fmt.Println(urlsToDel)

		if seekProfile(ext, *pf) {
			// access file, open it

			// fmt.Printf(data)
			// FOR EACH APP IN APPS TO DELETE, DO A SCANNER ON THE FILE AND COMPARE EACH LINE to the app
			// access file, open them
			// apps := readFile("appDir.txt")
			// scan through appDir
			// scanner := bufio.NewScanner(strings.NewReader(apps))
			// var profpathplus string
			// i := 0
			// for scanner.Scan() {
			// 	if i == 0 {
			// 		profpathplus = scanner.Text() + "/" + *pf + ext
			// 		i++
			// 	} else {
			// 		break
			// 	}
			// }
			profpathplus := profpath + "/" + *pf + ext

			for _, item := range urlsToDel {
				found := false
				data := readFile(profpathplus)
				scanner := bufio.NewScanner(strings.NewReader(data))
				for scanner.Scan() {
					// fmt.Println(scanner.Text())
					line := scanner.Text()
					delMe := "url:" + item
					if line == delMe {
						// DELETE THE DATA aka REPLACE IT WITH NOTHING ------------------------------
						newContents := strings.Replace(data, string(delMe), "", -1)
						// fmt.Println("NEW: \n", newContents)
						err := ioutil.WriteFile(profpathplus, []byte(newContents), 0)
						if err != nil {
							panic(err)
						} else {
							found = true
							fmt.Println("Deleted ", item, " from ", *pf)
						}
					}
				}
				// loop through all the lines and if found still false, let us know
				if found == false {
					fmt.Println("Could not find ", item, " in ", *pf, "make sure it's spelled correctly and has no trailing spaces")
				}
			}

		} else {
			fmt.Println("Typo? We couldn't find that profile")
		}

	} else if *del != "no" && *del != *pf && *del != "profile" {
		fmt.Println("Try again - specify what you want to delete '-del app' or '-del url'")
	}

	// DELETES THE WHOLE PROFILE --------------------------------------------------------------------------
	if *del == "profile" || *del == *pf {
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
