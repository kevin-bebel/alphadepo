package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	//_ := mux.NewRouter()
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	http.HandleFunc("/", handlePushEvent)
	http.HandleFunc("/create", handleCreateEvent)
	Info.Println("About to listen on", os.Getenv("PORT"), ". Go to https:", os.Getenv("IP")+":"+os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("IP")+":"+os.Getenv("PORT"), nil)
	log.Fatal(err)
}
