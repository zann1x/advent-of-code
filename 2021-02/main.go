package main

import (
	"github.com/zann1x/advent-of-code/util"
	"log"
	"strings"
)

func main() {
	inputFile := "2021-02/input.txt"

	res1 := part1(inputFile)
	log.Println("result of part1:", res1)

	res2 := part2(inputFile)
	log.Println("result of part2:", res2)
}

func part1(inputFile string) int {
	defer util.Track("part1")()

	input := util.LoadFileContentsIntoStringArray(inputFile)

	horizontalPos := 0
	depth := 0
	for _, val := range input {
		course := strings.Split(val, " ")
		direction := course[0]
		directionVal := util.ConvertStringToInt(course[1])
		switch direction {
		case "forward":
			horizontalPos += directionVal
			break
		case "up":
			depth -= directionVal
			break
		case "down":
			depth += directionVal
			break
		}
	}

	return horizontalPos * depth
}

func part2(inputFile string) int {
	defer util.Track("part2")()

	input := util.LoadFileContentsIntoStringArray(inputFile)

	aim := 0
	horizontalPos := 0
	depth := 0
	for _, val := range input {
		course := strings.Split(val, " ")
		direction := course[0]
		directionVal := util.ConvertStringToInt(course[1])
		switch direction {
		case "forward":
			horizontalPos += directionVal
			depth += aim * directionVal
			break
		case "up":
			aim -= directionVal
			break
		case "down":
			aim += directionVal
			break
		}
	}

	return horizontalPos * depth
}
