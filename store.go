package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
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
	ch := make(chan error)
	for _, s := range stores {
		go func(s Storer) { ch <- s.Store(newID, p) }(s)
	}
	for i := 0; i < len(stores); i++ {
		if err := <-ch; err != nil {
			log.Println(err)
		}
	}
}
