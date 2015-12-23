package mp3

import "sync"

func NewInMemoryArtworkGallery(loader ArtworkLoader) *InMemoryArtworkGallery {
	return &InMemoryArtworkGallery{
		loader,
		make(map[string]*Artwork),
		make(map[string]bool),
		&sync.Mutex{},
	}
}

type InMemoryArtworkGallery struct {
	loader   ArtworkLoader
	gallery  map[string]*Artwork
	attempts map[string]bool
	*sync.Mutex
}

func (g *InMemoryArtworkGallery) ArtworkFor(mp3 Mp3) *Artwork {
	g.Lock()
	defer g.Unlock()

	// if g.attempts[mp3.ID] == false {
	// 	return nil
	// }

	if g.gallery[mp3.ID] != nil {
		return g.gallery[mp3.ID]
	}

	if artwork := g.loader(mp3); artwork != nil {
		g.gallery[mp3.ID] = artwork
		g.attempts[mp3.ID] = true

		return artwork
	}

	g.attempts[mp3.ID] = false
	return nil
}
