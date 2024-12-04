package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func MullItOver(corruptedInput string) (result int) {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	for _, match := range re.FindAllString(corruptedInput, -1) {
		numbers := strings.Split(strings.Replace(match[4:], ")", "", -1), ",")
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		result += num1 * num2
	}
	return
}

func MullItOverP2(corruptedInput string) (result int) {
	re := regexp.MustCompile(`(?m)do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	do := true
	for _, location := range re.FindAllStringIndex(corruptedInput, -1) {
		found := corruptedInput[location[0]:location[1]]
		if found == "don't()" {
			do = false
		}
		if found == "do()" {
			do = true
		}
		if strings.HasPrefix(found, "mul") &&
			do {
			num1, num2 := parseMul(found)
			result += num1 * num2
		}
	}
	return
}

func parseMul(match string) (int, int) {
	numbers := strings.Split(strings.Replace(match[4:], ")", "", -1), ",")
	num1, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Fatal(err)
	}
	num2, err := strconv.Atoi(numbers[1])
	if err != nil {
		log.Fatal(err)
	}
	return num1, num2
}
