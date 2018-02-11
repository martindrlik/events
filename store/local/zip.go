package local

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
)

func Zip(ids []int64, w io.Writer) error {
	archive := zip.NewWriter(w)
	for _, id := range ids {
		f, err := archive.Create(fmt.Sprintf("%d", id))
		if err != nil {
			return err
		}
		b, err := ioutil.ReadFile(FileName(id))
		if err != nil {
			return err
		}
		_, err = f.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
