package main

import (
	"fmt"
	"internal/filereader"
)

func Day6() {
	fmt.Println("============== Day6 ==============")
	Day5Part1()
	Day5Part2()
}

func Day6Part1() {
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

func Day6Part2() {
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
