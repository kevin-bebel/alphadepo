package main

import (
	"encoding/json"
	"fmt"
	"github.com/asdine/storm"
	cjson "github.com/asdine/storm/codec/json"
	"gopkg.in/src-d/go-git.v4"
	_ "gopkg.in/src-d/go-git.v4/plumbing"
	_ "gopkg.in/src-d/go-git.v4/plumbing/transport"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

const (
	dbName = ""
)

func handlePushEvent(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		var event pushEvent
		json.Unmarshal(body, &event)
		//results = append(results, string(body))
		//fmt.Printlin(event)
		fmt.Fprint(w, event)
	} else {

	}

	//w.Header().Set("Content-Type", "text/plain")
	//w.Write([]byte("This is an example server controller file.\n"))
}

func handleCreateEvent(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "X-GitHub-Delivery")
	w.Header().Add("Access-Control-Allow-Headers", "X-GitHub-Event")
	if req.Method == "POST" {
		Info.Println("Entering Create Event Handler")
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		var event createEvent
		json.Unmarshal(body, &event)
		if event.RefType == "tag" {
			Info.Println("Create Event is a Tag")
			Info.Println("event : ", event)
			err = createNewVersion(event)
			if checkError(err, w) == true {
				return
			}

			//Checkout the branch associated to this tag
			//Upload the branch to S3.
			//Notify Of a Tag update
			Info.Println("Version Successfully Created")
			Info.Println("Repo URL : ", event.Repository.HTMLURL)
			repoUrl := event.Repository.HTMLURL
			_, err := git.PlainClone("/tmp/"+event.Repository.Name, false, &git.CloneOptions{
				URL: repoUrl,
			})

			if checkError(err, w) == true {
				return
			}
			Info.Println("Repository Cloned Successfully")
			Info.Println("Checking out branch : ", event.Ref)
			cmd := exec.Command("git", "checkout", "--track", "origin/"+event.Ref)
			cmd.Dir = "/tmp/" + event.Repository.Name
			_, err := cmd.Output()

			if checkError(err, w) == true {
				return
			}

			Info.Println("Branch Created Successfully")
			Info.Println("Uploading website to S3 Alpha Bucket")

			w.WriteHeader(http.StatusOK)
		} else {
			//Skip if not a tag and send back status ok
			w.WriteHeader(http.StatusOK)
		}
	}
}

func checkError(err error, w http.ResponseWriter) bool {
	if err != nil {
		Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return true
	}
	return false
}

func createNewVersion(event createEvent) error {
	Info.Println("Entering create new version methed ")
	version := Version{
		App: event.Repository.Name,
		Tag: event.Ref,
		//CreateDate: time.Now().Local().Format("2006-01-02"),
		CreatedAt: time.Now(),
	}

	Info.Println("version : ", version)
	Info.Println("Connecting to BoltDB")
	db, err := storm.Open("depo.db", storm.Codec(cjson.Codec))
	if err != nil {
		return err
	}

	defer db.Close()
	Info.Println("Inserting version into DB")
	db.Save(&version)
	Info.Println("Version Inserted Sucessfully")
	return nil
}

/*
	-- TODO --
	- Install MongoDB for later use
	- Create Event Should handle the creation of a tag and when it create store that version because a branch of that version will be create
	- handle PushEvent will check if the branch pushed was pushed to the tagged branch and if it update the alpha website.
	- Run a notification on Slack/Email of the update the has happened and its version number
*/
