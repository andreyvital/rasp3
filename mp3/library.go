package mp3

type Library interface {
	Add(mp3 *Mp3) error
	GetById(id string) *Mp3
	Search(query string) []*Mp3
}
