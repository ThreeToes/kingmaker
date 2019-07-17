package main

import (
	"flag"
	"log"
	"os"

	"github.com/threetoes/kingmaker"
)

func main() {
	cfg := kingmaker.GetCmdConfig()
	if _, err := os.Stat(cfg.EventFile); err != nil {
		log.Println("Event file issue: %s", err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}
	if _, err := os.Stat(cfg.CharacterFile); err != nil {
		log.Println("Character file issue: %s", err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}
	characters, err := kingmaker.LoadCharacters(cfg.CharacterFile)
	if err != nil {
		log.Println("Error while loading character file: ", err.Error())
		os.Exit(1)
	}
	for _, c := range characters {
		log.Printf("%+v\n", c)
	}
}
