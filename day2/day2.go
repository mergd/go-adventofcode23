package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var sum int
	var powerCubeSum int
	file, err := os.ReadFile("inputday2.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileStr := string(file)

	for iterator, gamestr := range strings.Split(fileStr, "\n") {
		println(iterator + 1)
		strs := strings.Split(gamestr, ":")
		var validGame bool = true

		valid, powercube := processGame(strs[1])
		powerCubeSum += powercube
		if !valid && validGame {
			println("game knocked out", iterator)
			validGame = false
		} else if validGame {
			sum += iterator + 1
			println("game added", sum)
		}
	}

	println(sum)
	println(powerCubeSum)

}

func processGame(input string) (bool, int) {

	// println(input)
	var isValid bool = true
	var cubeNum = [3]int{0, 0, 0}
	for _, round := range strings.Split(input, ";") {

		for _, cube := range strings.Split(round, ",") {
			// println("Cube ", cube)
			var numb int
			re := regexp.MustCompile("[0-9]+")
			match := re.FindString(cube)
			if match != "" {
				num, err := strconv.Atoi(match)
				if err != nil {
					log.Fatal(err)
				}
				numb = num
			}

			if strings.Contains(cube, "red") {
				if cubeNum[0] < numb {
					cubeNum[0] = numb
				}
			} else if strings.Contains(cube, "green") {
				if cubeNum[1] < numb {
					cubeNum[1] = numb
				}

			} else if strings.Contains(cube, "blue") {
				if cubeNum[2] < numb {
					cubeNum[2] = numb
				}

			}
		}
	}
	cubenum := cubeNum[0] * cubeNum[1] * cubeNum[2]
	println("cubenum", cubeNum[0], cubeNum[1], cubeNum[2], cubenum)

	return isValid, cubenum

}
