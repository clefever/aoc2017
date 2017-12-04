package day3

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"Ex. 1", args{1}, 0},
		{"Ex. 2", args{12}, 3},
		{"Ex. 3", args{23}, 2},
		{"Ex. 4", args{1024}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartOne(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartOne() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// TODO: Add PartTwo tests
