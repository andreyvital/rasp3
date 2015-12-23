package gql

import "github.com/CentaurWarchief/rasp3/mp3"

func GalleryArtworkResolver(g mp3.ArtworkGallery) ArtworkReferenceResolver {
	return func(mp3 *mp3.Mp3) *ArtworkReference {
		if artwork := g.ArtworkFor(mp3); artwork != nil {
			return &ArtworkReference{
				MIME:   artwork.MIME,
				Width:  artwork.Width,
				Height: artwork.Height,
			}
		}

		return nil
	}
}
