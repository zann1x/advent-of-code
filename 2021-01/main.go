package main

import (
	"github.com/zann1x/advent-of-code/util"
	"log"
	"math"
)

func main() {
	inputFile := "2021-01/input.txt"

	res1 := part1(inputFile)
	log.Println("result of part1:", res1)

	res2 := part2(inputFile)
	log.Println("result of part2:", res2)
}

func part1(inputFile string) int {
	defer util.Track("part1")()

	input := util.LoadFileContentsIntoIntArray(inputFile)

	prevVal := math.MaxInt
	resultCount := 0
	for _, val := range input {
		if val > prevVal {
			resultCount++
		}

		prevVal = val
	}

	return resultCount
}

func part2(inputFile string) int {
	defer util.Track("part2")()

	input := util.LoadFileContentsIntoIntArray(inputFile)

	prevSum := math.MaxInt
	resultCount := 0
	for i := 0; i < len(input)-2; i++ {
		j := i + 1
		k := i + 2
		sum := input[i] + input[j] + input[k]
		if sum > prevSum {
			resultCount++
		}

		prevSum = sum
	}

	return resultCount
}
