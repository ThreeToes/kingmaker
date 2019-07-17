package kingmaker

import (
	"testing"
)

func TestAttributePrecondition_PreconditionMet(t *testing.T) {
	type fields struct {
		Attribute AttributeType
		Operator  Operator
		Value     int
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
			name: "TestEquals",
			want: true,
			args: args{
				c: &Character{
					GivenName:  "Antoin",
					FamilyName: "DeBlois",
					Rank:       DUKE,
					Attributes: &Attributes{
						Diplomacy:   10,
						Intrigue:    10,
						Learning:    10,
						Martial:     10,
						Stewardship: 10,
					},
				},
			},
			fields: fields{
				Operator:  EQUAL,
				Attribute: MARTIAL,
				Value:     10,
			},
		},
		{
			name: "TestEqualsWhenNotEqual",
			want: false,
			args: args{
				c: &Character{
					GivenName:  "Antoin",
					FamilyName: "DeBlois",
					Rank:       DUKE,
					Attributes: &Attributes{
						Diplomacy:   10,
						Intrigue:    10,
						Learning:    10,
						Martial:     1,
						Stewardship: 1,
					},
				},
			},
			fields: fields{
				Operator:  EQUAL,
				Attribute: STEWARDSHIP,
				Value:     10,
			},
		},
		{
			name: "TestGreaterThan",
			want: true,
			args: args{
				c: &Character{
					GivenName:  "Antoin",
					FamilyName: "DeBlois",
					Rank:       DUKE,
					Attributes: &Attributes{
						Diplomacy:   10,
						Intrigue:    10,
						Learning:    10,
						Martial:     10,
						Stewardship: 10,
					},
				},
			},
			fields: fields{
				Operator:  GREATER_THAN,
				Attribute: INTRIGUE,
				Value:     5,
			},
		},
		{
			name: "TestGreaterThanWhenNotGreaterThan",
			want: false,
			args: args{
				c: &Character{
					GivenName:  "Antoin",
					FamilyName: "DeBlois",
					Rank:       DUKE,
					Attributes: &Attributes{
						Diplomacy:   4,
						Intrigue:    10,
						Learning:    10,
						Martial:     1,
						Stewardship: 10,
					},
				},
			},
			fields: fields{
				Operator:  GREATER_THAN,
				Attribute: DIPLOMACY,
				Value:     5,
			},
		},
		{
			name: "TestLessThan",
			want: true,
			args: args{
				c: &Character{
					GivenName:  "Antoin",
					FamilyName: "DeBlois",
					Rank:       DUKE,
					Attributes: &Attributes{
						Diplomacy:   10,
						Intrigue:    10,
						Learning:    10,
						Martial:     10,
						Stewardship: 10,
					},
				},
			},
			fields: fields{
				Operator:  LESS_THAN,
				Attribute: MARTIAL,
				Value:     20,
			},
		},
		{
			name: "TestLessThanWhenNotLessThan",
			want: false,
			args: args{
				c: &Character{
					GivenName:  "Antoin",
					FamilyName: "DeBlois",
					Rank:       DUKE,
					Attributes: &Attributes{
						Diplomacy:   10,
						Intrigue:    10,
						Learning:    10,
						Martial:     25,
						Stewardship: 10,
					},
				},
			},
			fields: fields{
				Operator:  LESS_THAN,
				Attribute: MARTIAL,
				Value:     20,
			},
		},
		{
			name: "TestNotEqual",
			want: true,
			args: args{
				c: &Character{
					GivenName:  "Antoin",
					FamilyName: "DeBlois",
					Rank:       DUKE,
					Attributes: &Attributes{
						Diplomacy:   10,
						Intrigue:    10,
						Learning:    10,
						Martial:     10,
						Stewardship: 10,
					},
				},
			},
			fields: fields{
				Operator:  NOT_EQUAL,
				Attribute: LEARNING,
				Value:     20,
			},
		},
		{
			name: "TestNotEqual",
			want: false,
			args: args{
				c: &Character{
					GivenName:  "Antoin",
					FamilyName: "DeBlois",
					Rank:       DUKE,
					Attributes: &Attributes{
						Diplomacy:   10,
						Intrigue:    10,
						Learning:    10,
						Martial:     10,
						Stewardship: 10,
					},
				},
			},
			fields: fields{
				Operator:  NOT_EQUAL,
				Attribute: LEARNING,
				Value:     10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AttributePrecondition{
				Attribute: tt.fields.Attribute,
				Operator:  tt.fields.Operator,
				Value:     tt.fields.Value,
			}
			if got := p.PreconditionMet(tt.args.c); got != tt.want {
				t.Errorf("AttributePrecondition.PreconditionMet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgePrecondition_PreconditionMet(t *testing.T) {
	type fields struct {
		Age      int
		Operator Operator
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
			fields: fields{
				Age:      10,
				Operator: GREATER_THAN,
			},
			want: true,
			name: "Greater Than",
			args: args{
				c: &Character{
					Age: 16,
				},
			},
		},
		{
			fields: fields{
				Age:      10,
				Operator: GREATER_THAN,
			},
			want: false,
			name: "Greater Than false",
			args: args{
				c: &Character{
					Age: 1,
				},
			},
		},
		{
			fields: fields{
				Age:      10,
				Operator: LESS_THAN,
			},
			want: true,
			name: "Less Than",
			args: args{
				c: &Character{
					Age: 9,
				},
			},
		},
		{
			fields: fields{
				Age:      10,
				Operator: LESS_THAN,
			},
			want: false,
			name: "Less Than False",
			args: args{
				c: &Character{
					Age: 16,
				},
			},
		},
		{
			fields: fields{
				Age:      10,
				Operator: EQUAL,
			},
			want: true,
			name: "Equal",
			args: args{
				c: &Character{
					Age: 10,
				},
			},
		},
		{
			fields: fields{
				Age:      10,
				Operator: EQUAL,
			},
			want: false,
			name: "Equal False",
			args: args{
				c: &Character{
					Age: 16,
				},
			},
		},
		{
			fields: fields{
				Age:      10,
				Operator: NOT_EQUAL,
			},
			want: true,
			name: "Not equal",
			args: args{
				c: &Character{
					Age: 16,
				},
			},
		},
		{
			fields: fields{
				Age:      10,
				Operator: NOT_EQUAL,
			},
			want: false,
			name: "Not Equal False",
			args: args{
				c: &Character{
					Age: 10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgePrecondition{
				Age:      tt.fields.Age,
				Operator: tt.fields.Operator,
			}
			if got := a.PreconditionMet(tt.args.c); got != tt.want {
				t.Errorf("AgePrecondition.PreconditionMet() = %v, want %v", got, tt.want)
			}
		})
	}
}
