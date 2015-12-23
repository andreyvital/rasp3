package mp3

import "github.com/twinj/uuid"

type Mp3 struct {
	ID   string `json:"id"`
	File string `json:"file"`
	Id3  *ID3
}

func New(file string, id3 *ID3) *Mp3 {
	return &Mp3{
		ID:   uuid.NewV4().String(),
		File: file,
		Id3:  id3,
	}
}

func NewWithoutId3(file string) *Mp3 {
	return &Mp3{
		ID:   uuid.NewV4().String(),
		File: file,
		Id3:  nil,
	}
}
