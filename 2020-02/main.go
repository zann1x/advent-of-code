package main

import (
	"github.com/zann1x/advent-of-code/util"
	"log"
	"strconv"
	"strings"
)

type parsedLine struct {
	ruleChar byte
	ruleOne  int
	ruleTwo  int
	pass     string
}

func main() {
	inputFile := "2020-02/input.txt"

	res1 := part1(inputFile)
	log.Println("result of part1:", res1)

	res2 := part2(inputFile)
	log.Println("result of part2:", res2)
}

func part1(inputFile string) int {
	defer util.Track("part1")()

	input := util.LoadFileContentsIntoStringArray(inputFile)
	var validCnt int
	for _, line := range input {
		parsed := parse(line)

		foundCnt := strings.Count(parsed.pass, string(parsed.ruleChar))
		if foundCnt < parsed.ruleOne || foundCnt > parsed.ruleTwo {
			continue
		}

		validCnt++
	}

	return validCnt
}

func part2(inputFile string) int {
	defer util.Track("part2")()

	input := util.LoadFileContentsIntoStringArray(inputFile)
	var validCnt int
	for _, line := range input {
		parsed := parse(line)

		if len(parsed.pass) < parsed.ruleOne || len(parsed.pass) < parsed.ruleTwo {
			continue
		}

		isCharInBothPositions := parsed.pass[parsed.ruleOne-1] == parsed.ruleChar && parsed.pass[parsed.ruleTwo-1] == parsed.ruleChar
		isCharInNeitherPosition := parsed.pass[parsed.ruleOne-1] != parsed.ruleChar && parsed.pass[parsed.ruleTwo-1] != parsed.ruleChar
		if isCharInBothPositions || isCharInNeitherPosition {
			continue
		}

		validCnt++
	}

	return validCnt
}

func parse(line string) parsedLine {
	var result parsedLine

	tokens := strings.Split(line, ": ")

	rule := strings.Split(tokens[0], " ")
	result.pass = tokens[1]

	ruleCnt := strings.Split(rule[0], "-")
	ruleCharByteArr := []byte(rule[1])
	result.ruleChar = ruleCharByteArr[0]

	var err error
	result.ruleOne, err = strconv.Atoi(ruleCnt[0])
	if err != nil {
		panic(err)
	}
	result.ruleTwo, err = strconv.Atoi(ruleCnt[1])
	if err != nil {
		panic(err)
	}

	return result
}
