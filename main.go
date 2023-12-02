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
	var runningTotal int = 0
	puzzleData, fsError := os.ReadFile("file.txt")
	if fsError != nil {
		fmt.Println("Error reading file")
		return
	}

	puzzleDataString := string(puzzleData)

	for _, line := range strings.Split(puzzleDataString, "\n") {
		runningTotal += calculateLine(line)
	}

	println(runningTotal)

}
