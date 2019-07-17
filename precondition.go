package kingmaker

type Operator string
type AttributeType string

const (
	EQUAL        Operator      = "="
	GREATER_THAN Operator      = ">"
	LESS_THAN    Operator      = "<"
	NOT_EQUAL    Operator      = "!="
	DIPLOMACY    AttributeType = "diplomacy"
	MARTIAL      AttributeType = "martial"
	LEARNING     AttributeType = "learning"
	INTRIGUE     AttributeType = "intrigue"
	STEWARDSHIP  AttributeType = "stewardship"
)

type Precondition interface {
	PreconditionMet(c *Character) bool
}

type AttributePrecondition struct {
	Attribute AttributeType `json:"attribute"`
	Operator  Operator      `json:"operator"`
	Value     int           `json:"value"`
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

func getAttribute(c *Character, attr AttributeType) int {
	switch attr {
	case MARTIAL:
		return c.Attributes.Martial
	case LEARNING:
		return c.Attributes.Learning
	case DIPLOMACY:
		return c.Attributes.Diplomacy
	case INTRIGUE:
		return c.Attributes.Intrigue
	case STEWARDSHIP:
		return c.Attributes.Stewardship
	default:
		return -1
	}
}
