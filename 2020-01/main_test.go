package main

import "testing"

func TestParts(t *testing.T) {
	inputFile := "input_test.txt"

	if got := part1(inputFile); got != 514579 {
		t.Error("got", got, "wanted", 514579)
	}

	if got := part2(inputFile); got != 241861950 {
		t.Error("got", got, "wanted", 241861950)
	}
}
