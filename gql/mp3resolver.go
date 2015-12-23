package gql

import (
	"strings"

	"github.com/CentaurWarchief/rasp3/mp3"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/graphql-go/graphql"
)

func MP3Resolver(l mp3.Library) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		var q Query

		if p.Args["query"] != nil && p.Args["query"].(string) != "" {
			q = From(l.Search(p.Args["query"].(string)))
		} else {
			q = From(l.All())
		}

		if p.Args["limit"] != nil {
			q = q.Take(p.Args["limit"].(int))
		}

		if p.Args["artist"] != nil && p.Args["artist"].(string) != "" {
			q = q.Where(func(m T) (bool, error) {
				return strings.Contains(
					m.(mp3.MP3).ID3.Artist,
					p.Args["artist"].(string),
				), nil
			})
		}

		if p.Args["album"] != nil && p.Args["album"].(string) != "" {
			q = q.Where(func(m T) (bool, error) {
				return strings.Contains(
					m.(mp3.MP3).ID3.Album,
					p.Args["album"].(string),
				), nil
			})
		}

		return q.Results()
	}
}
