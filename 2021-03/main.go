package main

import (
	"github.com/zann1x/advent-of-code/util"
	"log"
)

func main() {
	inputFile := "2021-03/input.txt"

	res1 := part1(inputFile)
	log.Println("result of part1:", res1)

	res2 := part2(inputFile)
	log.Println("result of part2:", res2)
}

func part1(inputFile string) int {
	defer util.Track("part1")()

	input := util.LoadFileContentsIntoStringArray(inputFile)

	horizontalOneCount := make([]int, len(input[0]))
	for _, val := range input {
		for pos, char := range val {
			if string(char) == "1" {
				horizontalOneCount[pos]++
			}
		}
	}

	gamma := uint64(0b0)
	epsilon := uint64(0b0)
	for _, oneCnt := range horizontalOneCount {
		if oneCnt > len(input)/2 {
			gamma = gamma<<1 + 1
		} else {
			gamma = gamma << 1
		}

		epsilon = epsilon<<1 + 1
	}

	epsilon = gamma ^ epsilon

	return int(gamma) * int(epsilon)
}

func part2(inputFile string) int {
	defer util.Track("part2")()

	input := util.LoadFileContentsIntoStringArray(inputFile)

	oxygenCandidates := input
	for pos := 0; len(oxygenCandidates) > 1; pos++ {
		oneCnt := countOnesAtPos(oxygenCandidates, pos)
		if float64(oneCnt) >= float64(len(oxygenCandidates))/2 {
			oxygenCandidates = getEntriesWithValueAtPos(oxygenCandidates, '1', pos)
		} else {
			oxygenCandidates = getEntriesWithValueAtPos(oxygenCandidates, '0', pos)
		}
	}
	oxygen := convertBinaryStringToInt(oxygenCandidates[0])

	co2Candidates := input
	for pos := 0; len(co2Candidates) > 1; pos++ {
		oneCnt := countOnesAtPos(co2Candidates, pos)
		if float64(oneCnt) >= float64(len(co2Candidates))/2 {
			co2Candidates = getEntriesWithValueAtPos(co2Candidates, '0', pos)
		} else {
			co2Candidates = getEntriesWithValueAtPos(co2Candidates, '1', pos)
		}
	}
	co2 := convertBinaryStringToInt(co2Candidates[0])

	return oxygen * co2
}

func countOnesAtPos(arr []string, pos int) int {
	cnt := 0
	for _, val := range arr {
		if val[pos] == '1' {
			cnt++
		}
	}

	return cnt
}

func getEntriesWithValueAtPos(arr []string, value byte, pos int) []string {
	var res []string
	for _, entry := range arr {
		if entry[pos] == value {
			res = append(res, entry)
		}
	}

	return res
}

func convertBinaryStringToInt(s string) int {
	result := int64(0b0)
	for _, val := range s {
		result = result << 1
		if val == '1' {
			result++
		}
	}

	return int(result)
}
