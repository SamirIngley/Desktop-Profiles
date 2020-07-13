package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	fmt.Println("find: ", nameFile)
	fmt.Println("list of txt files: ", extList)

	// looks for our file
	for i, file := range extList {
		if file == nameFile {
			fmt.Printf("Exists: ")
			fmt.Printf(extList[i])
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

// func createProfile(name string, file string) {

// }

func main() {
	// figure out if adding to a profile or opening a profile
	fmt.Println("Desktop Profiles...")
	profile := flag.String("profile", "profile-name", "profile name")
	open := flag.Bool("open", true, "open this profile")
	edit := flag.Bool("edit", false, "do you want to modify a file")
	flag.Parse()

	fmt.Println(*profile)
	fmt.Println(*open)
	fmt.Println(*edit)

	ext := ".txt"

	// seekProfile(ext, *profile)

	if seekProfile(ext, *profile) {
		fmt.Println("Accessing file...")
		// open or edit
		if *open {
			// access file, open them
			data := readFile(*profile + ext)
			fmt.Printf(data)
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
