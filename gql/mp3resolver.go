package gql

import (
	"github.com/CentaurWarchief/rasp3/mp3"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/graphql-go/graphql"
)

func MP3Resolver(l mp3.Library) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		limit := 10

		if p.Args["limit"] != nil {
			limit = p.Args["limit"].(int)
		}

		if limit == 0 || limit > 30 {
			limit = 10
		}

		var q Query

		if p.Args["query"] != nil && p.Args["query"].(string) != "" {
			q = From(l.Search(p.Args["query"].(string)))
		} else {
			q = From(l.All())
		}

		return q.Take(limit).Results()
	}
}
