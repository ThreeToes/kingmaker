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

	events, err := kingmaker.LoadEvents(cfg.EventFile)
	if err != nil {
		log.Println("Error while loading events: ", err.Error())
		os.Exit(1)
	}

	for _, e := range events {
		for _, c := range characters {
			if e.PreconditionsMet(c) {
				s, _ := kingmaker.FillTemplate(&kingmaker.EventContext{
					Event:           e,
					ActiveCharacter: c,
				})
				log.Println(s)
			}
		}
	}
}
