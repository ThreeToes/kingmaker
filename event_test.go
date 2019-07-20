package kingmaker

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestEvent_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	type expectedFields struct {
		Id            string
		Template      string
		Preconditions []Precondition
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		expected expectedFields
	}{
		{
			name: "no preconditions",
			args: args{
				b: []byte(`
					{
						"id":"sickskid",
						"template": "Davo did a sick skid",
						"preconditions":[]
					}
				`),
			},
			expected: expectedFields{
				Preconditions: nil,
				Template:      "Davo did a sick skid",
				Id:            "sickskid",
			},
			wantErr: false,
		},
		{
			name: "single precondition",
			args: args{
				b: []byte(`
					{
						"id":"sickskid",
						"template": "Davo did a sick skid",
						"preconditions":[
							{
								"operator": "<",
								"age": 25
							}
						]
					}
				`),
			},
			expected: expectedFields{
				Preconditions: []Precondition{
					&AgePrecondition{
						Operator: LESS_THAN,
						Age:      25,
					},
				},
				Template: "Davo did a sick skid",
				Id:       "sickskid",
			},
			wantErr: false,
		},
		{
			name: "mixed precondition",
			args: args{
				b: []byte(`
					{
						"id":"sickskid",
						"template": "Davo did a sick skid",
						"preconditions":[
							{
								"operator": "<",
								"age": 25
							},
							{
								"operator": "=",
								"value": 10,
								"attribute":"intrigue"
							}
						]
					}
				`),
			},
			expected: expectedFields{
				Preconditions: []Precondition{
					&AgePrecondition{
						Operator: LESS_THAN,
						Age:      25,
					},
					&AttributePrecondition{
						Operator:  EQUAL,
						Value:     10,
						Attribute: INTRIGUE,
					},
				},
				Template: "Davo did a sick skid",
				Id:       "sickskid",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{}
			if err := e.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Event.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.expected.Id, e.Id)
			assert.Equal(t, tt.expected.Template, e.Template)
			assert.Len(t, e.Preconditions, len(tt.expected.Preconditions))
			for i, p := range e.Preconditions {
				if !assert.True(t, reflect.DeepEqual(p, tt.expected.Preconditions[i])) {
					t.FailNow()
				}
			}
		})
	}
}
