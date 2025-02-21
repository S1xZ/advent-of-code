package main

import (
	"fmt"
	"os"
	"strings"
)

func Solve1(instruction string) bool {
	isNice := false
	vowelCount := 0
	haveDoubleLetter := false
	haveForbidenCombination := false
	prevChar := ""
	for i := 0; i < len(instruction); i++ {
		if string(instruction[i]) == "a" ||
			string(instruction[i]) == "e" ||
			string(instruction[i]) == "i" ||
			string(instruction[i]) == "o" ||
			string(instruction[i]) == "u" {
			vowelCount += 1
		}
		if prevChar == string(instruction[i]) {
			haveDoubleLetter = true
		}
		if prevChar == "a" && string(instruction[i]) == "b" ||
			prevChar == "c" && string(instruction[i]) == "d" ||
			prevChar == "p" && string(instruction[i]) == "q" ||
			prevChar == "x" && string(instruction[i]) == "y" {
			haveForbidenCombination = true
		}
		prevChar = string(instruction[i])
	}
	if vowelCount >= 3 &&
		haveDoubleLetter &&
		!haveForbidenCombination {
		isNice = true
	}
	return isNice
}

func Solve2(instruction string) bool {
	isNice := false
	twoPrevChar := false
	havePairs := false
	strMap := map[string]int{}
	for i := 1; i < len(instruction); i++ {
		_, exists := strMap[instruction[i-1:i+1]]
		if !exists {
			strMap[instruction[i-1:i+1]] = 1
		} else if instruction[i-2:i] != instruction[i-1:i+1] {
			havePairs = true
		}
		if i >= 2 {
			if instruction[i-2] == instruction[i] {
				twoPrevChar = true
			}
		}
	}
	if havePairs && twoPrevChar {
		isNice = true
	}
	return isNice
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	strdat := strings.TrimRight(string(dat), "\n")
	strArr := strings.Split(strdat, "\n")
	res1 := 0
	for i := 0; i < len(strArr); i++ {
		if Solve1(strArr[i]) {
			res1 += 1
		}
	}
	res2 := 0
	for i := 0; i < len(strArr); i++ {
		if Solve2(strArr[i]) {
			res2 += 1
		}
	}
	fmt.Println(res1)
	fmt.Println(res2)
}
