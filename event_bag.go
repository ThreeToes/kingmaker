package kingmaker

import "math/rand"

type EventBag struct {
	events []*Event
	random *rand.Rand
}

func (b *EventBag) AddEvent(e *Event) {
	b.events = append(b.events, e)
}

func (b *EventBag) DrawEvent() *Event {
	if len(b.events) == 0 {
		return nil
	}
	elen := len(b.events)
	choice := int(b.random.Uint32()) % elen
	return b.events[choice]
}
