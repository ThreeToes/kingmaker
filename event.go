package kingmaker

import (
	"encoding/json"
	"io/ioutil"
)

type Event struct {
	Id            string         `json:"id"`
	Preconditions []Precondition `json:"preconditions"`
	Template      string         `json:"template"`
}

func (e *Event) PreconditionsMet(c *Character) bool {
	for _, p := range e.Preconditions {
		if !p.PreconditionMet(c) {
			return false
		}
	}
	return true
}

func (e *Event) UnmarshalJSON(b []byte) error {
	var objMap map[string]json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(objMap["id"], &e.Id)
	if err != nil {
		return err
	}
	err = json.Unmarshal(objMap["template"], &e.Template)
	if err != nil {
		return err
	}
	var msgs []*json.RawMessage
	err = json.Unmarshal(objMap["preconditions"], &msgs)
	if err != nil {
		return err
	}
	e.Preconditions, err = unmarshalPrecondtions(msgs)
	if err != nil {
		return err
	}
	return nil
}

func LoadEvents(eventFile string) ([]*Event, error) {
	contents, err := ioutil.ReadFile(eventFile)
	if err != nil {
		return nil, err
	}
	events := []*Event{}
	err = json.Unmarshal([]byte(contents), &events)
	return events, err
}
