package main

import "testing"

func TestParts(t *testing.T) {
	inputFile := "input_test.txt"

	if got := part1(inputFile); got != 2 {
		t.Error("got", got, "wanted", 2)
	}

	if got := part2(inputFile); got != 1 {
		t.Error("got", got, "wanted", 1)
	}
}
