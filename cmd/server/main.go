package main

import (
	"log"

	"github.com/caarlos0/env"

	"github.com/joshbohde/example"
	"github.com/joshbohde/example/http"
	"github.com/joshbohde/example/memory"
)

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`
}

func main() {
	config := Config{}

	err := env.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}

	addresses := memory.AddressService{
		Addresses: map[int]example.Address{1: {ID: 2, Street: "Foo", ZipCode: "12345"}},
	}

	users := memory.UserService{
		Users:          map[int]example.User{1: {ID: 1, Name: "Josh"}},
		AddressService: &addresses,
	}

	server := http.Server{
		AddressService: &addresses,
		UserService:    &users,
	}

	server.Run(config.Port)

}
