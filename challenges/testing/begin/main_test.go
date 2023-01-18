// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want 	int
	}{
		"empty": {input: "", want: 0},
		"one": {input: "a2 32 pepe", want: 5},
		"two": {input: "8/1 abcdefgh 22", want: 8},
	}

	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := lc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = map[string]struct {
		input	string
		want	int
	}{
		"empty": {input: "", want: 0},
		"one": {input: "asf1325", want: 4},
		"two": {input: "asf151asfv215", want: 6},
	}

	nc := numberCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := nc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCoun(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one": {input: "asf*/sa.1325", want: 3},
		"two": {input: "a}sf151a|sfv2`1`5", want: 4},
	}

	sc := symbolCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := sc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
