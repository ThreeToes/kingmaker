package kingmaker

import "testing"

func Test_fillTemplate(t *testing.T) {
	type args struct {
		w *EventContext
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				w: &EventContext{
					Event: &Event{
						Template: "hello, world",
					},
					ActiveCharacter: &Character{},
				},
			},
			want:    "hello, world",
			name:    "plain string",
			wantErr: false,
		},
		{
			args: args{
				w: &EventContext{
					Event: &Event{
						Template: "hello, {{.ActiveCharacter.GivenName}}",
					},
					ActiveCharacter: &Character{
						GivenName: "Dave",
					},
				},
			},
			want:    "hello, Dave",
			name:    "Single Value Fill in",
			wantErr: false,
		},
		{
			args: args{
				w: &EventContext{
					Event: &Event{
						Template: "hello, {{.ActiveCharacter.GivenName",
					},
					ActiveCharacter: &Character{
						GivenName: "Dave",
					},
				},
			},
			want:    "",
			name:    "Bad template",
			wantErr: true,
		},
		{
			args: args{
				w: &EventContext{
					Event: &Event{
						Template: "hello, {{.ActiveCharacter.NonExistentAttribute}}",
					},
					ActiveCharacter: &Character{
						GivenName: "Dave",
					},
				},
			},
			want:    "",
			name:    "Bad template variable",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FillTemplate(tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("FillTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FillTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
