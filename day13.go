package main

import (
	"fmt"
	"internal/filereader"
	"sort"
	"strconv"
)

type DataPacketPair struct {
	left  []int
	right []int
}

func Day13() {
	fmt.Println("============== Day13 ==============")
	Day13Part1()
	Day13Part2()
}

func Day13Part1() {
	fmt.Println("============== Part1 =============")
	packetData := filereader.ReadFile("./input/day-13/test.txt")
	pairs := createPairs(packetData)
	validPairs := sumValidPairs(pairs)

	if validPairs != 13 {
		fmt.Println("Test went wrong:", validPairs)
	} else {
		fmt.Println("Test passed:", validPairs)
	}

	packetData = filereader.ReadFile("./input/day-13/input.txt")
	pairs = createPairs(packetData)
	validPairs = sumValidPairs(pairs)

	fmt.Println("Output:", validPairs)
}

func sumValidPairs(pairs [][]any) int {
	valid := 0
	for i, pair := range pairs {
		if sub := rightOrder(pair[0], pair[1]); sub == 0 || sub == 1 {
			valid += i + 1
		}
	}
	return valid
}

func rightOrder(left any, right any) int {
	leftInt, leftIsInt := left.(int)
	rightInt, rightIsInt := right.(int)
	if leftIsInt && rightIsInt {
		if leftInt > rightInt {
			return -1
		} else if rightInt > leftInt {
			return 1
		}
		return 0
	}

	leftList, leftIsList := left.([]any)
	rightList, rightIsList := right.([]any)
	if !leftIsList {
		leftList = []any{leftInt}
	}
	if !rightIsList {
		rightList = []any{rightInt}
	}
	max := maxFromIntSlice([]int{len(leftList), len(rightList)})
	for i := 0; i < max; i++ {
		if i >= len(leftList) {
			return 1
		}
		if i >= len(rightList) {
			return -1
		}
		if sub := rightOrder(leftList[i], rightList[i]); sub != 0 {
			return sub
		}
	}
	return 0
}

func maxFromIntSlice(numbers []int) int {
	max := 0
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

func createPairs(packetData []string) [][]any {
	pairs := make([][]any, 0)
	pair := make([]any, 0, 2)
	for _, data := range packetData {
		if data == "" {
			continue
		}
		packet, _ := parsePacket(data)
		pair = append(pair, packet)
		if len(pair) == 2 {
			pairs = append(pairs, pair)
			pair = make([]any, 0, 2)
		}
	}
	return pairs
}

func parsePacket(packet string) (any, int) {
	chars := []rune(packet)
	id := 0
	output := make([]any, 0)
	numberCharacters := make([]rune, 0)
	for id < len(chars) {
		char := chars[id]
		switch char {
		case '[':
			x, i := parsePacket(string(chars[id+1:]))
			output = append(output, x)
			id += i + 1
		case ']':
			if len(numberCharacters) > 0 {
				n, _ := strconv.Atoi(string(numberCharacters))
				output = append(output, n)
				numberCharacters = make([]rune, 0)
			}
			id++
			return output, id
		case ',':
			if len(numberCharacters) > 0 {
				n, _ := strconv.Atoi(string(numberCharacters))
				output = append(output, n)
				numberCharacters = make([]rune, 0)
			}
			id++
		default:
			numberCharacters = append(numberCharacters, char)
			id++
		}
	}
	return output, id
}

func Day13Part2() {
	fmt.Println("============== Part2 =============")
	packetData := filereader.ReadFile("./input/day-13/test.txt")
	dividers := []string{"[[2]]", "[[6]]"}
	packetData = append(packetData, dividers...)

	pairs := createPairs(packetData)
	key := findDecoderKey(pairs)

	if key != 140 {
		fmt.Println("Test went wrong:", key)
	} else {
		fmt.Println("Test passed:", key)
	}

	packetData = filereader.ReadFile("./input/day-13/input.txt")
	packetData = append(packetData, dividers...)
	pairs = createPairs(packetData)
	key = findDecoderKey(pairs)

	fmt.Println("Output:", key)
}

func findDecoderKey(pairs [][]any) int {
	allPackets := make([]any, 0, len(pairs)*2)
	for _, pair := range pairs {
		allPackets = append(allPackets, pair...)
	}
	sort.Slice(allPackets, func(i int, j int) bool {
		left := allPackets[i]
		right := allPackets[j]
		return rightOrder(left, right) == 1
	})
	out := 1
	for i, packet := range allPackets {
		asString := fmt.Sprintf("%v", packet)
		if asString == "[[[2]]]" || asString == "[[[6]]]" {
			fmt.Println("found dividers")
			out *= i + 1
		}
	}
	return out
}
