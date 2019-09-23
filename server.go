package main

import (
	"log"
	"net/http"
	"os/exec"
	"archive/zip"
	"io/ioutil"
	"os"
	"strings"

	erpc "github.com/Varunram/essentials/rpc"
	utils "github.com/Varunram/essentials/utils"
)

func ZipWriter(orgName string) error {
	baseFolder, err := os.Getwd()
	if err != nil {
		return err
	}

	// Get a Buffer to Write To
	outFile, err := os.Create(baseFolder + "/temp.zip")
	if err != nil {
		return err
	}
	defer outFile.Close()

	baseFolder += "/" + orgName + "/"

	w := zip.NewWriter(outFile)
	err = addFiles(w, baseFolder, "")
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) error {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			data, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				return err
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				return err
			}

			_, err = f.Write(data)
			if err != nil {
				return err
			}
		} else {
			newBase := basePath + file.Name() + "/"
			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
	return nil
}

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

		if len(strings.Split(org, " ")) > 1 {
			org = strings.Split(org, " ")[0] // avoid deadling with parsing multi word GitHub imports
		}

		if platform == "" {
			platform = "plat"
		}

		if len(strings.Split(platform, " ")) > 1 {
			platform = strings.Split(platform, " ")[0] // avoid deadling with parsing multi word GitHub imports
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

		_, err = exec.Command("./gen.sh", org, platform, invo, recpo, emailentity, abl).Output()
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		err = ZipWriter(org)
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		http.ServeFile(w, r, "temp.zip")
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
