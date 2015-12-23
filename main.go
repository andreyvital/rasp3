package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/CentaurWarchief/rasp3/config"
	"github.com/CentaurWarchief/rasp3/gql"
	"github.com/CentaurWarchief/rasp3/mp3"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	defer func() {
		os.Exit(1)
	}()

	var cfg config.Config

	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		fmt.Println(err.Error())
		return
	}

	library := mp3.NewInMemoryLibrary()

	for _, path := range cfg.Discover {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		go mp3.Discover(library, path)
	}

	var h http.Handler

	r := mux.NewRouter().StrictSlash(true)

	r.Methods("POST").Path("/query").HandlerFunc(gql.QueryHandler(library))

	h = r
	h = GetCors().Handler(h)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ListenPort), h); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func GetCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Content-Type"},
		MaxAge:           604800,
	})
}
