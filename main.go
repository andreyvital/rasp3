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

	r := mux.NewRouter().StrictSlash(true)

	r.Methods("POST").Path("/query").HandlerFunc(gql.QueryHandler(library))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ListenPort), r); err != nil {
		fmt.Println(err.Error())
		return
	}
}
