package main

import (
	"fmt"
	"internal/filereader"
	"math"
	"strconv"
	"strings"
)

type Space struct {
	x         int
	y         int
	spaceType string
}

func (s *Space) toString() string {
	return fmt.Sprintf("%d,%d", s.x, s.y)
}

func Day14() {
	fmt.Println("============== Day14 ==============")
	//Day14Part1()
	Day14Part2()
}

func Day14Part1() {
	fmt.Println("============== Part1 =============")
	caveData := filereader.ReadFile("./input/day-14/test.txt")
	cave, minX, maxX, minY, maxY := setupCave(caveData)
	sandUnits := dropSand(cave, minX, maxX, minY, maxY)

	if sandUnits != 24 {
		fmt.Println("Test went wrong:", sandUnits)
	} else {
		fmt.Println("Test passed:", sandUnits)
	}

	caveData = filereader.ReadFile("./input/day-14/input.txt")
	cave, minX, maxX, minY, maxY = setupCave(caveData)

	sandUnits = dropSand(cave, minX, maxX, minY, maxY)

	fmt.Println("Output:", sandUnits)
}
func dropSand(cave map[string]*Space, minX int, maxX int, minY int, maxY int) int {
	sandDroppingEternally := false
	sandCount := 0
	for !sandDroppingEternally {
		space := &Space{x: 500, y: 0, spaceType: "o"}
		sandDropping := true
		for sandDropping {
			bottomSpace := &Space{x: space.x, y: space.y + 1, spaceType: "o"}
			if _, ok := cave[bottomSpace.toString()]; !ok {
				space = bottomSpace
				if space.y > maxY {
					fmt.Println("space bottomless:", space)
					sandDroppingEternally = true
					sandDropping = false
				}
			} else {
				leftSpace := &Space{x: space.x - 1, y: space.y + 1, spaceType: "o"}
				rightSpace := &Space{x: space.x + 1, y: space.y + 1, spaceType: "o"}
				if _, ok := cave[leftSpace.toString()]; !ok {
					space = leftSpace
				} else if _, ok := cave[rightSpace.toString()]; !ok {
					space = rightSpace
				} else {
					cave[space.toString()] = space
					sandDropping = false
					sandCount++
				}
			}
		}
		showCave(cave, minX, maxX, minY, maxY)
	}

	return sandCount
}

func setupCave(caveInput []string) (map[string]*Space, int, int, int, int) {
	cave := make(map[string]*Space)
	maxY := 0
	maxX := 0
	minY := 0
	minX := 500

	for _, formation := range caveInput {
		formationPoints := strings.Split(formation, " -> ")
		for idx := 0; idx < len(formationPoints)-1; idx++ {
			startPoint := strings.Split(formationPoints[idx], ",")
			endPoint := strings.Split(formationPoints[idx+1], ",")

			startingX, _ := strconv.Atoi(startPoint[0])
			startingY, _ := strconv.Atoi(startPoint[1])
			endingX, _ := strconv.Atoi(endPoint[0])
			endingY, _ := strconv.Atoi(endPoint[1])

			if startingX < minX {
				minX = startingX
			}

			if endingX < minX {
				minX = endingX
			}

			if endingX > maxX {
				maxX = endingX
			}

			if startingX > maxX {
				maxX = startingX
			}

			if endingY > maxY {
				maxY = endingY
			}

			if startingY > maxY {
				maxY = startingY
			}

			if startingX == endingX {
				beginY := int(math.Min(float64(startingY), float64(endingY)))
				endY := int(math.Max(float64(startingY), float64(endingY)))
				for y := beginY; y <= endY; y++ {
					space := &Space{x: startingX, y: y, spaceType: "#"}
					if _, ok := cave[space.toString()]; !ok {
						cave[space.toString()] = space
					}
				}
			} else if startingY == endingY {
				beginX := int(math.Min(float64(startingX), float64(endingX)))
				endX := int(math.Max(float64(startingX), float64(endingX)))
				for x := beginX; x <= endX; x++ {
					space := &Space{x: x, y: endingY, spaceType: "#"}
					if _, ok := cave[space.toString()]; !ok {
						cave[space.toString()] = space
					}
				}
			}

		}
	}

	return cave, minX, maxX, minY, maxY
}

func showCave(cave map[string]*Space, minX int, maxX int, minY int, maxY int) {
	for y := minY; y <= maxY; y++ {

		row := ""
		for x := minX; x <= maxX; x++ {
			location := fmt.Sprintf("%d,%d", x, y)

			if loc, ok := cave[location]; ok {
				row += loc.spaceType
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
}

func Day14Part2() {
	fmt.Println("============== Part2 =============")
	caveData := filereader.ReadFile("./input/day-14/test.txt")
	cave, minX, maxX, minY, maxY := setupCave(caveData)
	sandUnits := findSourceBlockedSand(cave, minX, maxX, minY, maxY)

	if sandUnits != 93 {
		fmt.Println("Test went wrong:", sandUnits)
	} else {
		fmt.Println("Test passed:", sandUnits)
	}

	caveData = filereader.ReadFile("./input/day-14/input.txt")
	cave, minX, maxX, minY, maxY = setupCave(caveData)

	sandUnits = findSourceBlockedSand(cave, minX, maxX, minY, maxY)

	fmt.Println("Output:", sandUnits)
}

func findSourceBlockedSand(cave map[string]*Space, minX int, maxX int, minY int, maxY int) int {
	foundBlockage := false
	sandCount := 0
	for !foundBlockage {
		space := &Space{x: 500, y: 0, spaceType: "o"}
		if _, ok := cave[space.toString()]; ok {
			foundBlockage = true
			break
		}

		sandDropping := true
		for sandDropping {
			bottomSpace := &Space{x: space.x, y: space.y + 1, spaceType: "o"}
			if bottomSpace.y == maxY+2 {
				cave[space.toString()] = space
				sandDropping = false
				sandCount++
			}
			if _, ok := cave[bottomSpace.toString()]; !ok {
				space = bottomSpace
			} else {
				leftSpace := &Space{x: space.x - 1, y: space.y + 1, spaceType: "o"}
				rightSpace := &Space{x: space.x + 1, y: space.y + 1, spaceType: "o"}
				if _, ok := cave[leftSpace.toString()]; !ok {
					space = leftSpace
				} else if _, ok := cave[rightSpace.toString()]; !ok {
					space = rightSpace
				} else {
					cave[space.toString()] = space
					sandDropping = false
					sandCount++
				}
			}
		}
		showCave(cave, minX, maxX, minY, maxY)
	}

	return sandCount
}
