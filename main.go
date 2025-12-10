package main

import (
	"fmt"
	"log"

	"github.com/baletor/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("Baletor")
	if err != nil {
		log.Fatal(err)
	}

	cfg2, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg2)
}
