package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("10.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	mr := make([][]rune, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		lineRune := []rune(line)
		mr = append(mr, lineRune)
	}

	m := make([][]int, 0)
	ends := make([][]int, 0)
	for y, line := range mr {
		nums := make([]int, 0)

		for x, ch := range line {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}

			if num == 0 {
				end := []int{x, y}
				ends = append(ends, end)
			}

			nums = append(nums, num)
		}

		m = append(m, nums)
	}

	res := 0
	for _, end := range ends {
		plus := FoundRoute(m, end)
		fmt.Println(plus)
		res += plus
	}
	fmt.Println(res)
}

func FoundDist(m [][]int, start []int) int {
	queue := make([][]int, 1)
	queue[0] = start

	distMap := make(map[[2]int]struct{})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x := node[0]
		y := node[1]

		if m[y][x] == 9 {
			distMap[[2]int{x, y}] = struct{}{}
			continue
		}

		if y >= 1 && m[y-1][x]-m[y][x] == 1 {
			queue = append(queue, []int{x, y - 1})
		}

		if y < len(m)-1 && m[y+1][x]-m[y][x] == 1 {
			queue = append(queue, []int{x, y + 1})
		}

		if x >= 1 && m[y][x-1]-m[y][x] == 1 {
			queue = append(queue, []int{x - 1, y})
		}

		if x < len(m[y])-1 && m[y][x+1]-m[y][x] == 1 {
			queue = append(queue, []int{x + 1, y})
		}
	}

	return len(distMap)
}

func FoundRoute(m [][]int, end []int) int {
	num := 0
	x := end[0]
	y := end[1]

	if m[y][x] == 9 {
		return 1
	}

	if y >= 1 && m[y-1][x]-m[y][x] == 1 {
		num += FoundRoute(m, []int{x, y - 1})
	}

	if y < len(m)-1 && m[y+1][x]-m[y][x] == 1 {
		num += FoundRoute(m, []int{x, y + 1})
	}

	if x >= 1 && m[y][x-1]-m[y][x] == 1 {
		num += FoundRoute(m, []int{x - 1, y})
	}

	if x < len(m[y])-1 && m[y][x+1]-m[y][x] == 1 {
		num += FoundRoute(m, []int{x + 1, y})
	}

	return num
}
