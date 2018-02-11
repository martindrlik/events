package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/martindrlik/events/store/local"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	type Download struct {
		IDs []int64
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	downloadRequest := Download{}
	err := decoder.Decode(&downloadRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = local.Zip(downloadRequest.IDs, w)
	if err != nil {
		log.Println(err)
		http.Error(w, "download: something went wrong", http.StatusInternalServerError)
	}
}
