package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/apiserver"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"
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
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	user := &model.User{}
	fmt.Print("Login:")
	fmt.Scan(&user.Login)
	fmt.Print("Password:")
	fmt.Scan(&user.Password)
	fmt.Print("Name:")
	fmt.Scan(&user.Name)
	fmt.Print("Surname:")
	fmt.Scan(&user.Surname)
	//fmt.Scanf("name: %s", &user.Name)
	//fmt.Scanf("surname: %s", &user.Surname)
	if err := user.Validate(); err != nil {
		log.Fatal("Not valid data")
	}

	if err := user.BeforeCreate(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	user.Sanitize()
	if err := db.QueryRow("INSERT INTO users (login, encrypted_password, name, surname) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Login,
		user.EncryptedPassword,
		user.Name,
		user.Surname,
	).Scan(&user.ID); err != nil {
		log.Fatal(err)
	}

	if err := db.QueryRow("INSERT INTO user_roles (user_id, role) VALUES ($1, $2)",
		user.ID,
		"admin",
	); err != nil {
		log.Fatal(err)
	}

}
