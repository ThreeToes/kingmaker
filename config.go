package kingmaker

import "flag"

type Config struct {
	CharacterFile string
	EventFile     string
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
