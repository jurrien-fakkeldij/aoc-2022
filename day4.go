package main

import (
	"fmt"
	"internal/filereader"
	"strconv"
	"strings"
)

func Day4() {
	fmt.Println("============== Day4 ==============")
	Day4Part1()
	Day4Part2()
}

func Day4Part1() {
	fmt.Println("============== Part1 =============")
	assignments := filereader.ReadFile("./input/day-4/test.txt")

	noOfAssignments := findContainingAssigments(assignments)

	if noOfAssignments != 2 {
		fmt.Println("Test went wrong:", noOfAssignments)
	} else {
		fmt.Println("Test passed:", noOfAssignments)
	}

	assignments = filereader.ReadFile("./input/day-4/input.txt")
	noOfAssignments = findContainingAssigments(assignments)

	fmt.Println("Output:", noOfAssignments)
}

func findContainingAssigments(assignments []string) int {
	countContainingAssignments := 0
	for _, assignment := range assignments {
		pairs := strings.Split(assignment, ",")
		firstPairFirstAssi, _ := strconv.Atoi(strings.Split(pairs[0], "-")[0])
		firstPairSecondAssi, _ := strconv.Atoi(strings.Split(pairs[0], "-")[1])
		secondPairFirstAssi, _ := strconv.Atoi(strings.Split(pairs[1], "-")[0])
		secondPairSecondAssi, _ := strconv.Atoi(strings.Split(pairs[1], "-")[1])

		if firstPairFirstAssi <= secondPairFirstAssi &&
			firstPairSecondAssi >= secondPairSecondAssi {
			countContainingAssignments++
		} else if firstPairFirstAssi >= secondPairFirstAssi &&
			firstPairSecondAssi <= secondPairSecondAssi {
			countContainingAssignments++
		}
	}

	return countContainingAssignments
}

func Day4Part2() {
	fmt.Println("============== Part2 =============")
	assignments := filereader.ReadFile("./input/day-4/test.txt")

	noOfAssignments := findOverlappingAssigments(assignments)

	if noOfAssignments != 4 {
		fmt.Println("Test went wrong:", noOfAssignments)
	} else {
		fmt.Println("Test passed:", noOfAssignments)
	}

	assignments = filereader.ReadFile("./input/day-4/input.txt")
	noOfAssignments = findOverlappingAssigments(assignments)

	fmt.Println("Output:", noOfAssignments)
}

func findOverlappingAssigments(assignments []string) int {

	countOverlappingAssignments := 0
	for _, assignment := range assignments {
		pairs := strings.Split(assignment, ",")
		firstPairFirstAssi, _ := strconv.Atoi(strings.Split(pairs[0], "-")[0])
		firstPairSecondAssi, _ := strconv.Atoi(strings.Split(pairs[0], "-")[1])
		secondPairFirstAssi, _ := strconv.Atoi(strings.Split(pairs[1], "-")[0])
		secondPairSecondAssi, _ := strconv.Atoi(strings.Split(pairs[1], "-")[1])

		if firstPairFirstAssi <= secondPairFirstAssi &&
			firstPairSecondAssi >= secondPairSecondAssi {
			countOverlappingAssignments++
		} else if firstPairFirstAssi >= secondPairFirstAssi &&
			firstPairSecondAssi <= secondPairSecondAssi {
			countOverlappingAssignments++
		} else if firstPairFirstAssi <= secondPairFirstAssi && // 1-10, 2-12
			firstPairSecondAssi >= secondPairFirstAssi {
			countOverlappingAssignments++
		} else if secondPairSecondAssi <= firstPairSecondAssi && //6-10, 3-9
			secondPairSecondAssi >= firstPairFirstAssi {
			countOverlappingAssignments++
		} else if firstPairFirstAssi >= secondPairFirstAssi && //6-10, 1-7
			firstPairSecondAssi <= secondPairFirstAssi {
			countOverlappingAssignments++
		} else if secondPairSecondAssi >= firstPairSecondAssi &&
			secondPairSecondAssi <= firstPairFirstAssi {
			countOverlappingAssignments++
		}
	}
	return countOverlappingAssignments
}
