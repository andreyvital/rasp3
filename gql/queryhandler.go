package gql

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CentaurWarchief/rasp3/mp3"

	"github.com/graphql-go/graphql"
)

func QueryHandler(l mp3.Library) http.HandlerFunc {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Root",
			Fields: graphql.Fields{
				"mp3": &graphql.Field{
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"query": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"album": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"artist": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Type:    graphql.NewList(MP3),
					Resolve: MP3Resolver(l),
				},
				"artists": &graphql.Field{
					Type:    graphql.NewList(Artist),
					Resolve: ArtistResolver(l),
				},
			},
		}),
	})

	if err != nil {
		log.Println(err)
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/graphql" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		query, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer r.Body.Close()

		res := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})

		if res.HasErrors() {
			for _, err := range res.Errors {
				log.Println(err.Error())
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(res.Data); err != nil {
			log.Println(err)
		}
	}
}
