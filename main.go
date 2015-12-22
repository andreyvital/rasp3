package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/CentaurWarchief/rasp3/config"
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

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ListenPort), nil); err != nil {
		fmt.Println(err.Error())
		return
	}
}
