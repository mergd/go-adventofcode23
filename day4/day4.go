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
	pattern := regexp.MustCompile(`Card .*: (.*) \| (.*)`)
	match := pattern.FindStringSubmatch(str)

	_, winners, card := match[0], match[1], match[2]

	println(winners)
	println(card)

	m := make(map[int]bool)
	var matchCount int

	for _, strnum := range strings.Split(winners, " ") {
		num, _ := strconv.Atoi(strnum)
		print(num, " ")
		m[num] = true
	}
	println("parsecard")
	for _, cardnum := range strings.Split(card, " ") {
		num, err := strconv.Atoi(cardnum)
		if err != nil {
			println(err.Error())
		}

		print(num, " ")

		if m[num] && num != 0 {
			fmt.Printf("Match Found:  %d \n", num)

			matchCount++
		}

	}
	println("card")
	println("matchCount", matchCount)
	return int(math.Pow(2, float64(matchCount-1)))

}

func part2ParseCard(mapping map[int]int, countMapping map[int]int, str string) int {
	pattern := regexp.MustCompile(`(\d{1,3}): (.*) \| (.*)`)
	match := pattern.FindStringSubmatch(str)

	_, cardNumber, winners, card := match[0], match[1], match[2], match[3]
	num1, _ := strconv.Atoi(cardNumber)
	fmt.Printf("Game %d \n", num1)
	println(winners)
	println(card)
	m := make(map[int]bool)

	var matchCount int

	for _, strnum := range strings.Split(winners, " ") {
		num, _ := strconv.Atoi(strnum)

		m[num] = true
	}

	for _, cardnum := range strings.Split(card, " ") {
		num, err := strconv.Atoi(cardnum)
		if err != nil {
			println(err.Error())
		}

		if m[num] && num != 0 {

			matchCount++
		}

	}

	countMapping[num1] += 1
	val := countMapping[num1]
	for i := 0; i <= matchCount; i++ {
		countMapping[num1+i] += val
		fmt.Printf("Card %d now has the value %d \n", num1+i, countMapping[num1+i])
	}

	fmt.Printf("matchCount %d for match %d, but there are %d instances \n", matchCount, num1, countMapping[num1])
	return (countMapping[num1] / 2)
	// return int(math.Pow(2, float64(matchCount-1)))

}

func main() {
	var runningTotal int

	file, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	filstr := string(file)

	mapping := make(map[int]int)
	countMapping := make(map[int]int)

	for _, str := range strings.Split(filstr, "\n") {
		// runningTotal += parseCard(str)
		runningTotal += part2ParseCard(mapping, countMapping, str)
	}

	println("runningTotal", runningTotal)

}
