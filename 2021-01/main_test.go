package main

import "testing"

func TestParts(t *testing.T) {
	inputFile := "input_test.txt"

	if got := part1(inputFile); got != 7 {
		t.Error("got", got, "wanted", 7)
	}

	if got := part2(inputFile); got != 5 {
		t.Error("got", got, "wanted", 5)
	}
}
