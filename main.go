package main

import (
	"log"
	"os"
	"os/exec"
	// "bytes"
	scan "github.com/Varunram/essentials/scan"
)

func main() {
	log.Println("welcome to create-openx-app")

	log.Println("Enter organization name: ")
	orgName, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Enter platform name: ")
	platformName, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Would you like to have additional options for investors? (press n/N for no)")
	invVote, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	if invVote == "n" || invVote == "N" {
		log.Println("you have requested for investors not to have the option to vote towards projects")
	} else {
		invVote = "investor"
	}

	log.Println("Would you like to have additional options for recipients? (press n/N for no)")
	recpVote, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	if recpVote == "n" || recpVote == "N" {
		log.Println("you have requested for recipients not to have the option to vote towards projects")
	} else {
		recpVote = "recipient"
	}

	if platformName == "" || orgName == "" {
		log.Println("platform name or org name empty, quitting")
		os.Exit(1)
	}

	// trigger the gen script
	cmd, err := exec.Command("./gen.sh", orgName, platformName, invVote, recpVote).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("replaced names", string(cmd))
}
