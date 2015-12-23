package mp3

import (
	"bytes"
	"image/jpeg"
	"image/png"

	"github.com/mikkyang/id3-go"
	"github.com/mikkyang/id3-go/v2"
)

func Id3ArtworkLoader(mp3 Mp3) *Artwork {
	f, err := id3.Open(mp3.File)

	if err != nil {
		return nil
	}

	frame := f.Frame(AttachedPictureFrame)

	if frame == nil {
		return nil
	}

	image := frame.(*v2.ImageFrame)

	bytes := bytes.NewReader(image.Data())

	var width, height int

	mime := image.MIMEType()

	if mime == "image/jpeg" || mime == "image/jpg" {
		cfg, err := jpeg.DecodeConfig(bytes)

		if err == nil {
			width = cfg.Width
			height = cfg.Height
		}
	}

	if mime == "image/png" {
		cfg, err := png.DecodeConfig(bytes)

		if err == nil {
			width = cfg.Width
			height = cfg.Height
		}
	}

	return &Artwork{
		Binary: image.Data(),
		MIME:   mime,
		Width:  width,
		Height: height,
	}
}
