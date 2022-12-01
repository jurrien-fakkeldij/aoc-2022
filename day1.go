package main

import (
	"fmt"
	"internal/filereader"
	"sort"
	"strconv"
)

func Day1() {
	fmt.Println("============== Day1 ==============")
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("============== Part1 =============")
	caloriesList := filereader.ReadFile("./input/day-1/test.txt")
	_, amount := maxCalories(caloriesList)

	if amount != 24000 {
		fmt.Println("Test went wrong:", amount)
	} else {
		fmt.Println("Test passed:", amount)
	}

	caloriesList = filereader.ReadFile("./input/day-1/input.txt")
	_, amount = maxCalories(caloriesList)
	fmt.Println("Output:", amount)
}

func maxCalories(allCalories []string) (int, int) {
	elf := 1
	max := 0
	currentCalories := 0
	currentElf := 1

	for _, calorie := range allCalories {
		if len(calorie) > 0 {
			cal, _ := strconv.Atoi(calorie)
			currentCalories += cal
		} else {
			if max < currentCalories {
				max = currentCalories
				elf = currentElf
			}
			currentElf++
			currentCalories = 0
		}
	}

	if max < currentCalories {
		max = currentCalories
		elf = currentElf
	}
	currentElf++
	currentCalories = 0

	return elf, max
}

func Part2() {
	fmt.Println("============== Part2 =============")
	caloriesList := filereader.ReadFile("./input/day-1/test.txt")
	amount := top3Callories(caloriesList)

	if amount != 45000 {
		fmt.Println("Test went wrong:", amount)
	} else {
		fmt.Println("Test passed:", amount)
	}

	caloriesList = filereader.ReadFile("./input/day-1/input.txt")
	amount = top3Callories(caloriesList)
	fmt.Println("Output:", amount)
}

func top3Callories(allCalories []string) int {
	currentCalories := 0
	sortedCalories := []int{}
	for _, calorie := range allCalories {
		if len(calorie) > 0 {
			cal, _ := strconv.Atoi(calorie)
			currentCalories += cal
		} else {
			sortedCalories = append(sortedCalories, currentCalories)
			currentCalories = 0
		}
	}
	sortedCalories = append(sortedCalories, currentCalories)
	currentCalories = 0

	sort.Sort(sort.Reverse(sort.IntSlice(sortedCalories)))
	return (sortedCalories[0] + sortedCalories[1] + sortedCalories[2])
}
