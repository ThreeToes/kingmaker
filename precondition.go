package kingmaker

import "strings"

type Operator string

const (
	EQUAL        Operator = "="
	GREATER_THAN Operator = ">"
	LESS_THAN    Operator = "<"
	NOT_EQUAL    Operator = "!="
)

type Precondition interface {
	PreconditionMet(c *Character) bool
}

type AttributePrecondition struct {
	Attribute string   `json:"attribute"`
	Operator  Operator `json:"operator"`
	Value     int      `json:"value"`
}

func (p *AttributePrecondition) PreconditionMet(c *Character) bool {
	val := getAttribute(c, p.Attribute)
	switch p.Operator {
	case EQUAL:
		return val == p.Value
	case GREATER_THAN:
		return val > p.Value
	case LESS_THAN:
		return val < p.Value
	case NOT_EQUAL:
		return val != p.Value
	}
	return false
}

func getAttribute(c *Character, attr string) int {
	switch strings.ToLower(attr) {
	case "martial":
		return c.Attributes.Martial
	case "learning":
		return c.Attributes.Learning
	case "diplomacy":
		return c.Attributes.Diplomacy
	case "intrigue":
		return c.Attributes.Intrigue
	case "stewardship":
		return c.Attributes.Stewardship
	default:
		return -1
	}
}
