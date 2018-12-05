package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input/day2.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	inputStrings := make([]string, 0)

	var twos, threes = 0, 0

	for scanner.Scan() {
		inputStrings = append(inputStrings, scanner.Text())
	}

	for _, s := range inputStrings {
		two, three := checkRepeatingChars(s)
		if two { twos++ }
		if three { threes++ }
	}

	for i, s := range inputStrings {
		for j := i+1; j < len(inputStrings); j++ {
			similarity := noOfDifferencesBetweenStrings(s, inputStrings[j])
			if similarity == 1 {
				fmt.Printf("S1: %s, S2: %s\n", s, inputStrings[j])
			}
		}
	}

	id := twos*threes
	fmt.Printf("Twos: %v, Threes: %v, ID: %v\n", twos, threes, id)
}

func checkRepeatingChars(input string) (bool, bool) {
	charMap := make(map[rune]int)

	for _, c := range input {
		charMap[c] += 1
	}

	var twos, threes = false, false

	for _, v := range charMap {
		if v == 2 {
			twos = true
		}
		if v == 3 {
			threes = true
		}
	}
	return twos, threes
}

func noOfDifferencesBetweenStrings(s1 string, s2 string) int {
	differences := 0

	for i,c1 := range s1 {
		if uint8(c1) != s2[i] {
			differences++
		}
	}

	return differences
}
