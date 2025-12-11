package main

import (
	"log"
	"os"

	"github.com/baletor/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registered: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}

}
