package giantsquid

import (
	"reflect"
	"testing"
)

var grid3 = [][]int{{14, 21, 17, 24, 4}, {10, 16, 15, 9, 19}, {18, 8, 23, 26, 20}, {22, 11, 13, 6, 5}, {2, 0, 12, 3, 7}}
var grid1 = [][]int{{22, 13, 17, 11, 0}, {8, 2, 23, 4, 24}, {21, 9, 14, 16, 7}, {6, 10, 3, 18, 5}, {1, 12, 20, 15, 19}}
var grid2 = [][]int{{3, 15, 0, 2, 22}, {9, 18, 13, 17, 5}, {19, 8, 7, 25, 23}, {20, 11, 10, 24, 4}, {14, 21, 16, 12, 6}}

var seq1 = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

var output1 = [5][5]int{{1, 1, 1, 1, 1}, {0, 0, 0, 1, 0}, {0, 0, 1, 0, 0}, {0, 1, 0, 0, 1}, {1, 1, 0, 0, 1}}

var combinedInput = [][][]int{grid1, grid2, grid3}

// func TestMatch(t *testing.T) {
// 	type args struct {
// 		grid [5][5]int
// 		seq  []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want [5][5]int
// 	}{
// 		{
// 			name: "test case 1",
// 			args: args{grid1, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}},
// 			want: output1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Match(tt.args.grid, tt.args.seq); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Match() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestCheckRowDone(t *testing.T) {
	type args struct {
		grid [5][5]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{output1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckRowDone(tt.args.grid); got != tt.want {
				t.Errorf("CheckRowDone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckColumnDone(t *testing.T) {
	type args struct {
		grid [5][5]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{output1},
			want: false,
		},
		{
			name: "test2",
			args: args{[5][5]int{{1, 0, 1, 1, 1}, {1, 0, 0, 1, 0}, {1, 0, 1, 0, 0}, {1, 1, 0, 0, 1}, {1, 1, 0, 0, 1}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckColumnDone(tt.args.grid); got != tt.want {
				t.Errorf("CheckColumnDone() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestBingo(t *testing.T) {
// 	type args struct {
// 		input [5][5]int
// 		seq   []int
// 	}
// 	tests := []struct {
// 		name        string
// 		args        args
// 		matchOutput [5][5]int
// 		grid        [5][5]int
// 		number      int
// 	}{
// 		{
// 			name:        "test case 1",
// 			args:        args{grid1, seq1},
// 			matchOutput: output1,
// 			grid:        grid1,
// 			number:      24,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1, got2, got3, got4 := Bingo(tt.args.input, tt.args.seq)
// 			if !reflect.DeepEqual(got, tt.matchOutput) {
// 				t.Errorf("Bingo() got = %v, want %v", got, tt.matchOutput)
// 			}
// 			if !reflect.DeepEqual(got1, tt.grid) {
// 				t.Errorf("Bingo() got1 = %v, want %v", got1, tt.grid)
// 			}
// 			if got2 != tt.number {
// 				t.Errorf("Bingo() got2 = %v, want %v", got2, tt.number)
// 			}
// 		})
// 	}
// }

func TestBingo(t *testing.T) {
	type args struct {
		grid [][]int
		seq  []int
	}
	tests := []struct {
		name  string
		args  args
		want  [5][5]int
		want1 [][]int
		want2 int
		want3 bool
		want4 []int
	}{
		{
			name:  "test case 1",
			args:  args{grid1, seq1},
			want:  output1,
			want1: grid1,
			want2: 24,
			want3: true,
			want4: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4 := Bingo(tt.args.grid, tt.args.seq)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bingo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Bingo() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Bingo() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("Bingo() got3 = %v, want %v", got3, tt.want3)
			}
			if !reflect.DeepEqual(got4, tt.want4) {
				t.Errorf("Bingo() got4 = %v, want %v", got4, tt.want4)
			}
		})
	}
}

func TestRunAllBingo(t *testing.T) {
	type args struct {
		inputs [][][]int
		seq    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{combinedInput, seq1},
			want: 4512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunAllBingo(tt.args.inputs, tt.args.seq); got != tt.want {
				t.Errorf("RunAllBingo() = %v, want %v", got, tt.want)
			}
		})
	}
}
