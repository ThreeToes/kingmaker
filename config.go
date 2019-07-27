package kingmaker

import (
	"flag"
	"os"
)

type Config struct {
	CharacterFile string `json:"character_file"`
	EventFile     string `json:"event_file"`
	WorldConfig   *struct {
		DateStructure *DateStructure `json:"date_structure"`
		StartDate     *Date          `json:"start_date"`
	} `json:"world_config"`
	EventConfig *struct {
		EventChance int `json:"event_chance"`
	} `json:"event_config"`
}

func GetCmdConfig() *Config {
	configFilePtr := flag.String("config", "", "Path to config file")

	flag.Parse()
	if configFilePtr == nil {
		return &Config{}
	}

	if _, err := os.Stat(*configFilePtr); err != nil {
		return &Config{}
	}

	//bs, err :=

	return &Config{
		CharacterFile: *configFilePtr,
		EventFile:     *eventPtr,
	}
}
