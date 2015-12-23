package mp3

import (
	"strings"
	"sync"
)

func NewInMemoryLibrary() *InMemoryLibrary {
	return &InMemoryLibrary{
		make(map[string]*Mp3),
		&sync.Mutex{},
	}
}

type InMemoryLibrary struct {
	collection map[string]*Mp3
	*sync.Mutex
}

func (l *InMemoryLibrary) Add(mp3 *Mp3) error {
	l.Lock()
	defer l.Unlock()
	l.collection[mp3.ID] = mp3

	return nil
}

func (l *InMemoryLibrary) GetById(id string) *Mp3 {
	return l.collection[id]
}

func (l *InMemoryLibrary) Search(query string) []*Mp3 {
	res := make([]*Mp3, 0)

	for _, mp3 := range l.collection {
		if mp3.ID == query {
			res = append(res, mp3)
		}

		if mp3.Id3 == nil {
			continue
		}

		if strings.Contains(mp3.Id3.Title, query) {
			res = append(res, mp3)
		}

		if strings.Contains(mp3.Id3.Artist, query) {
			res = append(res, mp3)
		}

		if strings.Contains(mp3.Id3.Album, query) {
			res = append(res, mp3)
		}
	}

	return res
}

func (l *InMemoryLibrary) All() []*Mp3 {
	res := make([]*Mp3, 0, len(l.collection))

	for _, mp3 := range l.collection {
		res = append(res, mp3)
	}

	return res
}

func (l *InMemoryLibrary) Count() int {
	return len(l.collection)
}
