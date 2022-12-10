package main

import (
	"fmt"
	"internal/filereader"
)

func Day11() {
	fmt.Println("============== Day11 ==============")
	Day11Part1()
	Day11Part2()
}

func Day11Part1() {
	fmt.Println("============== Part1 =============")
	instructions := filereader.ReadFile("./input/day-11/test.txt")

	signalStrength := readInstructions(instructions, 20, 40)

	if signalStrength != 13140 {
		fmt.Println("Test went wrong:", signalStrength)
	} else {
		fmt.Println("Test passed:", signalStrength)
	}

	instructions = filereader.ReadFile("./input/day-11/input.txt")
	signalStrength = readInstructions(instructions, 20, 40)

	fmt.Println("Output:", signalStrength)
}

func Day11Part2() {
	fmt.Println("============== Part2 =============")
	instructions := filereader.ReadFile("./input/day-11/test.txt")
	readImageInstructions(instructions, 40)

	fmt.Println("============ SOLUTION ============")

	instructions = filereader.ReadFile("./input/day-11/input.txt")
	readImageInstructions(instructions, 40)
}
