go doc
https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/DocsDeploy

[![Go Report Card](https://goreportcard.com/badge/github.com/SamirIngley/Desktop-Profiles)](https://goreportcard.com/report/github.com/SamirIngley/Desktop-Profiles)


# Commands:

* list profiles: go run main.go 

* open a profile:  go run main.go -pf profile-name 
* list profile contents: go run main.go -pf profile-name

* create or add:  go run main.go -pf profile-name -add app:app-name
* delete app or ul:  go run main.go -pf profile-name -del url:website.com

* to delete profile:  go run main.go -pf -del profile-name


# Example:

an example profile has been provided in the profiles/ folder
all your profiles will be in this folder as well. 

To run the example, type:
go run main.go -pf example

To add an app (Slack) to the example, type:
go run main.go -pf example -add app:Slack

To add a website (Google) to the example, type:
go run main.go -pf example -add url:google.com

Same for deleting, except use the -del flag

To CREATE YOUR OWN PROFILE called "work" with gmail, type:
go run main.go -pf work -add url:gmail.com

To Delete a profile:
go run main.go -pf work -del work

You should get a confirmation message after making changes to any profile. 


## IMPORTANT:

### PROFILES: 
> Profiles are .txt files. No spaces in the profile, empty lines are fine
> For urls do not include https://www.

### INPUT:
> Type anything for yes, "no" for no flag, more details can be found about the input by typing the "-help" flag: go run main.go -help
> Currently case sensitive - apps must be typed exactly as shown in appDir / on your pc
>If you're having trouble specifying an app, Find it in the appDir.txt file and ignore the number in front of it")

### APP DIRECTORY:
>If you added more new apps to your pc, delete the appDir file and a new one will be created for you next time you run the program.



Open functionality help from:
https://github.com/skratchdot/open-golang

