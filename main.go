package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateLine(line string) int {
	var firstInt int
	var lastInt int
	// grep every mention of number with int

	for _, char := range strings.Split(line, "") {
		i, err := strconv.ParseInt(char, 10, 64)
		if err != nil {
			continue

		} else if firstInt == 0 {
			firstInt = int(i)
		} else {
			lastInt = int(i)
		}
	}

	if lastInt == 0 && firstInt > 0 {
		lastInt = firstInt
	}
	fmt.Println(firstInt*10 + lastInt)
	return firstInt*10 + lastInt

}

func main() {

	type Tuple struct {
		str string
		num int
	}

	var strs = []Tuple{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}, {"five", 5}, {"six", 6}, {"seven", 7}, {"eight", 8}, {"nine", 9}}

	var runningTotal int = 0
	puzzleData, fsError := os.ReadFile("file.txt")
	if fsError != nil {
		fmt.Println("Error reading file")
		return
	}

	puzzleDataString := string(puzzleData)

	for _, num := range strs {
		puzzleDataString = strings.Replace(puzzleDataString, num.str, strconv.Itoa(num.num), -1)

	}

	for iterator, line := range strings.Split(puzzleDataString, "\n") {
		if iterator > 10 {

			print(line)
		}
		runningTotal += calculateLine(line)
	}

	println(runningTotal)

}
