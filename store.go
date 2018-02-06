package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/martindrlik/events/store/local"
)

var globalID int64

func storeHandler(w http.ResponseWriter, r *http.Request) {
	readBody := func() ([]byte, error) {
		p := make([]byte, 5000)
		n, err := r.Body.Read(p)
		if err != io.EOF && err != nil {
			return nil, err
		}
		if n <= 0 {
			return nil, errors.New("read 0 bytes from request body")
		}
		if n >= len(p) {
			return nil, fmt.Errorf("read %d bytes from request body, it is too much", n)
		}
		return p, nil
	}
	defer r.Body.Close()
	p, err := readBody()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newID := atomic.AddInt64(&globalID, 1)
	err = local.Store(newID, p)
	if err != nil {
		log.Println(err)
		http.Error(w, "local store failed", http.StatusInternalServerError)
		return
	}
	go storeRemotely(newID, p)
}

func storeRemotely(id int64, p []byte) {
	ch := make(chan error)
	for _, s := range remotes {
		go func(s storer) { ch <- s.Store(id, p) }(s)
	}
	for i := 0; i < len(remotes); i++ {
		if err := <-ch; err != nil {
			log.Println(err)
		}
	}
}
