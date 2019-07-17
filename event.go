package kingmaker

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
