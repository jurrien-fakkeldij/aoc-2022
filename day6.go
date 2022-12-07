package main

import (
	"fmt"
	"internal/filereader"
	"strings"
)

func Day6() {
	fmt.Println("============== Day6 ==============")
	Day6Part1()
	Day6Part2()
}

func Day6Part1() {
	fmt.Println("============== Part1 =============")
	instructions := filereader.ReadFile("./input/day-6/test.txt")

	marker, position := findMarker(instructions[0])

	if marker != "jpqm" && position == 7 {
		fmt.Println("Test went wrong:", marker, position)
	} else {
		fmt.Println("Test passed:", marker, position)
	}

	instructions = filereader.ReadFile("./input/day-6/input.txt")
	marker, position = findMarker(instructions[0])

	fmt.Println("Output:", marker, position)
}

func findMarker(instructions string) (string, int) {
	startPosition := 2
	markerSet := instructions[:startPosition+1]
	markerPosition := 0
	noDuplicateChars := false
	for i := startPosition + 1; i < len(instructions); i++ {
		markerSet += string(instructions[i])
		for j := 0; j < len(markerSet); j++ {
			newMarkerSet := removeDuplicateValues(markerSet)
			if strings.Compare(newMarkerSet, markerSet) != 0 {
				//fmt.Println(markerSet, i, string(character))
				markerSet = markerSet[1:]
				//fmt.Println(markerSet, i)
				break
			} else {
				noDuplicateChars = true
			}
		}

		if noDuplicateChars {
			markerPosition = i + 1
			break
		}
	}
	return markerSet, markerPosition
}

func removeDuplicateValues(line string) string {
	keys := make(map[string]bool)
	list := ""

	for _, entry := range line {
		if _, value := keys[string(entry)]; !value {
			keys[string(entry)] = true
			list += string(entry)
		}
	}
	return list
}

func Day6Part2() {
	fmt.Println("============== Part2 =============")
	instructions := filereader.ReadFile("./input/day-6/test.txt")

	message, position := findMessage(instructions[0])

	if position != 19 {
		fmt.Println("Test went wrong:", message, position)
	} else {
		fmt.Println("Test passed:", message, position)
	}

	instructions = filereader.ReadFile("./input/day-6/input.txt")
	message, position = findMessage(instructions[0])

	fmt.Println("Output:", message, position)
}

func findMessage(instructions string) (string, int) {
	startPosition := 12
	markerSet := instructions[:startPosition+1]
	markerPosition := 0
	noDuplicateChars := false
	for i := startPosition + 1; i < len(instructions); i++ {
		markerSet += string(instructions[i])
		for j := 0; j < len(markerSet); j++ {
			newMarkerSet := removeDuplicateValues(markerSet)
			if strings.Compare(newMarkerSet, markerSet) != 0 {
				//fmt.Println(markerSet, i, string(character))
				markerSet = markerSet[1:]
				//fmt.Println(markerSet, i)
				break
			} else {
				noDuplicateChars = true
			}
		}

		if noDuplicateChars {
			markerPosition = i + 1
			break
		}
	}
	return markerSet, markerPosition
}
