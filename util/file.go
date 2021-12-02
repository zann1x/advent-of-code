package util

import (
	"bufio"
	"os"
)

func LoadFileContentsIntoIntArray(path string) []int {
	strArr := LoadFileContentsIntoStringArray(path)
	var intArr []int
	for _, strVal := range strArr {
		intVal := ConvertStringToInt(strVal)
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
		i := ConvertStringToInt(scanner.Text())
		input[i] = true
	}

	return input
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
}
