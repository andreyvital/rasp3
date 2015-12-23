package mp3

type Library interface {
	Add(mp3 MP3) error
	All() []MP3
	Search(query string) []MP3
	Count() int
}
