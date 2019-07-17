package kingmaker

import (
	"encoding/json"
	"io/ioutil"
)

type Rank int

const (
	PEASANT Rank = iota
	BARON   Rank = iota
	COUNT   Rank = iota
	DUKE    Rank = iota
	KING    Rank = iota
	EMPEROR Rank = iota
)

type Character struct {
	Id         string      `json:"id"`
	GivenName  string      `json:"given_name"`
	FamilyName string      `json:"family_name"`
	Attributes *Attributes `json:"attributes"`
	Rank       Rank        `json:"rank"`
	Age        int         `json:"age"`
}

type Attributes struct {
	Diplomacy   int
	Learning    int
	Martial     int
	Intrigue    int
	Stewardship int
}

func LoadCharacters(filepath string) ([]*Character, error) {
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	chars := []*Character{}
	err = json.Unmarshal([]byte(contents), &chars)
	return chars, err
}
