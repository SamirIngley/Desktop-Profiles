# Desk Profile

App & website launcher
Create any number of profiles to open any combination of your frequently used apps and websites from the cli at once! 

** Follow the instructions below to get started ** 

Note: This was built in Go on a Mac (for a Mac)

[![Go Report Card](https://goreportcard.com/badge/github.com/SamirIngley/Desktop-Profiles)](https://goreportcard.com/report/github.com/SamirIngley/Desktop-Profiles)

![Image1](gopherIMG.png)
************************************************************************************************

## :floppy_disk: Install:

1. Download / clone this repo

2. Run `go run desk.go` -> this will walk you through creating your app directory file

3. You're ready to roll! Checkout the commands and the example below


## :mega: Commands:

* When you first download the package, run `go run desk.go`, this will give you instructions on providing your Applications folders

* list available profiles: `go run desk.go` 

* open a profile:  `go run desk.go -pf profile-name` 
* list profile contents: `go run desk.go -pf profile-name`

* create or add:  `go run desk.go -pf profile-name -add app:app-name`
* delete app or ul:  `go run desk.go -pf profile-name -del url:website.com`

* to delete profile:  `go run desk.go -pf -del profile-name`


## :goal_net: Example:

An example profile has been provided in the profiles/ folder
all your profiles will be in this folder as well. 

To run the example, type:
`go run desk.go -pf example`

To add an app (Slack) to the example, type:
`go run desk.go -pf example -add app:Slack`

To add a website (Google) to the example, type:
`go run desk.go -pf example -add url:google.com`

Same for deleting, except use the `-del` flag

To CREATE YOUR OWN PROFILE called "work" with gmail, type:
`go run desk.go -pf work -add url:gmail.com`

To Delete a profile:
`go run desk.go -pf work -del work`

You should get a confirmation message after making changes to any profile. 


## :warning: IMPORTANT:

### Input: 

* No trailing spaces when adding or deleting apps. Must be typed exactly as is in the appDir.txt

* Type anything for yes, type "no" for no, more details can be found about the input by typing the "-help" flag: go run desk.go -help (Exception: for -del flag when deleting a profile -> must be the profile name or "profile")

* Currently case sensitive - apps must be typed exactly as shown on your pc

* If you're having trouble specifying an app, find it in the appDir.txt file (which is created when you first run `go run desk.go`) and ignore the number in front of it")

### APP DIRECTORY:

* If you added more new apps to your pc, just delete the appDir file and a new one will be created for you next time you run the program.


Future updates:
- needs to handle trailing space when deleting apps
- need to be able to close apps and urls too



Open functionality help from:
https://github.com/skratchdot/open-golang

go doc
https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/DocsDeploy
