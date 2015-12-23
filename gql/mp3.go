package gql

import "github.com/graphql-go/graphql"

var MP3 = graphql.NewObject(graphql.ObjectConfig{
	Name: "MP3",
	Fields: graphql.Fields{
		"file": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"size": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"id3": &graphql.Field{
			Type: ID3,
		},
	},
})
