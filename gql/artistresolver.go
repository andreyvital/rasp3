package gql

import (
	"strings"

	"github.com/CentaurWarchief/rasp3/mp3"
	"github.com/graphql-go/graphql"
)

type AlbumObject struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

type ArtistObject struct {
	Count  int            `json:"count"`
	Name   string         `json:"name"`
	Albums []*AlbumObject `json:"albums"`
}

func ArtistResolver(l mp3.Library) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		artists := make(map[string]*ArtistObject, 0)
		albums := make(map[string]map[string]*AlbumObject, 0)

		for _, mp3 := range l.All() {
			if mp3.ID3.Artist == "" {
				continue
			}

			artist := strings.ToLower(mp3.ID3.Artist)

			if artists[artist] == nil {
				artists[artist] = &ArtistObject{
					Name: mp3.ID3.Artist,
				}
			}

			if mp3.ID3.Album == "" {
				continue
			}

			album := strings.ToLower(mp3.ID3.Album)

			if albums[artist] == nil {
				albums[artist] = make(map[string]*AlbumObject, 0)
			}

			if albums[artist][album] == nil {
				albums[artist][album] = &AlbumObject{
					Count: 0,
					Name:  mp3.ID3.Album,
				}
			}

			albums[artist][album].Count++
			artists[artist].Count = len(albums[artist])
		}

		for artist, _ := range albums {
			for _, album := range albums[artist] {
				artists[artist].Albums = append(
					artists[artist].Albums,
					album,
				)
			}
		}

		res := []*ArtistObject{}

		for _, artist := range artists {
			res = append(res, artist)
		}

		return res, nil
	}
}
