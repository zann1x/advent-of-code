package main

import "testing"

var inputFile = "input_test.txt"

func TestPart1(t *testing.T) {
	want := 150
	if got := part1(inputFile); got != want {
		t.Error("got", got, "wanted", want)
	}
}

func TestPart2(t *testing.T) {
	want := 900
	if got := part2(inputFile); got != want {
		t.Error("got", got, "wanted", want)
	}
}
