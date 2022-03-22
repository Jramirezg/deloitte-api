package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gihub.com/deloitte-api/pkg/replace"
)

var retrieved_string string

func ReturnStringReplaced(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: string")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error when retrieving string ", err)
	}
	err = json.Unmarshal(reqBody, &retrieved_string)
	if err != nil {
		log.Fatal("Error unmarshalling string ", err)
	}
	log.Println("String received: ", retrieved_string)
	new_string := replace.Replace(retrieved_string)
	log.Println("New string ", new_string)

	json.NewEncoder(w).Encode(new_string)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to test API!")
	fmt.Println("Endpoint Hit: homePage")
}
