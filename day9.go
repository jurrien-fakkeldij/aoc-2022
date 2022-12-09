package main

import (
	"fmt"
	"internal/filereader"
	"math"
	"strconv"
)

func Day9() {
	fmt.Println("============== Day9 ==============")
	Day9Part1()
	Day9Part2()
}

type Position struct {
	x int
	y int
}

type Vector struct {
	x int
	y int
}

func (p *Position) toString() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func Day9Part1() {
	fmt.Println("============== Part1 =============")
	instructions := filereader.ReadFile("./input/day-9/test.txt")

	visited := uniquePositionsVisited(instructions, false)

	if visited != 13 {
		fmt.Println("Test went wrong:", visited)
	} else {
		fmt.Println("Test passed:", visited)
	}

	instructions = filereader.ReadFile("./input/day-9/input.txt")
	visited = uniquePositionsVisited(instructions, false)

	fmt.Println("Output:", visited)
}

func uniquePositionsVisited(instructions []string, debug bool) int {
	uniqueLocationsVisited := make(map[string]*Position)
	startingPosition := &Position{x: 0, y: 0}
	uniqueLocationsVisited[startingPosition.toString()] = startingPosition
	currentPosition := startingPosition
	previousPosition := currentPosition

	for _, instruction := range instructions {
		dX := 0
		dY := 0
		switch instruction[0] {
		case 'U':
			dY = 1
		case 'D':
			dY = -1
		case 'L':
			dX = -1
		case 'R':
			dX = 1
		}

		movement, _ := strconv.Atoi(string(instruction[2:]))
		if debug {
			fmt.Println(movement)
		}

		for i := 1; i <= movement; i++ {
			newPosition := &Position{x: currentPosition.x + dX, y: currentPosition.y + dY}
			if debug {
				fmt.Println("new:", newPosition.toString())
			}
			if newPosition.x == previousPosition.x && newPosition.y == previousPosition.y {
				currentPosition = newPosition
				continue
			}
			if debug {
				fmt.Println("curr:", currentPosition.toString())
				fmt.Println("prev:", previousPosition.toString())
				//fmt.Println(math.Abs(float64(previousPosition.x)-float64(newPosition.x)), math.Abs(float64(previousPosition.y)-float64(newPosition.y)))
			}

			if math.Abs(float64(previousPosition.x)-float64(newPosition.x)) == 1 && math.Abs(float64(previousPosition.y)-float64(newPosition.y)) == 1 {
				currentPosition = newPosition
				continue
			}

			if math.Abs(float64(previousPosition.x)-float64(newPosition.x)) == 0 && math.Abs(float64(previousPosition.y)-float64(newPosition.y)) == 1 {
				currentPosition = newPosition
				continue
			}

			if math.Abs(float64(previousPosition.x)-float64(newPosition.x)) == 1 && math.Abs(float64(previousPosition.y)-float64(newPosition.y)) == 0 {
				currentPosition = newPosition
				continue
			}

			previousPosition = currentPosition
			currentPosition = newPosition
			if debug {
				fmt.Println("prev:", previousPosition.toString())
			}
			if _, ok := uniqueLocationsVisited[previousPosition.toString()]; !ok {
				uniqueLocationsVisited[previousPosition.toString()] = previousPosition
			}
		}
	}
	return len(uniqueLocationsVisited)
}

func Day9Part2() {
	fmt.Println("============== Part2 =============")
	instructions := filereader.ReadFile("./input/day-9/test2.txt")

	visited := uniquePositionsVisitedForMore(instructions, false)

	if visited != 36 {
		fmt.Println("Test went wrong:", visited)
	} else {
		fmt.Println("Test passed:", visited)
	}

	instructions = filereader.ReadFile("./input/day-9/input.txt")
	visited = uniquePositionsVisitedForMore(instructions, false)

	fmt.Println("Output:", visited)
}

func uniquePositionsVisitedForMore(instructions []string, debug bool) int {
	uniqueLocationsVisited := make(map[string]*Position)

	snakePositions := []*Position{&Position{x: 0, y: 0}, &Position{x: 0, y: 0}, &Position{x: 0, y: 0}, &Position{x: 0, y: 0}, &Position{x: 0, y: 0}, &Position{x: 0, y: 0}, &Position{x: 0, y: 0}, &Position{x: 0, y: 0}, &Position{x: 0, y: 0}, &Position{x: 0, y: 0}}

	for index, instruction := range instructions {
		dX := 0
		dY := 0
		switch instruction[0] {
		case 'U':
			dY = 1
		case 'D':
			dY = -1
		case 'L':
			dX = -1
		case 'R':
			dX = 1
		}
		movement, _ := strconv.Atoi(string(instruction[2:]))

		for i := 1; i <= movement; i++ {
			snakePositions[0].x = snakePositions[0].x + dX
			snakePositions[0].y = snakePositions[0].y + dY

			if index < 3 {
				debug = true
			} else {
				debug = false
			}
			snakePositions = determineNewPositions(snakePositions, debug)
			uniqueLocationsVisited[snakePositions[9].toString()] = snakePositions[9]
		}
	}
	return len(uniqueLocationsVisited)
}

func determineNewPositions(snake []*Position, debug bool) []*Position {

	for index, currentPosition := range snake {
		if index == 0 {
			continue
		}

		newPosition := snake[index-1]
		diffInX := math.Abs(float64(currentPosition.x - newPosition.x))
		diffInY := math.Abs(float64(currentPosition.y - newPosition.y))
		if diffInY <= 1 && diffInX <= 1 {
			continue
		}

		xDirection := 1
		yDirection := 1

		if newPosition.y < currentPosition.y {
			yDirection = -1
		}
		if newPosition.x < currentPosition.x {
			xDirection = -1
		}
		if newPosition.x == currentPosition.x {
			currentPosition.y += 1 * yDirection
			continue
		}
		if newPosition.y == currentPosition.y {
			currentPosition.x += 1 * xDirection
			continue
		}
		currentPosition.x += 1 * xDirection
		currentPosition.y += 1 * yDirection
	}
	if debug {
		for index, pos := range snake {
			fmt.Printf("%d %v\n", index, pos)
		}
	}
	return snake
}
