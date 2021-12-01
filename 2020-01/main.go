package main

import (
	"github.com/zann1x/advent-of-code/util"
	"log"
)

func main() {
	inputFile := "2020-01/input.txt"

	res1 := part1(inputFile)
	log.Println("result of part1:", res1)

	res2 := part2(inputFile)
	log.Println("result of part2:", res2)
}

func part1(inputFile string) int {
	defer util.Track("part1")()

	input := util.LoadFileContentsIntoIntMap(inputFile)
	for i := range input {
		j := 2020 - i
		if input[j] == false {
			continue
		}

		return i * j
	}

	panic("No result found")
}

func part2(inputFile string) int {
	defer util.Track("part2")()

	inputMap := util.LoadFileContentsIntoIntMap(inputFile)
	inputArr := util.LoadFileContentsIntoIntArray(inputFile)
	for i := 0; i < len(inputArr); i++ {
		for j := i + 1; j < len(inputArr); j++ {
			k := 2020 - inputArr[i] - inputArr[j]
			if inputMap[k] == false {
				continue
			}

			return inputArr[i] * inputArr[j] * k
		}
	}

	panic("No result found")
}
