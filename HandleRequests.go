package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//HTTP requests
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc(
		"/api/v1/balance/{id}/{currency}", GetBalance).Methods("GET")
	myRouter.HandleFunc(
		"/api/v1/time/balance/{id}/{beginTime}/{endTime}", GetListOfEventsByTime).Methods("GET")
	myRouter.HandleFunc(
		"/api/v1/post/enrollment", CreateNewEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
