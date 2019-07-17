package kingmaker

import (
	"testing"
)

func TestEvent_PreconditionsMet(t *testing.T) {
	type fields struct {
		Preconditions []Precondition
		Template      string
	}
	type args struct {
		c *Character
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Empty precondition lists",
			want: true,
			args: args{
				c: &Character{
					GivenName:  "Davo",
					FamilyName: "Davidson",
					Rank:       DUKE,
					Attributes: &Attributes{
						Stewardship: 10,
						Intrigue:    10,
						Diplomacy:   10,
						Learning:    10,
						Martial:     10,
					},
				},
			},
			fields: fields{
				Preconditions: nil,
				Template:      "",
			},
		},
		{
			name: "One False",
			want: false,
			args: args{
				c: &Character{
					GivenName:  "Davo",
					FamilyName: "Davidson",
					Rank:       DUKE,
					Attributes: &Attributes{
						Stewardship: 10,
						Intrigue:    10,
						Diplomacy:   10,
						Learning:    10,
						Martial:     10,
					},
				},
			},
			fields: fields{
				Preconditions: []Precondition{
					&AttributePrecondition{
						Value:     11,
						Operator:  EQUAL,
						Attribute: STEWARDSHIP,
					},
				},
				Template: "",
			},
		},
		{
			name: "One True",
			want: true,
			args: args{
				c: &Character{
					GivenName:  "Davo",
					FamilyName: "Davidson",
					Rank:       DUKE,
					Attributes: &Attributes{
						Stewardship: 10,
						Intrigue:    10,
						Diplomacy:   10,
						Learning:    10,
						Martial:     10,
					},
				},
			},
			fields: fields{
				Preconditions: []Precondition{
					&AttributePrecondition{
						Value:     15,
						Operator:  LESS_THAN,
						Attribute: STEWARDSHIP,
					},
				},
				Template: "",
			},
		},
		{
			name: "Multi False",
			want: false,
			args: args{
				c: &Character{
					GivenName:  "Davo",
					FamilyName: "Davidson",
					Rank:       DUKE,
					Attributes: &Attributes{
						Stewardship: 10,
						Intrigue:    10,
						Diplomacy:   10,
						Learning:    10,
						Martial:     10,
					},
				},
			},
			fields: fields{
				Preconditions: []Precondition{
					&AttributePrecondition{
						Value:     11,
						Operator:  EQUAL,
						Attribute: DIPLOMACY,
					},
					&AttributePrecondition{
						Value:     8,
						Operator:  LESS_THAN,
						Attribute: MARTIAL,
					},
				},
				Template: "",
			},
		},
		{
			name: "Multi True",
			want: true,
			args: args{
				c: &Character{
					GivenName:  "Davo",
					FamilyName: "Davidson",
					Rank:       DUKE,
					Attributes: &Attributes{
						Stewardship: 10,
						Intrigue:    10,
						Diplomacy:   10,
						Learning:    10,
						Martial:     10,
					},
				},
			},
			fields: fields{
				Preconditions: []Precondition{
					&AttributePrecondition{
						Value:     10,
						Operator:  EQUAL,
						Attribute: LEARNING,
					},
					&AttributePrecondition{
						Value:     13,
						Operator:  LESS_THAN,
						Attribute: INTRIGUE,
					},
				},
				Template: "",
			},
		},
		{
			name: "Mixed True False",
			want: false,
			args: args{
				c: &Character{
					GivenName:  "Davo",
					FamilyName: "Davidson",
					Rank:       DUKE,
					Attributes: &Attributes{
						Stewardship: 10,
						Intrigue:    10,
						Diplomacy:   10,
						Learning:    10,
						Martial:     10,
					},
				},
			},
			fields: fields{
				Preconditions: []Precondition{
					&AttributePrecondition{
						Value:     10,
						Operator:  EQUAL,
						Attribute: DIPLOMACY,
					},
					&AttributePrecondition{
						Value:     8,
						Operator:  LESS_THAN,
						Attribute: STEWARDSHIP,
					},
				},
				Template: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				Preconditions: tt.fields.Preconditions,
				Template:      tt.fields.Template,
			}
			if got := e.PreconditionsMet(tt.args.c); got != tt.want {
				t.Errorf("Event.PreconditionsMet() = %v, want %v", got, tt.want)
			}
		})
	}
}
