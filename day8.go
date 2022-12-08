package main

import (
	"fmt"
	"internal/filereader"
	"strconv"
)

type Tree struct {
	x       int
	y       int
	height  int
	visible bool
}

type Size struct {
	width  int
	height int
}

func Day8() {
	fmt.Println("============== Day8 ==============")
	Day8Part1()
	Day8Part2()
}

func Day8Part1() {
	fmt.Println("============== Part1 =============")
	treeGridMap := filereader.ReadFile("./input/day-8/test.txt")

	treeGrid, gridSize := setupTreeGrid(treeGridMap)

	visibleTrees := numberOfVisibleTrees(treeGrid, gridSize)

	if visibleTrees != 21 {
		fmt.Println("Test went wrong:", visibleTrees)
	} else {
		fmt.Println("Test passed:", visibleTrees)
	}

	treeGridMap = filereader.ReadFile("./input/day-8/input.txt")
	treeGrid, gridSize = setupTreeGrid(treeGridMap)

	visibleTrees = numberOfVisibleTrees(treeGrid, gridSize)

	fmt.Println("Output:", visibleTrees)
}

func setupTreeGrid(treeGridMap []string) (map[string]*Tree, *Size) {
	treeGrid := make(map[string]*Tree)
	height := 0
	width := 0
	for x, row := range treeGridMap {
		for y, col := range row {
			location := fmt.Sprint("%d,%d", x, y)
			height, _ := strconv.Atoi(string(col))
			treeGrid[location] = &Tree{height: height, visible: false, x: x, y: y}
			width = y + 1
		}
		height = x + 1
	}
	return treeGrid, &Size{height: height, width: width}
}

func numberOfVisibleTrees(grid map[string]*Tree, size *Size) int {
	visibleTrees := 0
	for _, tree := range grid {
		if tree.visible {
			continue
		}

		if isTreeHigherThanOneNeighbour(tree, grid, size) {
			tree.visible = true
			visibleTrees++
		}
	}
	return visibleTrees
}

func isTreeHigherThanOneNeighbour(tree *Tree, grid map[string]*Tree, size *Size) bool {
	if tree.x == size.width-1 || tree.x == 0 {
		return true
	}

	if tree.y == size.height-1 || tree.y == 0 {
		return true
	}

	higherUp := true
	higherDown := true
	higherLeft := true
	higherRight := true

	for x := 1; x < size.width-tree.x; x++ {
		if tree.height <= grid[fmt.Sprint("%d,%d", tree.x+x, tree.y)].height {
			higherRight = false
			break
		}
	}

	for x := tree.x - 1; x >= 0; x-- {
		if tree.height <= grid[fmt.Sprint("%d,%d", x, tree.y)].height {
			higherLeft = false
			break
		}
	}

	for y := 1; y < size.height-tree.y; y++ {
		if tree.height <= grid[fmt.Sprint("%d,%d", tree.x, tree.y+y)].height {
			higherDown = false
			break
		}
	}

	for y := tree.y - 1; y >= 0; y-- {
		if tree.height <= grid[fmt.Sprint("%d,%d", tree.x, y)].height {
			higherUp = false
			break
		}
	}

	return higherUp || higherDown || higherLeft || higherRight
}

func Day8Part2() {
	fmt.Println("============== Part2 =============")
	treeGridMap := filereader.ReadFile("./input/day-8/test.txt")

	treeGrid, gridSize := setupTreeGrid(treeGridMap)

	scenicScore := highestScenicValue(treeGrid, gridSize)

	if scenicScore != 8 {
		fmt.Println("Test went wrong:", scenicScore)
	} else {
		fmt.Println("Test passed:", scenicScore)
	}

	treeGridMap = filereader.ReadFile("./input/day-8/input.txt")
	treeGrid, gridSize = setupTreeGrid(treeGridMap)

	scenicScore = highestScenicValue(treeGrid, gridSize)

	fmt.Println("Output:", scenicScore)
}

func highestScenicValue(grid map[string]*Tree, size *Size) int {
	score := -1
	for _, tree := range grid {
		if score == -1 {
			score = calculateScenicScore(tree, grid, size)
		} else {
			currentScore := calculateScenicScore(tree, grid, size)
			if currentScore > score {
				score = currentScore
			}
		}
	}
	return score
}

func calculateScenicScore(tree *Tree, grid map[string]*Tree, size *Size) int {

	amountUp := 0
	amountDown := 0
	amountLeft := 0
	amountRight := 0
	debug := false

	for x := 1; x < size.width-tree.x; x++ {
		if tree.height <= grid[fmt.Sprint("%d,%d", tree.x+x, tree.y)].height {
			if debug {
				fmt.Println("heigher tree found at:", tree.x+x, tree.y)
			}
			amountRight++
			break
		} else {
			amountRight++
		}
	}
	if debug {
		fmt.Println("amount right:", amountRight)
	}

	for x := tree.x - 1; x >= 0; x-- {
		if tree.height <= grid[fmt.Sprint("%d,%d", x, tree.y)].height {
			if debug {
				fmt.Println("heigher tree found at:", x, tree.y)
			}
			amountLeft++
			break
		} else {
			amountLeft++
		}
	}
	if debug {
		fmt.Println("amount left:", amountLeft)
	}

	for y := 1; y < size.height-tree.y; y++ {
		if tree.height <= grid[fmt.Sprint("%d,%d", tree.x, tree.y+y)].height {
			if debug {
				fmt.Println("heigher tree found at:", tree.x, tree.y+y)
			}
			amountDown++
			break
		} else {
			amountDown++
		}
	}
	if debug {
		fmt.Println("amount Down:", amountDown)
	}

	for y := tree.y - 1; y >= 0; y-- {
		if tree.height <= grid[fmt.Sprint("%d,%d", tree.x, y)].height {
			if debug {
				fmt.Println("heigher tree found at:", tree.x, y)
			}
			amountUp++
			break
		} else {
			amountUp++
		}
	}
	if debug {
		fmt.Println("amount Up:", amountUp)
	}

	return amountUp * amountDown * amountLeft * amountRight
}
