package lanternfish

import (
	"testing"
)

func TestCalculateNumberOfFish(t *testing.T) {
	type args struct {
		InitialState []int
		days         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{InitialState: []int{3, 4, 3, 1, 2}, days: 18},
		// 	want: 26,
		// },
		// {
		// 	name: "test2",
		// 	args: args{InitialState: []int{3, 4, 3, 1, 2}, days: 80},
		// 	want: 5934,
		// },
		{
			name: "test3",
			args: args{InitialState: []int{3, 4, 3, 1, 2}, days: 256},
			want: 26984457539,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateNumberOfFish(tt.args.InitialState, tt.args.days); got != tt.want {
				t.Errorf("CalculateNumberOfFish() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestCalculateNumberOfFish1(t *testing.T) {
// 	type args struct {
// 		InitialState []int
// 		days         int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			name: "test1",
// 			args: args{InitialState: []int{3, 4, 3, 1, 2}, days: 18},
// 			want: 26,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := CalculateNumberOfFish1(tt.args.InitialState, tt.args.days); got != tt.want {
// 				t.Errorf("CalculateNumberOfFish1() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_internalTimerCheck1(t *testing.T) {
// 	type args struct {
// 		fishLife int
// 		days     int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  int
// 		want1 []int
// 	}{
// 		{
// 			name:  "test1",
// 			args:  args{fishLife: 3, days: 18},
// 			want:  6,
// 			want1: []int{8, 8, 8},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := internalTimerCheck1(tt.args.fishLife, tt.args.days)
// 			if got != tt.want {
// 				t.Errorf("internalTimerCheck1() got = %v, want %v", got, tt.want)
// 			}
// 			if !reflect.DeepEqual(got1, tt.want1) {
// 				t.Errorf("internalTimerCheck1() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

// NEW TEST
