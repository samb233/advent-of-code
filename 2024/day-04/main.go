package main

import (
	"fmt"
	"os"
	// "strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("4.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	// res := FindXMAS(lines)
	res := FindMASInX(lines)
	fmt.Println(res)
}

// part 1
func FindXMAS(lines []string) int {
	linesRune := make([][]rune, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		lineRune := []rune(line)
		linesRune = append(linesRune, lineRune)
	}

	res := 0
	for y, line := range linesRune {
		for x, ch := range line {
			if ch == 'X' {
				res = res + FindXMASH(linesRune, x, y) + FindXMASV(linesRune, x, y) + FindXMASD(linesRune, x, y)
 			}
		}
	}

	return res
}

// part 2
func FindMASInX(lines []string) int {
	linesRune := make([][]rune, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		lineRune := []rune(line)
		linesRune = append(linesRune, lineRune)
	}

	res := 0
	for y, line := range linesRune {
		for x, ch := range line {
			if ch == 'A' {
				res = res + FindMASInXD(linesRune, x, y)
 			}
		}
	}

	return res
}

func FindXMASH(lines [][]rune, x, y int) int {
	cnt := 0
	if x >= 3 {
		if lines[y][x - 1] == 'M' &&
			lines[y][x - 2] == 'A' &&
			lines[y][x - 3] == 'S' {
			cnt ++
		}
	}

	if len(lines[y]) > x + 3 {
		if lines[y][x + 1] == 'M' &&
			lines[y][x + 2] == 'A' &&
			lines[y][x + 3] == 'S' {
			cnt ++
		}
	}

	return cnt
}

func FindXMASV(lines [][]rune, x, y int) int {
	cnt := 0
	if y >= 3 {
		if lines[y - 1][x] == 'M' &&
			lines[y - 2][x] == 'A' &&
			lines[y - 3][x] == 'S' {
			cnt ++
		}
	}

	if len(lines) > y + 3 {
		if lines[y + 1][x] == 'M' &&
			lines[y + 2][x] == 'A' &&
			lines[y + 3][x] == 'S' {
			cnt ++
		}
	}

	return cnt
}

func FindXMASD(lines [][]rune, x, y int) int {
	cnt := 0
	if y >= 3 && x >= 3 {
		if lines[y - 1][x - 1] == 'M' &&
			lines[y - 2][x - 2] == 'A' &&
			lines[y - 3][x - 3] == 'S' {
			cnt ++
		}
	}

	if len(lines) > y + 3 && len(lines[y]) > x + 3 {
		if lines[y + 1][x + 1] == 'M' &&
			lines[y + 2][x + 2] == 'A' &&
			lines[y + 3][x + 3] == 'S' {
			cnt ++
		}
	}

	if y >= 3 && len(lines[y]) > x + 3 {
		if lines[y - 1][x + 1] == 'M' &&
			lines[y - 2][x + 2] == 'A' &&
			lines[y - 3][x + 3] == 'S' {
			cnt ++
		}
	}

	if len(lines) > y + 3 && x >= 3 {
		if lines[y + 1][x - 1] == 'M' &&
			lines[y + 2][x - 2] == 'A' &&
			lines[y + 3][x - 3] == 'S' {
			cnt ++
		}
	}

	return cnt
}

func FindMASInXD(lines [][]rune, x, y int) int {
	cnt := 0
	if y < 1 || x < 1 || len(lines[y]) <= x + 1 || len(lines) <= y + 1 {
		return 0
	}

	Xmas1 := lines[y - 1][x - 1] == 'M' && lines[y + 1][x + 1] == 'S'
	Xmas2 := lines[y - 1][x - 1] == 'S' && lines[y + 1][x + 1] == 'M'
	Xmas3 := lines[y - 1][x + 1] == 'S' && lines[y + 1][x - 1] == 'M'
	Xmas4 := lines[y - 1][x + 1] == 'M' && lines[y + 1][x - 1] == 'S'

	if (Xmas1 && Xmas3) || (Xmas1 && Xmas4) || (Xmas2 && Xmas3) || (Xmas2 && Xmas4) {
		cnt ++
	}

	return cnt
}
