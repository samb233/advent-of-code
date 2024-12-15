package main

import (
	"fmt"
	"os"
	// "sort"
	// "strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("15.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n\n")

	sX := 0
	sY := 0

	m := make([][]rune, 0)
	mLines := strings.Split(lines[0], "\n")
	for y, line := range mLines {
		if len(line) == 0 {
			continue
		}

		mR := []rune(line)
		m = append(m, mR)

		for x, ch := range line {
			if ch == '@' {
				sX = x
				sY = y
			}
		}
	}

	m2 := make([][]rune, 0)
	for _, line := range mLines {
		if len(line) == 0 {
			continue
		}

		mr := make([]rune, 0)

		for _, ch := range line {
			if ch == '@' {
				mr = append(mr, '@')
				mr = append(mr, '.')
			} else if ch == '#' {
				mr = append(mr, '#')
				mr = append(mr, '#')
			} else if ch == '.' {
				mr = append(mr, '.')
				mr = append(mr, '.')
			} else if ch == 'O' {
				mr = append(mr, '[')
				mr = append(mr, ']')
			}
		}

		m2 = append(m2, mr)
	}

	for y, line := range m2 {
		for x, ch := range line {
			if ch == '@' {
				sX, sY = x, y
			}

		}
	}

	moves := make([][]int, 0)
	for _, ch := range lines[1] {
		m := make([]int, 0)
		switch ch {
		case '<':
			m = []int{-1, 0}
		case '^':
			m = []int{0, -1}
		case '>':
			m = []int{1, 0}
		case 'v':
			m = []int{0, 1}
		}

		if len(m) == 2 {
			moves = append(moves, m)
		}
	}

	// MovePart1(m, moves, sX, sY)
	MovePart2(m2, moves, sX, sY)

	for _, line := range m2 {
		for _, ch := range line {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}

	res := CalGPSPart2(m2)
	fmt.Println(res)
}

func MovePart1(m [][]rune, moves [][]int, sX, sY int) {
	var Move func(int, int, []int)
	Move = func(X, Y int, move []int) {
		ch := m[Y][X]
		nX, nY := X+move[0], Y+move[1]

		// fmt.Println(string(ch), "-", string(m[nY][nX]))

		if m[nY][nX] == '#' {
			return
		}

		if m[nY][nX] == '.' {
			m[nY][nX] = ch
			m[Y][X] = '.'

			if ch == '@' {
				sX, sY = nX, nY
			}

			return
		}

		if m[nY][nX] == 'O' {
			Move(nX, nY, move)

			if m[nY][nX] == '.' {
				m[nY][nX] = ch
				m[Y][X] = '.'

				if ch == '@' {
					sX, sY = nX, nY
				}

			}
		}
	}

	for _, move := range moves {
		// fmt.Println(move[0], move[1])
		Move(sX, sY, move)

		// for _, line := range m {
		// 	for _, ch := range line {
		// 		fmt.Print(string(ch))
		// 	}
		// 	fmt.Println()
		// }

	}
}

func MovePart2(m [][]rune, moves [][]int, sX, sY int) {
	var Move func(int, int, []int)
	Move = func(X, Y int, move []int) {
		ch := m[Y][X]
		dx := move[0]
		dy := move[1]
		nX, nY := X+move[0], Y+move[1]
		// fmt.Println(string(ch), "-", string(m[nY][nX]))

		if m[nY][nX] == '#' {
			return
		}

		if m[nY][nX] == '.' {
			m[nY][nX] = ch
			m[Y][X] = '.'

			if ch == '@' {
				sX, sY = nX, nY
			}

			return
		}

		moveQueue := make([][]int, 0)
		moveQueue = append(moveQueue, []int{X, Y})

		needMoveQueue := make([][]int, 0)
		needMoveQueue = append(needMoveQueue, []int{nX, nY})

		canMove := true
		for len(needMoveQueue) > 0 {
			need := needMoveQueue[0]
			needMoveQueue = needMoveQueue[1:]

			nX := need[0]
			nY := need[1]

			if m[nY][nX] == '.' {
				continue
			}

			if m[nY][nX] == '#' {
				canMove = false
				break
			}

			for m[nY][nX] != '.' {
				// fmt.Println(moveQueue)
				if m[nY][nX] == '#' {
					canMove = false
					break
				}

				if dy != 0 {
					moveQueue = append(moveQueue, []int{nX, nY})

					nnY, nnX := nY, nX+1
					if m[nY][nX] == ']' {
						nnX = nX - 1
					}

					if m[nnY+dy][nnX] != '#' {
						moveQueue = append(moveQueue, []int{nnX, nnY})
						needMoveQueue = append(needMoveQueue, []int{nnX + dx, nnY + dy})
					} else {
						canMove = false
						break
					}

				} else {
					moveQueue = append(moveQueue, []int{nX, nY})
				}

				nY += dy
				nX += dx
			}
		}

		if !canMove {
			return
		}

		runeMap := make(map[[2]int]rune)
		for _, p := range moveQueue {
			_, ok := runeMap[[2]int{p[0], p[1]}]
			if !ok {
				runeMap[[2]int{p[0], p[1]}] = m[p[1]][p[0]]
			}

			m[p[1]][p[0]] = '.'
		}

		for p, r := range runeMap {
			m[p[1]+dy][p[0]+dx] = r
		}

		sX += dx
		sY += dy
		return
	}

	for _, move := range moves {
		// fmt.Println(move[0], move[1])
		Move(sX, sY, move)

		// for _, line := range m {
		// 	for _, ch := range line {
		// 		fmt.Print(string(ch))
		// 	}
		// 	fmt.Println()
		// }

	}
}

func CalGPS(m [][]rune) int {
	res := 0
	for y, line := range m {
		for x, ch := range line {
			if ch == 'O' {
				gps := 100*y + x
				res += gps
			}
		}
	}

	return res
}

func CalGPSPart2(m [][]rune) int {
	res := 0
	for y, line := range m {
		for x, ch := range line {
			if ch == '[' {
				gps := 100*y + x
				res += gps
			}
		}
	}

	return res
}
