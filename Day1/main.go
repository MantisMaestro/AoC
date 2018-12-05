package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("E:/Development/Go/AoC/Day1/tmp/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	changeList := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineInt, _ := strconv.Atoi(line)
		changeList = append(changeList, lineInt)
	}

	seenFrequencies := make(map[int]bool)
	duplicateFound := false
	frequency := 0

	for duplicateFound == false {
		for _,change := range changeList {
			frequency += change
			if seenFrequencies[frequency] == false {
				seenFrequencies[frequency] = true
			} else {
				fmt.Printf("Already seen frequency %v\n", frequency)
				duplicateFound = true
				break
			}
		}
	}
}
