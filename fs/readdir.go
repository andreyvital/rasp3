package fs

import (
	"os"
	"path/filepath"
)

func Readdir(root string) (files []string) {
	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files
}
