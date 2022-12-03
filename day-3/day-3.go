package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	// filename := "sample.txt"
	filename := "input.txt"

	txtlines := utils.ReadFile(filename)

	points := day3par2(txtlines)

	fmt.Printf("Final points are %d\n", points)
}

func day3(line []string) int {
	var score = 0
	for i := 0; i < len(line); i++ {
		str := line[i]
		mid := len(str) / 2
		chars := strings.Split(str, "")
		res := intersection(chars[0:mid], chars[mid:])
		score += valueMap[res[0]]
	}
	return score
}

func day3par2(line []string) int {
	var score = 0
	fmt.Printf("Lines count: %d\n", len(line)/3)

	nextLine := 0

	for i := 0; i < len(line)/3; i++ {
		line1 := nextLine
		line2 := line1 + 1
		line3 := line2 + 1
		nextLine = line3 + 1
		fmt.Printf("Lines are: %d, %d, %d \n", line1, line2, line3)

		elf1 := strings.Split(line[line1], "")
		elf2 := strings.Split(line[line2], "")
		elf3 := strings.Split(line[line3], "")

		elf1And2 := intersection(elf1, elf2)
		res := intersection(elf1And2, elf3)
		score += valueMap[res[0]]
	}
	return score
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

var (
	valueMap = map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6,
		"g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12,
		"m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18,
		"s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24,
		"y": 25, "z": 26,
		"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32,
		"G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38,
		"M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44,
		"S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50,
		"Y": 51, "Z": 52,
	}
)
