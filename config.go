package kingmaker

import "flag"

type Config struct {
	CharacterFile string `json:"character_file"`
	EventFile     string `json:"event_file"`
	WorldConfig   *struct {
	}
}

type DateStructure struct {
}

func GetCmdConfig() *Config {
	characterPtr := flag.String("c", "", "Path to characters file")
	eventPtr := flag.String("e", "", "Path to events file")

	flag.Parse()
	return &Config{
		CharacterFile: *characterPtr,
		EventFile:     *eventPtr,
	}
}
