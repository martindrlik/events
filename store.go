package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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
	globalID++
	for _, store := range stores {
		b := bytes.NewBuffer(p)
		if err := store.Store(globalID, b); err != nil {
			log.Println(err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}
}
