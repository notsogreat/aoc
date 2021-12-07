package dive

import (
	"testing"
)

var input = []Data{
	{
		position: "forward",
		value:    5,
	},
	{
		position: "down",
		value:    5,
	},
	{
		position: "forward",
		value:    8,
	},
	{
		position: "up",
		value:    3,
	},
	{
		position: "down",
		value:    8,
	},
	{
		position: "forward",
		value:    2,
	},
}

func TestSubmarineDepth(t *testing.T) {

	type args struct {
		input []Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{input},
			want: 150,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubmarineDepth(tt.args.input); got != tt.want {
				t.Errorf("SubmarineDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCorrectSubmarineDepth(t *testing.T) {
	type args struct {
		input []Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{input},
			want: 900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CorrectSubmarineDepth(tt.args.input); got != tt.want {
				t.Errorf("CorrectSubmarineDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
