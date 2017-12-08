package day7

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
		wantResult string
	}{
		{"Ex. 1", args{"pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\n" +
			"ktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\n" +
			"padx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\n" +
			"jptl (61)\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\n" +
			"cntj (57)"}, "tknk"},
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
		{"Ex. 1", args{"pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\n" +
			"ktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\n" +
			"padx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\n" +
			"jptl (61)\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\n" +
			"cntj (57)"}, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartTwo(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartTwo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
