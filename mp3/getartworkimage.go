package mp3

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetArtworkImage(l Library, g ArtworkGallery) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := mux.Vars(r)["id"]

		if ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		mp3 := l.GetById(ID)

		if mp3 == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		artwork := g.ArtworkFor(mp3)

		if artwork == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Write(artwork.Binary)
		w.Header().Set("Content-Type", artwork.MIME)
	}
}
