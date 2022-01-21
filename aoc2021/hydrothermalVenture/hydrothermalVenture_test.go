package hydrothermalventure

import (
	"reflect"
	"testing"
)

var line1 = lines{start: points{x: 0, y: 9}, end: points{x: 5, y: 9}}
var line2 = lines{start: points{x: 8, y: 0}, end: points{x: 0, y: 8}}
var line3 = lines{start: points{x: 9, y: 4}, end: points{x: 3, y: 4}}
var line4 = lines{start: points{x: 2, y: 2}, end: points{x: 2, y: 1}}
var line5 = lines{start: points{x: 7, y: 0}, end: points{x: 7, y: 4}}
var line6 = lines{start: points{x: 6, y: 4}, end: points{x: 2, y: 0}}
var line7 = lines{start: points{x: 0, y: 9}, end: points{x: 2, y: 9}}
var line8 = lines{start: points{x: 3, y: 4}, end: points{x: 1, y: 4}}
var line9 = lines{start: points{x: 0, y: 0}, end: points{x: 8, y: 8}}
var line10 = lines{start: points{x: 5, y: 5}, end: points{x: 8, y: 2}}

func TestGetPoints(t *testing.T) {
	type args struct {
		start points
		end   points
	}
	tests := []struct {
		name string
		args args
		want []points
	}{
		{
			name: "test1",
			args: args{start: points{x: 0, y: 9}, end: points{x: 5, y: 9}},
			want: []points{{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9}, {x: 3, y: 9}, {x: 4, y: 9}, {x: 5, y: 9}},
		},
		{
			name: "test2",
			args: args{start: points{x: 8, y: 0}, end: points{x: 0, y: 8}},
			want: []points{{x: 8, y: 0}, {x: 7, y: 1}, {x: 6, y: 2}, {x: 5, y: 3}, {x: 4, y: 4}, {x: 3, y: 5}, {x: 2, y: 6}, {x: 1, y: 7}, {x: 0, y: 8}},
		},
		{
			name: "test3",
			args: args{start: points{x: 9, y: 4}, end: points{x: 3, y: 4}},
			want: []points{{x: 9, y: 4}, {x: 8, y: 4}, {x: 7, y: 4}, {x: 6, y: 4}, {x: 5, y: 4}, {x: 4, y: 4}, {x: 3, y: 4}},
		},
		{
			name: "test4",
			args: args{start: points{x: 2, y: 2}, end: points{x: 2, y: 1}},
			want: []points{{x: 2, y: 2}, {x: 2, y: 1}},
		},
		{
			name: "test5",
			args: args{start: points{x: 7, y: 0}, end: points{x: 7, y: 4}},
			want: []points{{x: 7, y: 0}, {x: 7, y: 1}, {x: 7, y: 2}, {x: 7, y: 3}, {x: 7, y: 4}},
		},
		{
			name: "test6",
			args: args{start: points{x: 6, y: 4}, end: points{x: 2, y: 0}},
			want: []points{{x: 6, y: 4}, {x: 5, y: 3}, {x: 4, y: 2}, {x: 3, y: 1}, {x: 2, y: 0}},
		},
		{
			name: "test7",
			args: args{start: points{x: 0, y: 9}, end: points{x: 2, y: 9}},
			want: []points{{x: 0, y: 9}, {x: 1, y: 9}, {x: 2, y: 9}},
		},
		{
			name: "test8",
			args: args{start: points{x: 3, y: 4}, end: points{x: 1, y: 4}},
			want: []points{{x: 3, y: 4}, {x: 2, y: 4}, {x: 1, y: 4}},
		},
		{
			name: "test9",
			args: args{start: points{x: 0, y: 0}, end: points{x: 8, y: 8}},
			want: []points{{x: 0, y: 0}, {x: 1, y: 1}, {x: 2, y: 2}, {x: 3, y: 3}, {x: 4, y: 4}, {x: 5, y: 5}, {x: 6, y: 6}, {x: 7, y: 7}, {x: 8, y: 8}},
		},
		{
			name: "test10",
			args: args{start: points{x: 5, y: 5}, end: points{x: 8, y: 2}},
			want: []points{{x: 5, y: 5}, {x: 6, y: 4}, {x: 7, y: 3}, {x: 8, y: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPoints(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointsOverLap(t *testing.T) {
	var allLines = []lines{line1, line2, line3, line4, line5, line6, line7, line8, line9, line10}

	type args struct {
		input []lines
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{input: allLines},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointsOverLap(tt.args.input, false); got != tt.want {
				t.Errorf("PointsOverLap() = %v, want %v", got, tt.want)
			}
		})
	}
}
