package SonarSweep

import (
	"reflect"
	"testing"
)

func TestCalculateDepthMeasureIncrease(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Test Case 1
		{"test case 1", args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, 7},
		{"test case 2", args{[]int{607, 618, 618, 617, 647, 716, 769, 792}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDepthMeasureIncrease(tt.args.input); got != tt.want {
				t.Errorf("CountDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateDepthMeasureSlidingWindow(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name      string
		args      args
		want      []int
		wantDepth int
	}{
		{"testCase 1", args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, []int{607, 618, 618, 617, 647, 716, 769, 792}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			depthMeasureSlice, depthMeasureIncrease := CalculateDepthMeasureSlidingWindow(tt.args.input)
			if !reflect.DeepEqual(depthMeasureSlice, tt.want) {
				t.Errorf("CalculateDepthMeasureSlidingWindow() got = %v, want %v", depthMeasureSlice, tt.want)
			}
			if depthMeasureIncrease != tt.wantDepth {
				t.Errorf("CalculateDepthMeasureIncrease() depthMeasureIncrease = %v, want %v", depthMeasureIncrease, tt.wantDepth)
			}
		})
	}
}
