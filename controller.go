package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		var event createEvent
		json.Unmarshal(body, &event)
		//results = append(results, string(body))
		//fmt.Printlin(event)
		//fmt.Fprint(w, event)
		fmt.Fprint(w, event.RefType)
	} else {

	}
}

/*
	-- TODO --
	- Install MongoDB for later use
	- Create Event Should handle the creation of a tag and when it create store that version because a branch of that version will be create
	- handle PushEvent will check if the branch pushed was pushed to the tagged branch and if it update the alpha website.
	- Run a notification on Slack/Email of the update the has happened and its version number
*/
