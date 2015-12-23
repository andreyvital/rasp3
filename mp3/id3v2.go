package mp3

import (
	"io"
	"strconv"

	"github.com/ascherkus/go-id3/src/id3"
)

func ID3v2(r io.ReadSeeker) ID3 {
	f := id3.Read(r)

	year, _ := strconv.Atoi(f.Year)
	track, _ := strconv.Atoi(f.Track)

	return ID3{
		Title:  f.Name,
		Artist: f.Artist,
		Album:  f.Album,
		Year:   year,
		Track:  track,
	}
}
