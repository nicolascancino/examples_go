package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handle struct {
}

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func HandleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "its my api")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application.json")
	w.Write(response)
	//fmt.Fprintf(w, "Payload %v\n", user)
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
