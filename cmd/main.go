package main

import (
	"flag"
	"log"
	"restApi/internal/app/api"
	"restApi/internal/app/config"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := config.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := api.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
