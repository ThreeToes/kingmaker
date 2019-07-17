package kingmaker

type Event struct {
	Preconditions []*AttributePrecondition
	Template      string
}
