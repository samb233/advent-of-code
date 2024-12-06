package main

import (
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

type PositionD struct {
	X int
	Y int
	D int
}

func main() {
	input, err := os.ReadFile("6.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	linesRune := make([][]rune, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		lineRune := []rune(line)
		linesRune = append(linesRune, lineRune)
	}

	// res := GuardPosition(linesRune)
	res := GuardPositionPart2(linesRune)
	// PrintRune(lines)

	fmt.Println(res)
}

// part 1
func GuardPosition(m [][]rune) int {
	posMap := make(map[Position]bool)

	X := 0
	Y := 0

	for y, line := range m {
		for x, ch := range line {
			if ch == '^' {
				X = x
				Y = y
			}
		}
	}

	// 1: up; 2: right; 3: down; 4: left
	direct := 1
	lastPos := &Position{}
	for Y >= 0 && Y < len(m) && X >= 0 && X < len(m[Y]) {
		if m[Y][X] == '#' {
			direct++
			if direct > 4 {
				direct = 1
			}

			X = lastPos.X
			Y = lastPos.Y
		} else {
			lastPos.X = X
			lastPos.Y = Y
			posMap[Position{X, Y}] = true
		}

		switch direct {
		case 1:
			Y--
		case 2:
			X++
		case 3:
			Y++
		case 4:
			X--
		}

	}

	return len(posMap)
}

// part 2
func GuardPositionPart2(m [][]rune) int {
	startX := 0
	startY := 0
	for y, line := range m {
		for x, ch := range line {
			if ch == '^' {
				startX = x
				startY = y
			}
		}
	}

	res := 0
	for stoneY := 0; stoneY < len(m); stoneY++ {
		for stoneX := 0; stoneX < len(m[stoneY]); stoneX++ {
			if stoneX == startX && stoneY == startY {
				continue
			}

			X := startX
			Y := startY
			posMap := make(map[PositionD]bool)
			lastPos := &PositionD{}

			// 1: up; 2: right; 3: down; 4: left
			direct := 1

			for Y >= 0 && Y < len(m) && X >= 0 && X < len(m[Y]) {
				if m[Y][X] == '#' || (X == stoneX && Y == stoneY) {
					direct++
					if direct > 4 {
						direct = 1
					}

					X = lastPos.X
					Y = lastPos.Y
				} else {
					lastPos.X = X
					lastPos.Y = Y
				}

				// fmt.Println(X, ", ", Y, ", ", direct)
				if posMap[PositionD{X, Y, direct}] {
					res++
					break
				}

				posMap[PositionD{X, Y, direct}] = true

				switch direct {
				case 1:
					Y--
				case 2:
					X++
				case 3:
					Y++
				case 4:
					X--
				}

			}
		}
	}

	return res
}
