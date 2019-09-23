package main

import (
	"log"
	"os/exec"
	"net/http"

	erpc "github.com/Varunram/essentials/rpc"
	utils "github.com/Varunram/essentials/utils"
)

func triggerScript() {
	http.HandleFunc("/setup", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckPost(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		org := r.FormValue("org")                 // the name of the org that the user selects
		platform := r.FormValue("platform")       // the name of the platform that the user wants to use
		emailentity := r.FormValue("emailentity") // entity to send emails as
		abl := r.FormValue("abl")                 // add other blockchain handlers
		invo := r.FormValue("invo")               // investor options
		recpo := r.FormValue("recpo")             // recipient options

		if org == "" {
			org = "org"
		}

		if len(org) > 2 {
			org = org[0] // avoid deadling with parsing multi word GitHub imports
		}

		if platform == "" {
			platform = "plat"
		}

		if len(platform) > 2 {
			platform = platform[0]
		}

		if emailentity == "" {
			emailentity = "The Plat"
		}

		if abl == "" {
			abl = "y"
		}

		if invo == "" {
			invo = "y"
		}

		if recpo == "" {
			recpo = "y"
		}

		cmd, err := exec.Command("./gen.sh", org, platform, invo, recpo, emailentity, abl).Output()
		if err != nil {
			log.Fatal(err)
		}

		erpc.MarshalSend(w, string(cmd))
	})
}

func StartServer(portx int) {
	triggerScript()
	port, err := utils.ToString(portx)
	if err != nil {
		log.Fatal("Port not string")
	}

	log.Println("Starting RPC Server on Port: ", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
