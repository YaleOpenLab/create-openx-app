package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	// "bytes"
	scan "github.com/Varunram/essentials/scan"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Listen bool `short:"l" description:"Start the generator script in rpc mode"`
	Port   int  `short:"p" description:"The port on which the server runs on" default:"8081"`
}

func main() {
	fmt.Printf("%s", "Welcome to create-openx-app\n\n")
	fmt.Printf("%s", "All options are yes by default. Please press n/N for omitting specific options\n\n")

	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if opts.Listen {
		StartServer(opts.Port)
	}

	fmt.Printf("%s", "Enter organization name: ")
	orgName, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", "Enter platform name: ")
	platformName, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", "Enter the entity that you would like to send emails as: ")
	emailName, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", "This template uses Stellar as the main blockchain platform but support"+
		"for other blockchains will be added in the future. Do you want to add the other blockchain handlers? ")

	otherBlHandlers, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	if otherBlHandlers == "n" || otherBlHandlers == "N" {
		otherBlHandlers = "n"
	} else {
		otherBlHandlers = "y"
	}

	fmt.Printf("%s", "Would you like to have voting options for investors? ")
	invVote, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	if invVote == "n" || invVote == "N" {
		invVote = "n"
	} else {
		invVote = "y"
	}

	fmt.Printf("%s", "Would you like to have additional options for recipients? ")
	recpVote, err := scan.ScanString()
	if err != nil {
		log.Fatal(err)
	}

	if recpVote == "n" || recpVote == "N" {
		recpVote = "n"
	} else {
		recpVote = "y"
	}

	if platformName == "" || orgName == "" {
		fmt.Printf("%s", "platform name or org name empty, quitting")
		os.Exit(1)
	}

	fmt.Printf("%s", "\n\nRUNNING GENERATOR SCRIPT..\n\n")
	// trigger the gen script
	cmd, err := exec.Command("./gen.sh", orgName, platformName, invVote, recpVote, emailName, otherBlHandlers).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(cmd))
}
