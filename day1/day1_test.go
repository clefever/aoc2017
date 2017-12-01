package day1

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Ex. 1", args{"1122"}, 3},
		{"Ex. 2", args{"1111"}, 4},
		{"Ex. 3", args{"1234"}, 0},
		{"Ex. 4", args{"91212129"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.args.input); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Ex. 1", args{"1212"}, 6},
		{"Ex. 2", args{"1221"}, 0},
		{"Ex. 3", args{"123425"}, 4},
		{"Ex. 4", args{"123123"}, 12},
		{"Ex. 5", args{"12131415"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.input); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
