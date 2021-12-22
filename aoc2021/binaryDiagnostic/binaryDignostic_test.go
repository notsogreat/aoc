package binarydiagnostic

import (
	"reflect"
	"testing"
)

func Test_PowerConsumption(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "test case 1",
			args:    args{input: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}},
			want:    198,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PowerConsumption(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PowerConsumption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PowerConsumption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateMostCommon(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test case 1",
			args: args{input: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}},
			want: []string{"10111"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMostCommon(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateMostCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchValue(t *testing.T) {
	type args struct {
		input    []string
		value    string
		position int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test case 1",
			args: args{
				input:    []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
				value:    "1",
				position: 0,
			},
			want: []string{"11110", "10110", "10111", "10101", "11100", "10000", "11001"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchValue(tt.args.input, tt.args.value, tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateListCommon(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test case 1",
			args: args{input: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}},
			want: []string{"01010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateListCommon(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateListCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLifeSupport(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "test case 1",
			args:    args{input: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}},
			want:    230,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LifeSupport(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("LifeSupport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LifeSupport() = %v, want %v", got, tt.want)
			}
		})
	}
}
