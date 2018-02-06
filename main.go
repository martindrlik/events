package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/martindrlik/events/store/local"
)

var (
	addr = flag.String("addr", ":8080", "listen and server on addr TCP network address")
	home = flag.String("home", "", "events home directory, $EVENTSPATH or d is default")
)

var stores = map[string]interface {
	Store(id int64, r io.Reader) error
}{}

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
