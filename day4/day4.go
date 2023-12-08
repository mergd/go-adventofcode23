package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseCard(str string) int {
	pattern := regexp.MustCompile(`\d*:(.*)|(.*)`)
	match := pattern.FindStringSubmatch(str)

	_, winners, card := match[0], match[1], match[2]

	m := make(map[int]bool)
	var matchCount int

	for _, strnum := range strings.Split(winners, " ") {
		print(strnum, " ")
		num, _ := strconv.Atoi(strnum)
		m[num] = true
	}
	println("parsecard")
	for _, cardnum := range strings.Split(card, " ") {
		num, _ := strconv.Atoi(cardnum)
		print(num, " ")

		if m[num] {
			fmt.Printf("Match Found:  %d", num)

			matchCount++
		}

	}
	println("card")
	println("matchCount", matchCount)
	return int(math.Pow(2, float64(matchCount-1)))

}

func main() {
	var runningTotal int

	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	filstr := string(file)

	for _, str := range strings.Split(filstr, "\n") {
		runningTotal += parseCard(str)
	}

	println("runningTotal", runningTotal)

}
