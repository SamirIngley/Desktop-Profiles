# Desktop Profiles 
<p align="left">
  <a>
    <a href="https://goreportcard.com/badge/github.com/SamirIngley/Desktop-Profiles" />
    <img alt="goreportcard" src="https://goreportcard.com/badge/github.com/SamirIngley/Desktop-Profiles" target="_blank" />
    <a href="https://img.shields.io/github/downloads/SamirIngley/Desktop-Profiles/total"/>
  </a> 
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

#### :anger: Problem:
> I got tired of individually clicking through urls and apps to open what I need each time I switch to a new project or task. 
> I also did not want a cluttered desktop with everything open all the time, as this can be a large drain on the computer's resources and my attention span. 

#### :potable_water: Solution: 
> With this command line tool, easily create and customize profiles (groups) of apps and/or urls and open any of these profiles with one short command.

#### :apple: Example:
1. Create a profile called "read"
2. Add the Books app, your Notes app, and Merriam-Webster's website
3. Anytime you want to read, open the "read" profile and voila, there they are, as you last left them!

#### :1234: Instructions below
Note: This was built in Go on a Mac (for a Mac)

![Image1](READMEimg/gopherIMG.png)
************************************************************************************************

# :floppy_disk: Install:

### 0. Install Go and configure your GOPATH
* [Quick instructions](https://medium.com/@jimkang/install-go-on-mac-with-homebrew-5fa421fc55f5)
* [Detailed instructions](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go)

### 1. Download / clone this repo. 
* Move it where you want, then 'cd' into it. 
* Note: If you move the folder after the next step, refer to the :boom:Errors section at the bottom.

![Image2](READMEimg/download.png)


### 2. Run `go build` 
* This will create a binary file (executable) called desk. 
* Although this is a binary, you must be in the directory to run the program, `go install` will not make this global. If you know of a way for a golang program to store a string input and access it globally (through bash_profile?) please let me know!

### 3. Run `desk` 
* This will walk you through creating your app directory file. Follow the image and steps below. 

![Image3](READMEimg/download2.png)

* 3 of the most common paths for locations of Applications on Mac are shown there at the bottom between the short lines, like in the picture, copy and paste those -> don't forget to change {USER-NAME} to your computer user name, then type `done`
* **Once you complete this step, you're done! Checkout all the commands below**


# :mega: Commands:
:exclamation: You must be inside the directory to run the commands
![Image1](READMEimg/using.png)
<br>
##### :bulb: If you didn't create an executable replace `desk` with `go run desk.go` followed by the command

#### > List available profiles: `desk` 

#### > List input options: `desk -help`

#### > List profile contents: `desk -pf profile-name -l y` 

#### > Open a profile:  `desk -pf profile-name` 

#### > Add url(s): `desk -pf profile-name -add url`
Enter the url, and hit Enter after each one, type "done" when you've finished.

#### > Add app(s): `desk -pf profile-name -add app`
Enter the exact app name (ei: VS Code won't work but Visual Studio Code will), and hit Enter after each one, type "done" when you've finished.

#### > Create a profile (or add to an existing profile):  `desk -pf profile-name -add app`
Type the name of the profile you want to create and the app or url you'd like to add. Replace `app` with `url` if you want to add urls. Add each app or url, hit Enter after typing each one, then type `done` and hit Enter when you're finished.

#### > Delete from profile:  `desk -pf profile-name -del url`
(Same as adding except use -del instead of -add) or `go run desk.go -pf profile-name -del url`. Replace `url` with `app` if you want to add apps. Add each app or url, hit Enter after typing each one, then type `done` and hit Enter when you're finished.

#### > Delete profile:  `desk -pf profile-name -del profile` 
or `go run desk.go -pf profile-name -del profile-name` (after -del write "profile" or the name of the profile)

# :goal_net: Example:

An example profile has been provided in the profiles/ folder.
All your profiles can be found in this folder as well. 

To run the example, type:
`go run desk.go -pf example`

### To add an app (Slack) to the example, type:
`desk -pf example -add app`
Then type:
`Slack`
And hit Enter.

Type the names of any other apps you'd like in this profile and hit Enter after each one. 

Then type:
`done`
and hit Enter when you've finished. 

You should see a message saying you added an app to your profile.
To verify what was added, type:
`desk -pf example -l y` 
and a list should show up of the contents of the example profile.


### To add a website (Google) to the example, type:
`desk -pf example -add url`
and hit Enter.

Now, type:
`google.com`
And hit Enter.

You've now added "google.com" to the profile example.
Type the name of any other websites you'd like and hit enter after each one. 

And finally:
`done`
when you're finished. 

You should see a message saying you added a url(s) to your profile.
To verify what was added, type:
`desk -pf example -l y` 
and a list should show up of the contents of the example profile.

### To delete an app or url, do the same as above for adding, except use the `-del` flag instead of `-add`

### To delete a profile:
`desk -pf profile-name -del profile`

You should get a confirmation message asking if you're sure you want to delete the profile. 


## :warning: IMPORTANT:

### Input: 

* No trailing spaces when adding or deleting apps. Must be typed exactly as is in the appDir.txt

* Type anything for yes, type "no" for no, more details can be found about the input by typing the "-help" flag: `desk -help` (Exception: for -del flag when deleting a profile -> must be the profile name or the word "profile")

* Currently case sensitive - apps must be typed exactly as shown on your pc

* If you're having trouble specifying an app, find it in the appDir.txt file (which is created when you first run `go run desk.go`) and ignore the number in front of it")

### APP DIRECTORY:

* If you added more new apps to your pc, just delete the appDir.txt, .env, and "desk" files, **NOT** "desk.go", and a new ones will be created for you next time you run the program.
* Same applies if you want to add new paths to other Application folders

## Future updates:
- open specific "file" with "app" 
- needs to handle trailing space when deleting apps
- need to be able to close apps and urls too
- DRY for reading profile
- instead of appending to profile, look for blank line!


## :boom: Errors

If you run into ANY problems with opening apps, you've downloaded new apps, added incorrect paths while setting up, or you move your folder:

- First try to find the app in the "appDir.txt" file (ignore the number in front of it). This file is a collection of all the apps at all the paths you passed in. If it's not there, then you likely haven't passed the path to that app during setup and need to redo this step - see next step for instructions on how to do this. Your profiles will still exist afterwards.
- To Redo Setup: Delete the "appDir.txt" file, delete the ".env" file, and if you did `go build` - also delete the binary file "desk" but **NOT** "desk.go".
- If you do delete "desk.go" you will have to re download this repo or just that file.
- Re run the program with the install commands:`go run desk.go` or `go build` & `desk` 

From there you'll be prompted to set up your application paths again :) follow the pictures carefully.

If you're having issues with the binary file:
- Try typing `go run desk.go` everywhere instead of `desk` 

This was built on a Mac and for a Mac. 

Feel free to reach out to me if you run into any issues: samir.ingle7@gmail.com

## Contact:
* samir.ingle7@gmail.com
* https://www.samiringle.com

### Acknowledgements:
 
[![Go Report Card](https://goreportcard.com/badge/github.com/SamirIngley/Desktop-Profiles)](https://goreportcard.com/report/github.com/SamirIngley/Desktop-Profiles)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Open functionality help from:
https://github.com/skratchdot/open-golang

go doc
https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/DocsDeploy
