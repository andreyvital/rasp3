package mp3

type ArtworkGallery interface {
	ArtworkFor(mp3 *Mp3) *Artwork
}
