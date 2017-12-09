package day9

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
		{"Ex. 1", args{"{}"}, 1},
		{"Ex. 2", args{"{{{}}}"}, 6},
		{"Ex. 3", args{"{{},{}}"}, 5},
		{"Ex. 4", args{"{{{},{},{{}}}}"}, 16},
		{"Ex. 5", args{"{<a>,<a>,<a>,<a>}"}, 1},
		{"Ex. 6", args{"{{<ab>},{<ab>},{<ab>},{<ab>}}"}, 9},
		{"Ex. 7", args{"{{<!!>},{<!!>},{<!!>},{<!!>}}"}, 9},
		{"Ex. 8", args{"{{<a!>},{<a!>},{<a!>},{<ab>}}"}, 3},
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
		{"Ex. 1", args{"<>"}, 0},
		{"Ex. 2", args{"<random characters>"}, 17},
		{"Ex. 3", args{"<<<<>"}, 3},
		{"Ex. 4", args{"<{!>}>"}, 2},
		{"Ex. 5", args{"<!!>"}, 0},
		{"Ex. 6", args{"<!!!>>"}, 0},
		{"Ex. 7", args{"<{o\"i!a,<{i<a>"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartTwo(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartTwo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
