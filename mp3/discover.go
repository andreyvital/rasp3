package mp3

import (
	"strconv"
	"strings"

	"github.com/CentaurWarchief/rasp3/fs"
	"github.com/mikkyang/id3-go"
	"github.com/mikkyang/id3-go/v2"
)

func Discover(l Library, root string) {
	for _, file := range fs.Readdir(root) {
		if !IsMp3File(file) {
			continue
		}

		f, err := id3.Open(file)

		size := f.Size()

		if err != nil {
			l.Add(NewWithoutId3(file, size))
			continue
		}

		if f.Title() == "" || f.Artist() == "" || f.Album() == "" {
			l.Add(NewWithoutId3(file, size))
			continue
		}

		year, _ := strconv.Atoi(f.Year())

		var track int

		if frame := f.Frame(TrackFrame); frame != nil {
			track, _ = strconv.Atoi(frame.(*v2.TextFrame).Text())
		}

		l.Add(
			New(file, size, &ID3{
				Title:  trim(f.Title()),
				Artist: trim(f.Artist()),
				Album:  trim(f.Album()),
				Year:   year,
				Track:  track,
			}),
		)
	}
}

func trim(s string) string {
	return strings.TrimSpace(
		strings.Replace(s, "\u0000", "", -1),
	)
}
