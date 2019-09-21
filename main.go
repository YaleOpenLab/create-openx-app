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

	if platformName == "" || orgName == "" {
		log.Println("platform name or org name empty, quitting")
		os.Exit(1)
	}

	// trigger the gen script
	cmd, err := exec.Command("./gen.sh", orgName, platformName).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("replaced names", string(cmd))
}
