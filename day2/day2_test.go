package day2

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
		{"Ex. 1", args{"5 1 9 5\n7 5 3\n2 4 6 8"}, 18},
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
		{"Ex. 1", args{"5 9 2 8\n9 4 7 3\n3 8 6 5"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartTwo(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartTwo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
