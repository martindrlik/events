package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/martindrlik/events/store/remote"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	type DeleteRemote struct {
		Key string
	}
	type UpserRemote struct {
		Key string
		URL string
	}
	type Config struct {
		DeleteRemote DeleteRemote
		UpserRemote  UpserRemote
	}
	var config Config
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&config)
	if err != nil {
		log.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	if config.DeleteRemote.Key != "" {
		delete(remotes, config.DeleteRemote.Key)
	}
	if config.UpserRemote.Key != "" {
		remotes[config.UpserRemote.Key] = remote.Store{URL: config.UpserRemote.URL}
	}
}
