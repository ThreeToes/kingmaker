package kingmaker

type Event struct {
	Preconditions []Precondition
	Template      string
}

func (e *Event) PreconditionsMet(c *Character) bool {
	for _, p := range e.Preconditions {
		if !p.PreconditionMet(c) {
			return false
		}
	}
	return true
}
