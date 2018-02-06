package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/martindrlik/events/store/local"
)

var (
	addr = flag.String("addr", ":8080", "listen and server on addr TCP network address")
	home = flag.String("home", "", "events home directory, $EVENTSPATH or d is default")
)

type Storer interface {
	Store(id int64, p []byte) error
}

var stores = map[string]Storer{}

func main() {
	flag.Parse()
	if *home == "" {
		*home = os.Getenv("EVENTSPATH")
	}
	if *home == "" {
		*home = "d"
	}
	stores[""] = local.Store{Home: *home}
	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/store", storeHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
