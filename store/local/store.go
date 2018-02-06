package local

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

var Path string

func Store(id int64, p []byte) error {
	name := path.Join(Path, "events", fmt.Sprintf("%v", id))
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

type FileInfo struct {
	ModTime time.Time
	Name    string
	Size    int64
}

func Ls() ([]FileInfo, error) {
	name := path.Join(Path, "events")
	dir, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("store/local: open directory %s failed: %v", name, err)
	}
	files := make([]FileInfo, 0, 3000)
	for i := 0; i < 6; i++ {
		infos, err := dir.Readdir(500)
		if err != nil {
			return nil, fmt.Errorf("store/local: read directory %s failed: %v", name, err)
		}
		for _, info := range infos {
			files = append(files, FileInfo{
				ModTime: info.ModTime(),
				Name:    info.Name(),
				Size:    info.Size()})
		}
		if len(infos) < 500 {
			break
		}
	}
	return files, nil
}
