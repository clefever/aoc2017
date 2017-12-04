package day4

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
		{"Ex. 1", args{"aa bb cc dd ee"}, 1},
		{"Ex. 2", args{"aa bb cc dd aa"}, 0},
		{"Ex. 3", args{"aa bb cc dd aaa"}, 1},
		{"Combined", args{"aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa"}, 2},
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
		{"Ex. 1", args{"abcde fghij"}, 1},
		{"Ex. 2", args{"abcde xyz ecdab"}, 0},
		{"Ex. 3", args{"a ab abc abd abf abj"}, 1},
		{"Ex. 4", args{"iiii oiii ooii oooi oooo"}, 1},
		{"Ex. 5", args{"oiii ioii iioi iiio"}, 0},
		{"Combined", args{"abcde fghij\nabcde xyz ecdab\na ab abc abd abf abj\n" +
			"iiii oiii ooii oooi oooo\noiii ioii iioi iiio"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartTwo(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartTwo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
