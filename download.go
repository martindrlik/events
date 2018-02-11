package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/martindrlik/events/store/local"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	to, err := strconv.Atoi(r.FormValue("to"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ids := make([]int64, 0, 1+(to-from))
	for i := from; i <= to; i++ {
		ids = append(ids, int64(i))
	}
	err = local.Zip(ids, w)
	if err != nil {
		log.Println(err)
		http.Error(w, "creating zip to download failed", http.StatusInternalServerError)
	}
}
