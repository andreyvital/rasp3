package gql

import "github.com/graphql-go/graphql"

var Artist = graphql.NewObject(graphql.ObjectConfig{
	Name: "Artist",
	Fields: graphql.Fields{
		"count": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"albums": &graphql.Field{
			Type: graphql.NewList(Album),
		},
	},
})
