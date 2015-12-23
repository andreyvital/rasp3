package mp3

import (
	"strconv"

	"github.com/mikkyang/id3-go"
	"github.com/mikkyang/id3-go/v2"
	"mbios.io/mbsound/fs"
)

func Discover(l Library, root string) {
	for _, file := range fs.Readdir(root) {
		if !IsMp3File(file) {
			continue
		}

		f, err := id3.Open(file)

		if err != nil {
			l.Add(NewWithoutId3(file))
			continue
		}

		if f.Title() == "" || f.Artist() == "" || f.Album() == "" {
			l.Add(NewWithoutId3(file))
			continue
		}

		year, _ := strconv.Atoi(f.Year())

		var track int

		if frame := f.Frame(TrackFrame); frame != nil {
			track, _ = strconv.Atoi(frame.(*v2.TextFrame).Text())
		}

		l.Add(
			New(file, &ID3{
				Title:  f.Title(),
				Artist: f.Artist(),
				Album:  f.Album(),
				Year:   year,
				Track:  track,
			}),
		)
	}
}
