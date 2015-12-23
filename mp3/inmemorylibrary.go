package mp3

import (
	"strings"
	"sync"
)

func NewInMemoryLibrary() *InMemoryLibrary {
	return &InMemoryLibrary{
		make([]MP3, 0),
		&sync.Mutex{},
	}
}

type InMemoryLibrary struct {
	collection []MP3
	*sync.Mutex
}

func (l *InMemoryLibrary) Add(mp3 MP3) error {
	l.Lock()
	defer l.Unlock()

	l.collection = append(
		l.collection,
		mp3,
	)

	return nil
}

func (l InMemoryLibrary) Search(query string) []MP3 {
	res := make([]MP3, 0)

	for _, mp3 := range l.collection {
		if strings.Contains(mp3.ID3.Title, query) {
			res = append(res, mp3)
		}

		if strings.Contains(mp3.ID3.Artist, query) {
			res = append(res, mp3)
		}

		if strings.Contains(mp3.ID3.Album, query) {
			res = append(res, mp3)
		}
	}

	return res
}

func (l InMemoryLibrary) All() []MP3 {
	return l.collection
}

func (l InMemoryLibrary) Count() int {
	return len(l.collection)
}
