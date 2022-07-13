package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	log.Println(config)
	if err != nil {
		log.Println("Error decode")
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
