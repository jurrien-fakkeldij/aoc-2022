package main

import (
	"fmt"
	"internal/filereader"
)

func Day10() {
	fmt.Println("============== Day10 ==============")
	Day10Part1()
	Day10Part2()
}

func Day10Part1() {
	fmt.Println("============== Part1 =============")
	instructions := filereader.ReadFile("./input/day-10/test.txt")

	signalStrength := readInstructions(instructions, 20, 40)

	if signalStrength != 13140 {
		fmt.Println("Test went wrong:", signalStrength)
	} else {
		fmt.Println("Test passed:", signalStrength)
	}

	instructions = filereader.ReadFile("./input/day-10/input.txt")
	signalStrength = readInstructions(instructions, 20, 40)

	fmt.Println("Output:", signalStrength)
}

func readInstructions(instructions []string, firstCycle int, cycleIncrement int) int {
	running := true
	cycle := 0
	currentCheck := 0
	maxChecks := 220 - firstCycle/cycleIncrement
	instructionPointer := 0
	measurementPoints := make([]int, maxChecks+1)
	x := 1
	adding := false

	for running {
		cycle++
		if cycle == firstCycle+(cycleIncrement*currentCheck) {
			measurementPoints[currentCheck] = cycle * x
			currentCheck++
		}

		instruction := instructions[instructionPointer]

		if instruction == "noop" {
			instructionPointer++
		} else {
			var operation string
			var amount int
			fmt.Sscanln(instruction, &operation, &amount)
			if !adding && operation == "addx" {
				adding = true
			} else {
				x += amount
				adding = false
				instructionPointer++
			}
		}

		if maxChecks < currentCheck || instructionPointer == len(instructions) {
			running = false
		}
	}

	signalStrength := 0
	for _, measurement := range measurementPoints {
		signalStrength += measurement
	}

	return signalStrength
}

func Day10Part2() {
	fmt.Println("============== Part2 =============")
	instructions := filereader.ReadFile("./input/day-10/test.txt")
	readImageInstructions(instructions, 40)

	fmt.Println("============ SOLUTION ============")

	instructions = filereader.ReadFile("./input/day-10/input.txt")
	readImageInstructions(instructions, 40)
}

func readImageInstructions(instructions []string, cycleIncrement int) {
	running := true
	cycle := 0
	currentCheck := 0
	maxChecks := 240 / cycleIncrement
	instructionPointer := 0
	image := make([]string, maxChecks)

	x := 1
	adding := false
	currentCheckLine := ""
	currentLinePosition := 0
	for running {
		cycle++
		if currentLinePosition == x-1 || currentLinePosition == x || currentLinePosition == x+1 {
			currentCheckLine += "#"
			currentLinePosition++

		} else {
			currentCheckLine += "."
			currentLinePosition++
		}

		if cycle == cycleIncrement+cycleIncrement*currentCheck {
			image[currentCheck] = currentCheckLine
			currentLinePosition = 0
			currentCheckLine = ""
			currentCheck++
		}

		instruction := instructions[instructionPointer]

		if instruction == "noop" {
			instructionPointer++
		} else {
			var operation string
			var amount int
			fmt.Sscanln(instruction, &operation, &amount)
			if !adding && operation == "addx" {
				adding = true
			} else {
				x += amount
				adding = false
				instructionPointer++
			}
		}

		if maxChecks < currentCheck || instructionPointer == len(instructions) {
			running = false
		}
	}

	for _, line := range image {
		fmt.Println(line)
	}
}
