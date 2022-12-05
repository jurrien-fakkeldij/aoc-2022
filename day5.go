package main

import (
	"fmt"
	"internal/filereader"
	"sort"
	"strconv"
	"strings"
)

func Day5() {
	fmt.Println("============== Day5 ==============")
	Day5Part1()
	Day5Part2()
}

func Day5Part1() {
	fmt.Println("============== Part1 =============")
	cargoAndProcedures := filereader.ReadFile("./input/day-5/test.txt")

	items := findTopItems(cargoAndProcedures)

	if items != "CMZ" {
		fmt.Println("Test went wrong:", items)
	} else {
		fmt.Println("Test passed:", items)
	}

	cargoAndProcedures = filereader.ReadFile("./input/day-5/input.txt")
	items = findTopItems(cargoAndProcedures)

	fmt.Println("Output:", items)
}

func findTopItems(cargoAndProcedures []string) string {
	cargo, procedures, stacks := setupCargoAndProcedures(cargoAndProcedures)
	//fmt.Println(cargo)
	for _, procedure := range procedures {
		tempSlice := strings.Split(procedure, " ")
		//fmt.Println(procedure)
		amount, _ := strconv.Atoi(tempSlice[1])
		fromStack, _ := strconv.Atoi(tempSlice[3])
		toStack, _ := strconv.Atoi(tempSlice[5])

		for i := 0; i < amount; i++ {
			index := len(cargo[fromStack]) - 1               // Get the index of the top most element.
			element := (cargo[fromStack])[index]             // Index into the slice and obtain the element.
			cargo[fromStack] = (cargo[fromStack])[:index]    // Remove it from the stack by slicing it off.
			cargo[toStack] = append(cargo[toStack], element) // Add it to new stack
		}

		//fmt.Println(cargo)
	}
	topItems := ""
	//fmt.Println(cargo)
	sort.Ints(stacks)
	//fmt.Println(stacks)
	for _, stack := range stacks {
		topItems += cargo[stack][len(cargo[stack])-1]
	}
	return topItems
}

func setupCargoAndProcedures(cargoAndProcedures []string) (map[int][]string, []string, []int) {
	procedures := make([]string, 0)
	cargo := make(map[int][]string)
	stacks := make([]int, 0)

	readingCargo := true
	for _, cargoOrProcedure := range cargoAndProcedures {
		if len(cargoOrProcedure) == 0 {
			readingCargo = false
			continue
		}

		if readingCargo {
			currentBucket := 1
			for i := 0; i < len(cargoOrProcedure); i += 4 {
				if cargoOrProcedure[i] == '[' {
					if _, ok := cargo[currentBucket]; !ok {
						cargo[currentBucket] = make([]string, 0)
						stacks = append(stacks, currentBucket)
					}
					cargo[currentBucket] = append([]string{string(cargoOrProcedure[i+1])}, cargo[currentBucket]...)
					currentBucket++
				} else {
					currentBucket++
					continue
				}

			}
			continue
		}

		if strings.Contains(cargoOrProcedure, "move") {
			procedures = append(procedures, cargoOrProcedure)
		}
	}
	return cargo, procedures, stacks
}

func Day5Part2() {
	fmt.Println("============== Part2 =============")
	cargoAndProcedures := filereader.ReadFile("./input/day-5/test.txt")

	items := findTopItems9001(cargoAndProcedures)

	if items != "MCD" {
		fmt.Println("Test went wrong:", items)
	} else {
		fmt.Println("Test passed:", items)
	}

	cargoAndProcedures = filereader.ReadFile("./input/day-5/input.txt")
	items = findTopItems9001(cargoAndProcedures)

	fmt.Println("Output:", items)
}

func findTopItems9001(cargoAndProcedures []string) string {
	cargo, procedures, stacks := setupCargoAndProcedures(cargoAndProcedures)
	//fmt.Println(cargo)
	for _, procedure := range procedures {
		tempSlice := strings.Split(procedure, " ")
		//fmt.Println(procedure)
		amount, _ := strconv.Atoi(tempSlice[1])
		fromStack, _ := strconv.Atoi(tempSlice[3])
		toStack, _ := strconv.Atoi(tempSlice[5])

		index := len(cargo[fromStack]) - (amount)            // Get the index of the top most element.
		elements := (cargo[fromStack])[index:]               // Index into the slice and obtain the element.
		cargo[fromStack] = (cargo[fromStack])[:index]        // Remove it from the stack by slicing it off.
		cargo[toStack] = append(cargo[toStack], elements...) // Add it to new stack

		//fmt.Println(cargo)
	}
	topItems := ""
	//fmt.Println(cargo)
	sort.Ints(stacks)
	//fmt.Println(stacks)
	for _, stack := range stacks {
		topItems += cargo[stack][len(cargo[stack])-1]
	}
	return topItems
}
