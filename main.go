package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/CentaurWarchief/rasp3/config"
	"github.com/CentaurWarchief/rasp3/mp3"
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
		go mp3.Discover(
			library,
			path,
		)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ListenPort), nil); err != nil {
		fmt.Println(err.Error())
		return
	}
}
