# Desktop Profiles

Create profiles to open any combination of your frequently used apps and websites from the cli at once! 

Follow the instructions below to get started

[![Go Report Card](https://goreportcard.com/badge/github.com/SamirIngley/Desktop-Profiles)](https://goreportcard.com/report/github.com/SamirIngley/Desktop-Profiles)

************************************************************************************************

## :floppy_disk: Install:

1. Download / clone this repo

2. In the main.go file, there are *two* places where routes to your Applications folders are specified. You will need to change the `{USER-NAME}` part in both places to whatever your the name of your User is on your pc. The locations are specified below. If you don't know or want a list of users, in the cli type in `ls /users`. 

    * Location 1: In main.go, line 112, at the end of rootToAppsUser, replace {USER-NAME} with your user    ( inside func createAppDir() )
    * Location 2: In main.go, line 186, at the end of rootToAppsUSER, replace {USER-NAME} with your user    ( inside func getApplications() )


3. Run `go run main.go` -> this will create your app directory file

4. You're ready to roll! Checkout the commands and the example below


## :mega: Commands:

* list profiles: `go run main.go` 

* open a profile:  `go run main.go -pf profile-name` 
* list profile contents: `go run main.go -pf profile-name`

* create or add:  `go run main.go -pf profile-name -add app:app-name`
* delete app or ul:  `go run main.go -pf profile-name -del url:website.com`

* to delete profile:  `go run main.go -pf -del profile-name`


## :goal_net: Example:

An example profile has been provided in the profiles/ folder
all your profiles will be in this folder as well. 

To run the example, type:
`go run main.go -pf example`

To add an app (Slack) to the example, type:
`go run main.go -pf example -add app:Slack`

To add a website (Google) to the example, type:
`go run main.go -pf example -add url:google.com`

Same for deleting, except use the `-del` flag

To CREATE YOUR OWN PROFILE called "work" with gmail, type:
`go run main.go -pf work -add url:gmail.com`

To Delete a profile:
`go run main.go -pf work -del work`

You should get a confirmation message after making changes to any profile. 


## :warning: IMPORTANT:

### PROFILES: 
* Profiles are .txt files. No spaces in the profile, empty lines are fine

* For urls do not include "https://www."

### INPUT:
* Type anything for yes, type "no" for no, more details can be found about the input by typing the "-help" flag: go run main.go -help (Exception: for -del flag when deleting a profile -> must be the profile name or "profile")

* Currently case sensitive - apps must be typed exactly as shown on your pc

* If you're having trouble specifying an app, find it in the appDir.txt file (which is being created now) and ignore the number in front of it")

### APP DIRECTORY:
* If you added more new apps to your pc, delete the appDir file and a new one will be created for you next time you run the program.



Open functionality help from:
https://github.com/skratchdot/open-golang

go doc
https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/DocsDeploy