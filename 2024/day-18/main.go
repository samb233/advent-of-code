package main

import (
	"fmt"
	"os"
	// "sort"
	"strconv"
	"strings"
)

// part 1
func main() {
	input, err := os.ReadFile("18.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	coors := make([][]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		coor := GetInts(line)
		coors = append(coors, coor)
	}

	sX, sY := 0, 0
	eX, eY := 70, 70

	res := GoPart2(coors, sX, sY, eX, eY)
	fmt.Println(res)
}

var dirs [][]int = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func GoPart1(c [][]int, sX, sY, eX, eY int) int {
	cMap := make(map[[2]int]bool)

	for i := 0; i < 1024; i++ {
		cc := c[i]
		fmt.Println(cc)
		cMap[[2]int{cc[0], cc[1]}] = true
	}

	distMap := make(map[[2]int]int)
	distMap[[2]int{sX, sY}] = 0

	queue := make([][2]int, 1)
	queue[0] = [2]int{sX, sY}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		X, Y := node[0], node[1]
		dist := distMap[node]

		for _, dir := range dirs {
			nX, nY := X+dir[0], Y+dir[1]
			nNode := [2]int{nX, nY}

			if nX < 0 || nY < 0 || nX > eX || nY > eY {
				continue
			}

			fmt.Println(nNode)

			if cMap[nNode] {
				continue
			}

			nDist := dist + 1
			oDist, ok := distMap[nNode]
			if !ok || nDist < oDist {
				distMap[nNode] = nDist
				queue = append(queue, nNode)
			}
		}
	}

	res := distMap[[2]int{eX, eY}]
	return res
}

func GoPart2(c [][]int, sX, sY, eX, eY int) int {
	cMap := make(map[[2]int]bool)

	for t := 1024; t < len(c); t++ {
		fmt.Println(t)
		for i := 0; i < t; i++ {
			cc := c[i]
			cMap[[2]int{cc[0], cc[1]}] = true
		}

		distMap := make(map[[2]int]int)
		distMap[[2]int{sX, sY}] = 0

		queue := make([][2]int, 1)
		queue[0] = [2]int{sX, sY}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			X, Y := node[0], node[1]
			dist := distMap[node]

			for _, dir := range dirs {
				nX, nY := X+dir[0], Y+dir[1]
				nNode := [2]int{nX, nY}

				if nX < 0 || nY < 0 || nX > eX || nY > eY {
					continue
				}

				if cMap[nNode] {
					continue
				}

				nDist := dist + 1
				oDist, ok := distMap[nNode]
				if !ok || nDist < oDist {
					distMap[nNode] = nDist
					queue = append(queue, nNode)
				}
			}
		}

		res := distMap[[2]int{eX, eY}]
		if res == 0 {
			fmt.Println(c[t - 1])
			return 0
		}
	}

	return 0
}

// tools
func GetInts(s string) []int {
	nums := make([]int, 0)
	numS := make([]rune, 0)
	for i, c := range s {
		cIsNum := c >= '0' && c <= '9'
		cIsMark := len(numS) == 0 && (c == '-' || c == '+')
		if cIsNum || cIsMark {
			numS = append(numS, c)
		} else {
			if len(numS) == 0 {
				continue
			}

			num, _ := strconv.Atoi(string(numS))
			nums = append(nums, num)
			numS = make([]rune, 0)
		}

		if i == len(s)-1 {
			if len(numS) > 0 {
				num, _ := strconv.Atoi(string(numS))
				nums = append(nums, num)
			}
		}
	}

	return nums
}
