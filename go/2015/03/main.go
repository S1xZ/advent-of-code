package main

import (
	"fmt"
	"os"
	"strings"
)

func Sovle1(instruction string) int {
	res := 1
	sz := len(instruction)
	arr := make([][]bool, 2*sz)
	for i := range arr {
		arr[i] = make([]bool, 2*sz)
	}

	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = false
		}
	}

	pos_x := sz
	pos_y := sz
	arr[pos_x][pos_y] = true
	for i := 0; i < sz; i++ {

		if instruction[i] == "^"[0] {
			pos_y += 1
		}
		if instruction[i] == "v"[0] {
			pos_y -= 1
		}
		if instruction[i] == ">"[0] {
			pos_x += 1
		}
		if instruction[i] == "<"[0] {
			pos_x -= 1
		}

		if !arr[pos_x][pos_y] {
			arr[pos_x][pos_y] = true
			res += 1
		}
	}
	return res
}

func Sovle2(instruction string) int {
	res := 1
	sz := len(instruction)
	arr := make([][]bool, 2*sz)
	for i := range arr {
		arr[i] = make([]bool, 2*sz)
	}

	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = false
		}
	}

	pos_x := sz
	pos_y := sz
	pos_x_rb := sz
	pos_y_rb := sz

	arr[pos_x][pos_y] = true
	for i := 0; i < sz; i++ {
		if i%2 == 0 {
			if instruction[i] == "^"[0] {
				pos_y += 1
			}
			if instruction[i] == "v"[0] {
				pos_y -= 1
			}
			if instruction[i] == ">"[0] {
				pos_x += 1
			}
			if instruction[i] == "<"[0] {
				pos_x -= 1
			}
		} else {
			if instruction[i] == "^"[0] {
				pos_y_rb += 1
			}
			if instruction[i] == "v"[0] {
				pos_y_rb -= 1
			}
			if instruction[i] == ">"[0] {
				pos_x_rb += 1
			}
			if instruction[i] == "<"[0] {
				pos_x_rb -= 1
			}
		}
		if !arr[pos_x][pos_y] || !arr[pos_x_rb][pos_y_rb] {
			arr[pos_x][pos_y] = true
			arr[pos_x_rb][pos_y_rb] = true
			res += 1
		}
	}
	return res
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	strdat := strings.TrimRight(string(dat), "\n")
	fmt.Println(Sovle1(strdat))
	fmt.Println(Sovle2(strdat))
}
