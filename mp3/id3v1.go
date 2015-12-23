package mp3

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func byteString(b []byte) string {
	p := bytes.IndexByte(b, 0)

	if p == -1 {
		p = len(b)
	}

	return string(b[0:p])
}

// http://en.wikipedia.org/wiki/ID3#Layout
func ID3v1(r io.ReadSeeker) ID3 {
	b := make([]byte, 128)

	r.Seek(-128, os.SEEK_END)
	r.Read(b)

	if string(b[0:3]) != "TAG" {
		return ID3{}
	}

	b = b[3:]

	title := byteString(b[0:30])
	artist := byteString(b[30:60])
	album := byteString(b[60:90])
	year, _ := strconv.Atoi(byteString(b[90:94]))
	var track int

	if b[122] == 0 {
		track, _ = strconv.Atoi(fmt.Sprintf("%d", b[123]))
	}

	return ID3{
		Title:  title,
		Artist: artist,
		Album:  album,
		Year:   year,
		Track:  track,
	}
}
