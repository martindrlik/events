package remote

import (
	"fmt"
	"io"
	"net/http"
)

type Store struct {
	URL string
}

func (s Store) Store(id int64, r io.Reader) error {
	url := fmt.Sprintf("%s/%d", s.URL, id)
	resp, err := http.DefaultClient.Post(url, "", r)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("store/remote: failed to post request: responded with http status code %v", resp.StatusCode)
	}
	return nil
}
