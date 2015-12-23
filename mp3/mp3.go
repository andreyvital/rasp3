package mp3

import "github.com/twinj/uuid"

type Mp3 struct {
	ID   string `json:"id"`
	File string `json:"file"`
	Size int    `json:"size"`
	Id3  *ID3
}

func New(file string, size int, id3 *ID3) *Mp3 {
	return &Mp3{
		ID:   uuid.NewV4().String(),
		File: file,
		Size: size,
		Id3:  id3,
	}
}

func NewWithoutId3(file string, size int) *Mp3 {
	return &Mp3{
		ID:   uuid.NewV4().String(),
		File: file,
		Size: size,
		Id3:  nil,
	}
}
