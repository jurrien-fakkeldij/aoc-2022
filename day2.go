package main

import (
	"fmt"
	"internal/filereader"
	"strings"
)

func Day2() {
	fmt.Println("============== Day2 ==============")
	Day2Part1()
	Day2Part2()
}

func Day2Part1() {
	fmt.Println("============== Part1 =============")
	playList := filereader.ReadFile("./input/day-2/test.txt")

	score := playFromList(playList)

	if score != 15 {
		fmt.Println("Test went wrong:", score)
	} else {
		fmt.Println("Test passed:", score)
	}

	playList = filereader.ReadFile("./input/day-2/input.txt")
	score = playFromList(playList)
	fmt.Println("Output:", score)
}

func playFromList(playList []string) int {
	score := 0
	playBook := make(map[string]int)
	noPlays := 0
	for _, play := range playList {
		if val, ok := playBook[play]; ok {
			score += val
		} else {
			currentPlays := strings.Split(play, " ")
			fmt.Println(currentPlays)
			intermediateScore := determineScore(currentPlays[0], currentPlays[1])
			playBook[play] = intermediateScore
			score += intermediateScore
		}
		noPlays++
	}
	fmt.Println("plays: ", noPlays)
	return score
}

func determineScore(first string, second string) int {
	if first == "A" { //rock
		switch second {
		case "X": //rock
			return 4
		case "Y": //paper
			return 8
		case "Z": //scissor
			return 3
		}
	} else if first == "B" { // paper
		switch second {
		case "X":
			return 1
		case "Y":
			return 5
		case "Z":
			return 9
		}
	} else if first == "C" { //scissors
		switch second {
		case "X":
			return 7
		case "Y":
			return 2
		case "Z":
			return 6
		}
	}
	fmt.Println("No score something went wrong")

	return 0
}

func Day2Part2() {
	fmt.Println("============== Part2 =============")
	playList := filereader.ReadFile("./input/day-2/test.txt")

	score := playFromSecretList(playList)

	if score != 12 {
		fmt.Println("Test went wrong:", score)
	} else {
		fmt.Println("Test passed:", score)
	}

	playList = filereader.ReadFile("./input/day-2/input.txt")
	score = playFromSecretList(playList)
	fmt.Println("Output:", score)
}

func playFromSecretList(playList []string) int {
	score := 0
	playBook := make(map[string]int)
	noPlays := 0
	for _, play := range playList {
		if val, ok := playBook[play]; ok {
			score += val
		} else {
			currentPlays := strings.Split(play, " ")
			intermediateScore := determineSecretScore(currentPlays[0], currentPlays[1])
			playBook[play] = intermediateScore
			score += intermediateScore
		}
		noPlays++
	}
	fmt.Println("plays: ", noPlays)
	return score
}

func determineSecretScore(first string, second string) int {
	if first == "A" { //rock
		switch second {
		case "X": //lose
			return 3
		case "Y": //draw
			return 4
		case "Z": //win
			return 8
		}
	} else if first == "B" { // paper
		switch second {
		case "X": //lose
			return 1
		case "Y": //draw
			return 5
		case "Z": //win
			return 9
		}
	} else if first == "C" { //scissors
		switch second {
		case "X": //lose
			return 2
		case "Y": //draw
			return 6
		case "Z": //win
			return 7
		}
	}
	fmt.Println("No score something went wrong")

	return 0
}
