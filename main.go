package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	//_ := mux.NewRouter()
	log.SetOutput(os.Stdout)
	http.HandleFunc("/", handlePushEvent)
	http.HandleFunc("/create", handleCreateEvent)
	err := http.ListenAndServe(os.Getenv("IP")+":"+os.Getenv("PORT"), nil)
	log.Println("About to listen on", os.Getenv("PORT"), ". Go to https:", os.Getenv("IP")+":"+os.Getenv("PORT"))
	log.Fatal(err)
}
