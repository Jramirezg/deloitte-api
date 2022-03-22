package main

import (
	"log"
	"net/http"
	"strconv"

	"gihub.com/deloitte-api/pkg/api"
	"github.com/gorilla/mux"
)

const (
	port = 8080
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", api.HomePage)
	myRouter.HandleFunc("/string", api.ReturnStringReplaced).Methods("POST")
	// myRouter.HandleFunc("/string/{id}", api.returnSingleString)
	str_port := ":" + strconv.Itoa(port)
	log.Fatal(http.ListenAndServe(str_port, myRouter))
}

func main() {
	log.Println("Running API in port ", port)
	handleRequests()
}
