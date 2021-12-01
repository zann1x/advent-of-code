package util

import (
	"bufio"
	"os"
	"strconv"
)

func LoadFileContentsIntoIntArray(path string) []int {
	strArr := LoadFileContentsIntoStringArray(path)
	var intArr []int
	for _, strVal := range strArr {
		var intVal int
		var err error
		if intVal, err = strconv.Atoi(strVal); err != nil {
			panic(err)
		}
		intArr = append(intArr, intVal)
	}

	return intArr
}

func LoadFileContentsIntoStringArray(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer closeFile(file)

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func LoadFileContentsIntoIntMap(path string) map[int]bool {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer closeFile(file)

	input := map[int]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		input[i] = true
	}

	return input
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
}
