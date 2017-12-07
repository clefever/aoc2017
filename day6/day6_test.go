package day6

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Ex. 1", args{"0\t2\t7\t0"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartOne(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartOne() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Ex. 1", args{"0\t2\t7\t0"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartTwo(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartTwo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
