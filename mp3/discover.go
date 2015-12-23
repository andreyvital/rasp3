package mp3

import (
	"log"
	"os"

	"github.com/CentaurWarchief/rasp3/fs"
)

func Discover(l Library, root string) {
	for _, file := range fs.Readdir(root) {
		if !IsMp3File(file) {
			continue
		}

		f, err := os.OpenFile(file, os.O_RDONLY, 0666)

		if err != nil {
			log.Println(err)
			continue
		}

		defer f.Close()

		fi, err := f.Stat()

		if err != nil {
			log.Println(err)
			continue
		}

		size := int(fi.Size())

		l.Add(NewWithoutID3(file, size))
	}
}
