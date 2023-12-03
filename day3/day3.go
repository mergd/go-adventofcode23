package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type coord struct {
	col int
	row int
}

func findCoord(input string, lineLength int) []coord {
	var coords []coord
	for i := 0; i < len(input); i++ {
		if i == '*' {
			y := i / lineLength
			x := i % lineLength
			coords = append(coords, coord{x, y})
		}
	}

	return coords
}

func determineCoord(input string, lineLength int, coords []coord) int {
	var runningTotal int
	var partNum string
	var partCoords []coord
	for i, char := range input {
		if unicode.IsDigit(char) {
			println("isDigit", char)
			partNum += string(char)
			y := i / lineLength
			x := i % lineLength
			partCoords = append(partCoords, coord{x, y})

		} else if len(partNum) > 0 {
			println("pushing", partNum)

			for _, point := range partCoords {
				valid := testNeighbors(point, coords)

				if valid {
					val, _ := strconv.Atoi(partNum)
					// println("is valid", val)
					runningTotal += val
					println("runningTotal", runningTotal)
					break
				}
			}
			partNum = ""
		} else {
			continue
		}
	}

	return runningTotal
}

func testNeighbors(point coord, coords []coord) bool {
	var valid = false
	println("testing", point.col, point.row)
	valid = testPoint(coord{point.col + 1, point.col}, coords) || valid
	valid = testPoint(coord{point.col + 1, point.col + 1}, coords) || valid
	valid = testPoint(coord{point.col - 1, point.col}, coords) || valid
	valid = testPoint(coord{point.col - 1, point.col - 1}, coords) || valid
	valid = testPoint(coord{point.col, point.col + 1}, coords) || valid
	valid = testPoint(coord{point.col - 1, point.col + 1}, coords) || valid
	valid = testPoint(coord{point.col + 1, point.col - 1}, coords) || valid
	valid = testPoint(coord{point.col, point.col - 1}, coords) || valid
	println(" point valid? ", valid)
	return valid

}

func testPoint(point coord, coords []coord) bool {
	for _, c := range coords {
		if c == point {
			return true
		}
	}
	return false
}

func main() {
	var sum int
	var lineLength int
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	filstr := string(file)

	//  Clean symbols and replace with *
	filstr = strings.Replace(filstr, "#", "*", -1)
	filstr = strings.Replace(filstr, "+", "*", -1)
	filstr = strings.Replace(filstr, "$", "*", -1)

	lineLength = len(strings.Split(filstr, "\n")[0])
	print(strings.Split(filstr, "\n")[0])
	println("lineLength", lineLength)
	coords := findCoord(filstr, lineLength)
	println("coords", coords[0].col, coords[0].row)

	sum = determineCoord(filstr, lineLength, coords)
	println("sum", sum)

}
