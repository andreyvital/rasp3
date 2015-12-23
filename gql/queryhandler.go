package gql

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CentaurWarchief/rasp3/mp3"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/chnlr/baseurl"
	"github.com/graphql-go/graphql"
)

func QueryHandler(
	l mp3.Library,
	g mp3.ArtworkGallery,
) http.HandlerFunc {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Root",
			Fields: graphql.Fields{
				"mp3": &graphql.Field{
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Type: graphql.NewList(Mp3(GalleryArtworkResolver(g))),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						limit := 10

						if p.Args["limit"] != nil {
							limit = p.Args["limit"].(int)
						}

						if limit == 0 || limit > 30 {
							limit = 10
						}

						return From(l.All()).Take(limit).Results()
					},
				},
			},
		}),
	})

	if err != nil {
		log.Println(err)
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		query, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		r.Body.Close()

		res := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
			RootObject:    map[string]interface{}{"baseURL": baseurl.BaseUrl(r)},
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
