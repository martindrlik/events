package local

import (
	"fmt"
	"path"
)

func FileName(id int64) string {
	return path.Join(Path, "events", fmt.Sprintf("%v", id))
}
