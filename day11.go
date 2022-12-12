package main

import (
	"fmt"
	"internal/filereader"
	"internal/transformer"
	"math"
	"strconv"
	"strings"
)

type Monkey struct {
	items            []uint64
	operation        string
	opAmount         string
	checkWorryAmount int
	trueMonkey       int
	falseMonkey      int
}

var commonDivisor uint64 = 1

func Day11() {
	fmt.Println("============== Day11 ==============")
	Day11Part1()
	Day11Part2()
}

func Day11Part1() {
	fmt.Println("============== Part1 =============")
	monkeys := filereader.ReadFile("./input/day-11/test.txt")

	monkeyTable := setupMonkeys(monkeys)

	monkeyBusiness := playRounds(monkeyTable, 20, true)

	if monkeyBusiness != 10605 {
		fmt.Println("Test went wrong:", monkeyBusiness)
	} else {
		fmt.Println("Test passed:", monkeyBusiness)
	}

	monkeys = filereader.ReadFile("./input/day-11/input.txt")
	monkeyTable = setupMonkeys(monkeys)
	monkeyBusiness = playRounds(monkeyTable, 20, true)

	fmt.Println("Output:", monkeyBusiness)
}

func playRounds(monkeys map[int]*Monkey, rounds int, worryLevelDiminished bool) uint64 {
	currentRound := 1
	monkeysInspectionTable := make(map[int]int)

	for currentRound <= rounds {
		currentMonkeyIndex := 0
		for currentMonkeyIndex < len(monkeys) {
			if _, ok := monkeysInspectionTable[currentMonkeyIndex]; !ok {
				monkeysInspectionTable[currentMonkeyIndex] = 0
			}

			currentMonkey := monkeys[currentMonkeyIndex]
			if len(currentMonkey.items) > 0 {
				for _, item := range currentMonkey.items {
					currentWorry := item
					if currentMonkey.operation == "+" {
						if currentMonkey.opAmount == "old" {
							currentWorry += item
						} else {
							tempAmount, _ := strconv.Atoi(currentMonkey.opAmount)
							currentWorry += uint64(tempAmount)
						}
					} else if currentMonkey.operation == "*" {
						if currentMonkey.opAmount == "old" {
							currentWorry *= item
						} else {
							tempAmount, _ := strconv.Atoi(currentMonkey.opAmount)
							currentWorry *= uint64(tempAmount)
						}
					}

					if worryLevelDiminished {
						currentWorry = uint64(math.Floor(float64(currentWorry / 3)))
					} else {
						currentWorry = currentWorry % commonDivisor
					}

					if currentWorry%uint64(currentMonkey.checkWorryAmount) == 0 {
						monkeys[currentMonkey.trueMonkey].items = append(monkeys[currentMonkey.trueMonkey].items, currentWorry)
					} else {
						monkeys[currentMonkey.falseMonkey].items = append(monkeys[currentMonkey.falseMonkey].items, currentWorry)
					}

					monkeysInspectionTable[currentMonkeyIndex]++
				}
				currentMonkey.items = []uint64{}
			}
			currentMonkeyIndex++
		}
		currentRound++
	}

	highestMonkeyInspectionTimes := make([]int, 2)
	for _, monkeyInspectionTimes := range monkeysInspectionTable {
		if highestMonkeyInspectionTimes[0] < monkeyInspectionTimes {
			highestMonkeyInspectionTimes[1] = highestMonkeyInspectionTimes[0]
			highestMonkeyInspectionTimes[0] = monkeyInspectionTimes
		} else if highestMonkeyInspectionTimes[1] < monkeyInspectionTimes {
			highestMonkeyInspectionTimes[1] = monkeyInspectionTimes
		}
	}
	return uint64(highestMonkeyInspectionTimes[0]) * uint64(highestMonkeyInspectionTimes[1])
}

func setupMonkeys(monkeys []string) map[int]*Monkey {
	currentMonkeyIndex := 0
	currentMonkey := &Monkey{}

	monkeyTable := make(map[int]*Monkey)
	for id, line := range monkeys {
		if id == 0 {
			continue
		}

		if line == "" {
			monkeyTable[currentMonkeyIndex] = currentMonkey
			continue
		}

		if line[0:6] == "Monkey" {
			currentMonkeyIndex++
			currentMonkey = &Monkey{}
			continue
		}
		if line[0:2] == "  " {
			monkeyAttribute := strings.Split(line[2:], ":")
			if monkeyAttribute[0] == "Starting items" {
				worryItemLevels, _ := transformer.SliceAtoi(strings.Split(monkeyAttribute[1][1:], ", "))
				for _, item := range worryItemLevels {
					currentMonkey.items = append(currentMonkey.items, uint64(item))
				}
				continue
			}

			if monkeyAttribute[0] == "Operation" {
				currentMonkey.operation = string(monkeyAttribute[1][11])
				currentMonkey.opAmount = string(monkeyAttribute[1][13:])
				continue
			}

			if monkeyAttribute[0] == "Test" {
				worryAmount, _ := strconv.Atoi(monkeyAttribute[1][14:])
				currentMonkey.checkWorryAmount = worryAmount
				commonDivisor *= uint64(worryAmount)
				continue
			}

			if monkeyAttribute[0][0:2] == "  " {
				if monkeyAttribute[0] == "  If true" {

					monkeyToThrowTo, _ := strconv.Atoi(monkeyAttribute[1][17:])
					currentMonkey.trueMonkey = monkeyToThrowTo
					continue
				}

				if monkeyAttribute[0] == "  If false" {
					monkeyToThrowTo, _ := strconv.Atoi(monkeyAttribute[1][17:])
					currentMonkey.falseMonkey = monkeyToThrowTo
					continue
				}
			}
		}
	}
	monkeyTable[currentMonkeyIndex] = currentMonkey

	return monkeyTable
}

func Day11Part2() {
	monkeys := filereader.ReadFile("./input/day-11/test.txt")
	commonDivisor = 1
	monkeyTable := setupMonkeys(monkeys)

	monkeyBusiness := playRounds(monkeyTable, 10000, false)

	if monkeyBusiness != 2713310158 {
		fmt.Println("Test went wrong:", monkeyBusiness)
	} else {
		fmt.Println("Test passed:", monkeyBusiness)
	}

	commonDivisor = 1
	monkeys = filereader.ReadFile("./input/day-11/input.txt")
	monkeyTable = setupMonkeys(monkeys)
	monkeyBusiness = playRounds(monkeyTable, 10000, false)

	fmt.Println("Output:", monkeyBusiness)
}
