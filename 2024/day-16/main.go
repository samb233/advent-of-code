package main

import (
	"fmt"
	"os"
	// "sort"
	// "strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("16.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	sX, sY, eX, eY := 0, 0, 0, 0
	m := make([][]rune, 0)
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}

		mR := []rune(line)
		m = append(m, mR)

		for x, ch := range line {
			if ch == 'S' {
				sX, sY = x, y
			}

			if ch == 'E' {
				eX, eY = x, y
			}
		}
	}

	PlayPart2(m, sX, sY, eX, eY)
}

var adj4 [4][]int = [4][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func PlayPart1(m [][]rune, sX, sY, eX, eY int) {
	distMap := make(map[[3]int]int)
	distMap[[3]int{sX, sY, 0}] = 0

	// seen := make(map[[3]int]bool)
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{sX, sY, 0})

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		X := node[0]
		Y := node[1]
		dir := node[2]
		dist := distMap[node]

		if node[0] == eX && node[1] == eY {
			fmt.Println(dist)
		}

		for i, adj := range adj4 {
			nX, nY := X+adj[0], Y+adj[1]
			if m[nY][nX] == '#' {
				continue
			}

			turn := 0
			diff := Abs(i - dir)
			if diff == 1 || diff == 3 {
				turn = 1
			} else if diff == 2 {
				turn = 2
			}

			newDist := dist + 1 + 1000*turn
			newNode := [3]int{nX, nY, i}
			oldDist, ok := distMap[newNode]
			if !ok || newDist < oldDist {
				distMap[newNode] = newDist
				queue = append(queue, newNode)
			}
		}
	}

	return
}

func PlayPart2(m [][]rune, sX, sY, eX, eY int) {
	distMap := make(map[[3]int]int)
	distMap[[3]int{sX, sY, 0}] = 0

	queue := make([][3]int, 0)
	queue = append(queue, [3]int{sX, sY, 0})

	fromMap := make(map[[3]int][][3]int)

	minEnd := [3]int{}
	min := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		X := node[0]
		Y := node[1]
		dir := node[2]
		dist := distMap[node]

		if node[0] == eX && node[1] == eY {
			if min == 0 || dist < min {
				min = dist
				minEnd = [3]int{eX, eY, dir}
			}
		}

		for i, adj := range adj4 {
			nX, nY := X+adj[0], Y+adj[1]
			if m[nY][nX] == '#' {
				continue
			}

			turn := 0
			diff := Abs(i - dir)
			if diff == 1 || diff == 3 {
				turn = 1
			} else if diff == 2 {
				turn = 2
			}

			newDist := dist + 1 + 1000*turn
			newNode := [3]int{nX, nY, i}
			oldDist, ok := distMap[newNode]

			inserted := false
			if !ok || newDist < oldDist {
				distMap[newNode] = newDist
				queue = append(queue, newNode)

				from := make([][3]int, 0)
				from = append(from, node)
				fromMap[newNode] = from
				inserted = true
			}

			if !inserted && newDist == oldDist {
				from, _ := fromMap[newNode]
				from = append(from, node)
				fromMap[newNode] = from
			}
		}
	}

	queue = make([][3]int, 1)
	queue[0] = minEnd

	seen := make(map[[2]int]struct{})

	for len(queue) > 0 {
		node := queue[0]
		// fmt.Println(queue)
		// fmt.Println(distMap[node])

		queue = queue[1:]
		tX, tY := node[0], node[1]
		if tX == sX && tY == sY {
			break
		}

		from := fromMap[node]
		for _, fromNode := range from {
			fX, fY := fromNode[0], fromNode[1]
			seen[[2]int{fX, fY}] = struct{}{}
			queue = append(queue, fromNode)
		}
	}

	fmt.Println(len(seen) + 1)
	return
}

// tools
func Abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
