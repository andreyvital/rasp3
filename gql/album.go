package gql

import "github.com/graphql-go/graphql"

var Album = graphql.NewObject(graphql.ObjectConfig{
	Name: "Album",
	Fields: graphql.Fields{
		"count": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
