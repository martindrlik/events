package local

import (
	"fmt"
	"io"
	"os"
	"path"
)

type Store struct {
	Home string
}

func (s Store) Store(id int64, r io.Reader) error {
	name := path.Join(s.Home, "events", fmt.Sprintf("%v", id))
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("store/local: create file failed: %v", err)
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	if err != nil {
		return fmt.Errorf("store/local: write to file failed: %v", err)
	}
	return nil
}
