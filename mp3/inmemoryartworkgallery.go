package mp3

import "sync"

func NewInMemoryArtworkGallery(loader ArtworkLoader) *InMemoryArtworkGallery {
	return &InMemoryArtworkGallery{
		loader,
		make(map[string]*Artwork),
		make([]string, 0),
		&sync.Mutex{},
	}
}

type InMemoryArtworkGallery struct {
	loader   ArtworkLoader
	gallery  map[string]*Artwork
	attempts []string
	*sync.Mutex
}

func (g *InMemoryArtworkGallery) ArtworkFor(mp3 *Mp3) *Artwork {
	g.Lock()
	defer g.Unlock()

	if g.gallery[mp3.ID] == nil {
		for _, id := range g.attempts {
			if id == mp3.ID {
				return nil
			}
		}
	}

	if g.gallery[mp3.ID] != nil {
		return g.gallery[mp3.ID]
	}

	g.attempts = append(g.attempts, mp3.ID)

	if artwork := g.loader(mp3); artwork != nil {
		g.gallery[mp3.ID] = artwork
		return artwork
	}

	return nil
}
