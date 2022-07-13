package main

import (
	"flag"
	"fmt"
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
	fmt.Println(config)
	fmt.Println(configPath)
	_, err := toml.Decode(configPath, config)
	fmt.Println(config)
	if err != nil {
		fmt.Println("Error decode")
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
