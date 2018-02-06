package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/martindrlik/events/store/local"
)

func lsHandler(w http.ResponseWriter, r *http.Request) {
	files, err := local.Ls()
	if err != nil {
		log.Println(err)
		http.Error(w, "listing local store files failed", http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(files)
	if err != nil {
		log.Println(err)
		http.Error(w, "encoding files to json failed", http.StatusInternalServerError)
		return
	}
}
