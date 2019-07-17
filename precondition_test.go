package kingmaker

import "testing"

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
