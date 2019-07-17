package kingmaker

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventBag_AddEvent(t *testing.T) {
	type fields struct {
		events []*Event
		random *rand.Rand
	}
	type args struct {
		e *Event
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add new event from empty",
			args: args{
				e: &Event{
					Template:      "some template",
					Preconditions: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		// Reliable output
		random := rand.New(rand.NewSource(1))
		t.Run(tt.name, func(t *testing.T) {
			b := &EventBag{
				events: tt.fields.events,
				random: random,
			}
			b.AddEvent(tt.args.e)
			assert.Contains(t, b.events, tt.args.e)
		})
	}
}

func TestEventBag_DrawEvent(t *testing.T) {
	type fields struct {
		events []*Event
	}
	tests := []struct {
		name   string
		fields fields
		want   *Event
	}{
		{
			name: "empty",
			fields: fields{
				events: nil,
			},
			want: nil,
		},
		{
			name: "single",
			fields: fields{
				events: []*Event{
					{
						Template:      "hello",
						Preconditions: nil,
					},
				},
			},
			want: &Event{
				Template:      "hello",
				Preconditions: nil,
			},
		},
		//2596996162: Seed(1)
		{
			name: "two",
			fields: fields{
				events: []*Event{
					{
						Template:      "hello",
						Preconditions: nil,
					},
					{
						Template:      "goodbye",
						Preconditions: nil,
					},
				},
			},
			want: &Event{
				Template:      "hello",
				Preconditions: nil,
			},
		},
		{
			name: "three",
			fields: fields{
				events: []*Event{
					{
						Template:      "hello",
						Preconditions: nil,
					},
					{
						Template:      "good day",
						Preconditions: nil,
					},
					{
						Template:      "goodbye",
						Preconditions: nil,
					},
				},
			},
			want: &Event{
				Template:      "good day",
				Preconditions: nil,
			},
		},
	}
	for _, tt := range tests {
		// Reliable output
		random := rand.New(rand.NewSource(1))
		t.Run(tt.name, func(t *testing.T) {
			b := &EventBag{
				events: tt.fields.events,
				random: random,
			}
			if got := b.DrawEvent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventBag.DrawEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
