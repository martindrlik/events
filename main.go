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
)

func init() {
	flag.StringVar(&local.Path, "path", "", "path is events home directory, $EVENTSPATH or d is default")
}

type storer interface {
	Store(id int64, p []byte) error
}

var (
	remotes = map[string]storer{}
)

func main() {
	flag.Parse()
	if local.Path == "" {
		local.Path = os.Getenv("EVENTSPATH")
	}
	if local.Path == "" {
		local.Path = "d"
	}
	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/store", storeHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
