package gql

import "github.com/graphql-go/graphql"

var Artwork = graphql.NewObject(graphql.ObjectConfig{
	Name: "Artwork",
	Fields: graphql.Fields{
		"url": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"mime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"width": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"height": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
