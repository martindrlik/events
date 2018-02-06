package local

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
)

type Store struct {
	Home string
}

func (s Store) Store(id int64, p []byte) error {
	name := path.Join(s.Home, "events", fmt.Sprintf("%v", id))
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("store/local: create file failed: %v", err)
	}
	defer f.Close()
	b := bytes.NewBuffer(p)
	_, err = io.Copy(f, b)
	if err != nil {
		return fmt.Errorf("store/local: write to file failed: %v", err)
	}
	return nil
}
