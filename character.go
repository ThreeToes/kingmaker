package kingmaker

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
	GivenName  string      `json:"given_name"`
	FamilyName string      `json:"family_name"`
	Attributes *Attributes `json:"attributes"`
	Rank       Rank        `json:"rank"`
}

type Attributes struct {
	Diplomacy   int
	Learning    int
	Martial     int
	Intrigue    int
	Stewardship int
}
