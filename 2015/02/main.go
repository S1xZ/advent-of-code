package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	lenght int
	width  int
	height int
}

func ParseInstructions(instruction string) Cube {
	res := strings.Split(instruction, "x")

	lenght, _ := strconv.Atoi(res[0])
	width, _ := strconv.Atoi(res[1])
	height, _ := strconv.Atoi(res[2])

	return Cube{lenght, width, height}
}

func CalculateSerfaceArea(cb Cube) int {
	return 2*cb.lenght*cb.width + 2*cb.width*cb.height + 2*cb.lenght*cb.height
}

func GetSlack(cb Cube) int {
	return min(cb.lenght*cb.width, cb.width*cb.height, cb.lenght*cb.height)
}

func GetListOfInstructions(instructions string) []string {
	res := strings.Split(instructions, "\n")
	return res
}

func GetRibbonLenght(cb Cube) int {
	return min(2*cb.lenght+2*cb.height, 2*cb.width+2*cb.height, 2*cb.lenght+2*cb.width)
}

func GetBowLenght(cb Cube) int {
	return cb.lenght * cb.width * cb.height
}

func Solve1(instructions string) int {
	ans := 0
	inst_arr := GetListOfInstructions(instructions)
	for i := 0; i < len(inst_arr); i++ {
		cb := ParseInstructions(inst_arr[i])
		ans += CalculateSerfaceArea(cb) + GetSlack(cb)
	}
	return ans
}

func Solve2(instructions string) int {
	ans := 0
	inst_arr := GetListOfInstructions(instructions)
	for i := 0; i < len(inst_arr); i++ {
		cb := ParseInstructions(inst_arr[i])
		ans += GetRibbonLenght(cb) + GetBowLenght(cb)
	}
	return ans
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	strdat := strings.TrimRight(string(dat), "\n")
	fmt.Println(Solve1(strdat))
	fmt.Println(Solve2(strdat))
}
