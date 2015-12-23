package gql

import "github.com/graphql-go/graphql"

var ID3 = graphql.NewObject(graphql.ObjectConfig{
	Name: "ID3",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"artist": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"album": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"track": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
