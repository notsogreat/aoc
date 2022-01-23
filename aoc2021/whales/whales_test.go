package whales

import (
	"testing"
)

func Test_findMinAndMax(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		min  int
		max  int
	}{
		{
			name: "test1",
			args: args{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}},
			min:  1,
			max:  16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findMinAndMax(tt.args.values)
			if got != tt.min {
				t.Errorf("findMinAndMax() got = %v, want %v", got, tt.min)
			}
			if got1 != tt.max {
				t.Errorf("findMinAndMax() got1 = %v, want %v", got1, tt.max)
			}
		})
	}
}

func Test_calculateFuel(t *testing.T) {
	type args struct {
		positions []int
		value     int
		con       bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 2, true},
			want: 37,
		},
		{
			name: "test2",
			args: args{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 0, true},
			want: 49,
		},
		{
			name: "test3",
			args: args{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 16, true},
			want: 111,
		},
		{
			name: "test4",
			args: args{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 1, true},
			want: 41,
		},
		{
			name: "test5",
			args: args{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 5, false},
			want: 168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFuel(tt.args.positions, tt.args.value, tt.args.con); got != tt.want {
				t.Errorf("calculateFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alignCrabs(t *testing.T) {
	type args struct {
		values []int
		cons   bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, true},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alignCrabs(tt.args.values, tt.args.cons); got != tt.want {
				t.Errorf("alignCrabs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exponential(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{v: 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exponential(tt.args.v); got != tt.want {
				t.Errorf("exponential() = %v, want %v", got, tt.want)
			}
		})
	}
}
