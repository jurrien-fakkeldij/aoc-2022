package main

import (
	"fmt"
	"internal/filereader"
	"strings"
	"unicode"
)

func Day3() {
	fmt.Println("============== Day3 ==============")
	Day3Part1()
	Day3Part2()
}

func Day3Part1() {
	fmt.Println("============== Part1 =============")
	inventory := filereader.ReadFile("./input/day-3/test.txt")

	score := prioritiesInventory(inventory)

	if score != 157 {
		fmt.Println("Test went wrong:", score)
	} else {
		fmt.Println("Test passed:", score)
	}

	inventory = filereader.ReadFile("./input/day-3/input.txt")
	score = prioritiesInventory(inventory)

	fmt.Println("Output:", score)
}

func prioritiesInventory(inventory []string) int {
	wrongItems := make([]rune, 0)
	for _, rucksack := range inventory {
		compartiment1 := rucksack[0 : len(rucksack)/2]
		compartiment2 := rucksack[len(rucksack)/2:]
		for _, item := range compartiment1 {
			if strings.Contains(compartiment2, string(item)) {
				wrongItems = append(wrongItems, item)
				break
			}
		}
	}
	output := 0
	for _, item := range wrongItems {
		if unicode.IsUpper(item) {
			output += int(item) - 38
		} else {
			output += int(item) - 96
		}
	}

	return output
}

func IsUpper(r rune) bool {
	if !unicode.IsUpper(r) && unicode.IsLetter(r) {
		return false
	}
	return true
}

func IsLower(r rune) bool {
	if !unicode.IsLower(r) && unicode.IsLetter(r) {
		return false
	}
	return true
}

func Day3Part2() {
	fmt.Println("============== Part2 =============")
	inventory := filereader.ReadFile("./input/day-3/test.txt")

	score := findBadges(inventory)

	if score != 70 {
		fmt.Println("Test went wrong:", score)
	} else {
		fmt.Println("Test passed:", score)
	}

	inventory = filereader.ReadFile("./input/day-3/input.txt")
	score = findBadges(inventory)
	fmt.Println("Output:", score)
}

func findBadges(inventory []string) int {
	badges := make([]rune, 0)
	for i := 0; i < len(inventory); i += 3 {
		rucksack1 := inventory[i]
		rucksack2 := inventory[i+1]
		rucksack3 := inventory[i+2]
		for _, item := range rucksack1 {
			if strings.Contains(rucksack2, string(item)) && strings.Contains(rucksack3, string(item)) {
				badges = append(badges, item)
				break
			}
		}
	}
	output := 0
	for _, item := range badges {
		if unicode.IsUpper(item) {
			output += int(item) - 38
		} else {
			output += int(item) - 96
		}
	}

	return output
}
