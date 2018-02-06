package remote

import (
	"bytes"
	"fmt"
	"net/http"
)

type Store struct {
	URL string
}

func (s Store) Store(id int64, p []byte) error {
	b := bytes.NewBuffer(p)
	url := fmt.Sprintf("%s/%d", s.URL, id)
	resp, err := http.DefaultClient.Post(url, "", b)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("store/remote: failed to post request: responded with http status code %v", resp.StatusCode)
	}
	return nil
}
