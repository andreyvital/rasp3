package mp3

type MP3 struct {
	ID   string `json:"id"`
	File string `json:"file"`
	Size int    `json:"size"`
	ID3  ID3    `json:"id3"`
}

func New(file string, size int, id3 ID3) MP3 {
	return MP3{
		File: file,
		Size: size,
		ID3:  id3,
	}
}

func NewWithoutID3(file string, size int) MP3 {
	return MP3{
		File: file,
		Size: size,
	}
}
