package gql

import (
	"fmt"

	"github.com/CentaurWarchief/rasp3/mp3"
	"github.com/graphql-go/graphql"
)

func Mp3(ar ArtworkReferenceResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "MP3",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"file": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"id3": &graphql.Field{
				Type: ID3,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p.Source.(*mp3.Mp3).Id3 != nil {
						return p.Source.(*mp3.Mp3).Id3, nil
					}

					return nil, nil
				},
			},
			"artwork": &graphql.Field{
				Type: Artwork,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ref := ar(p.Source.(*mp3.Mp3))

					if ref != nil {
						ref.URL = fmt.Sprintf(
							"%s/mp3/%s/artwork",
							p.Info.RootValue.(map[string]interface{})["baseURL"].(string),
							p.Source.(*mp3.Mp3).ID,
						)

						return ref, nil
					}

					return nil, nil
				},
			},
		},
	})
}
