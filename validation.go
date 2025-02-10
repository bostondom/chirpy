package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Body string `json="body"`
	}
	var er struct{
		Error  string `json="error"`
	}
	var success struct {
		Valid bool `json:"valid"`
	}	
	err := json.NewDecoder(r.Body).Decode(&input)
	if err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-type","application/json")
		er.Error = "Something went wrong"
		dat,err := json.Marshal(er)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
	}
		w.Write(dat)
		return
	}
	if len(input.Body) >140 {
			w.WriteHeader(400)
		w.Header().Set("Content-type","application/json")
		er.Error = "Chirp is too long"
		dat,err := json.Marshal(er)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
	}
		w.Write(dat)
		return
	}
	w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type","application/json")
		success.Valid = true
		dat,err := json.Marshal(success)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
	}
		w.Write(dat)
}