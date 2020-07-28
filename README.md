# Desktop Profiles

App & website launcher
Create any number of profiles to open any combination of your frequently used apps and websites from the cli at once! 

** Follow the instructions below to get started ** 

Note: This was built in Go on a Mac (for a Mac)

[![Go Report Card](https://goreportcard.com/badge/github.com/SamirIngley/Desktop-Profiles)](https://goreportcard.com/report/github.com/SamirIngley/Desktop-Profiles)

![Image1](gopherIMG.png)
************************************************************************************************

## :floppy_disk: Install:

0. Install go and configure your GOPATH - here are some great [instructions](https://medium.com/@jimkang/install-go-on-mac-with-homebrew-5fa421fc55f5)

1. Download / clone this repo. Then 'cd' into it.

2. Run `go run desk.go` -> this will walk you through creating your app directory file

3. You're ready to roll! Checkout the commands and walk through the example below

4. For a list of the available flags, type `go run desk.go -help`

## :mega: Commands:

* When you download the package, run `go run desk.go`, this will give you instructions on providing the routes to the Applications folders you want to be able to access using this app
    - 3 of the most common routes for locations of Applications on Mac are shown at the bottom, just copy and paste those -> don't forget to change {USER-NAME} to your computer user name, then type `done`
    - Don't worry, if you find out you want to add more paths to Applications folders later - just delete the appDir file and run `go run desk.go`, now re-add all the paths you'd like to include

* list of input options: `go run desk.go -help`

* list available profiles: `go run desk.go` 

* list profile contents: `go run desk.go -pf profile-name -l y -o n`

* open a profile:  `go run desk.go -pf profile-name` 

* to create a profile, type the name of the profile and the app or url you'd like to add:  `go run desk.go -pf profile-name -add app`. Type `done` when you've finished adding. If you want to add a url, change `app` to `url`. 

* to add to an existing profile: `go run desk.go -pf profile-name -add app:app-name`

* delete app or url:  `go run desk.go -pf profile-name -del url` and type the names one by one hitting Enter after each. Type `done` when finished

* to delete profile:  `go run desk.go -pf -del profile-name`


## :goal_net: Example:

An example profile has been provided in the profiles/ folder.
All your profiles can be found in this folder as well. 

To run the example, type:
`go run desk.go -pf example`

### To add an app (Slack) to the example, type:
`go run desk.go -pf example -add app`
Then type:
`Slack`
And hit Enter.

You've now added Slack to your profile, type the name of any other apps you'd like in this profile and hit enter after each one. 

Then type:
`done`
and hit Enter when you've finished. 

### To add a website (Google) to the example, type:
`go run desk.go -pf example -add url`
and hit Enter.

Now, type:
`google.com`
And hit Enter.

You've now added google.com, type the name of any other websites you'd like and hit enter after each one. 

And finally:
`done`
when you're finished. 

### To delete an app or url, do the same as above for adding, except use the `-del` flag instead of `-add`

### To delete a profile:
`go run desk.go -pf profile-name -del work`

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
