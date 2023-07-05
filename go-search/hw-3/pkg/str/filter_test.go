package str

import "testing"

func TestFilterString(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{
				text: "",
			},
			want: "",
		},
		{
			name: "Normal text with space and symbols",
			args: args{
				text: " text , (text), [text]",
			},
			want: "text  text text",
		},
		{
			name: "Normal text with space and symbols",
			args: args{
				text: " text , (text), - [text] ",
			},
			want: "text  text  text",
		},
		{
			name: "Normal text with space and symbols",
			args: args{
				text: " text - go , (text), [text]",
			},
			want: "text  go  text text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterString(tt.args.text); got != tt.want {
				t.Errorf("FilterString() = %v, want %v", got, tt.want)
			}
		})
	}
}
