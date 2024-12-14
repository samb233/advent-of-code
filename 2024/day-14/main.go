package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var Lx int = 101
var Ly int = 103

type Robot struct {
	X, Y, Vx, Vy int
}

func (r *Robot) Tick() {
	r.X += r.Vx
	r.Y += r.Vy

	if r.X >= Lx {
		r.X -= Lx
	}

	if r.X < 0 {
		r.X += Lx
	}

	if r.Y >= Ly {
		r.Y -= Ly
	}

	if r.Y < 0 {
		r.Y += Ly
	}
}

func (r *Robot) Quadrant() int {
	Mx := (Lx - 1) / 2
	My := (Ly - 1) / 2

	if r.X < Mx && r.Y < My {
		return 1
	}

	if r.X > Mx && r.Y < My {
		return 2
	}

	if r.X < Mx && r.Y > My {
		return 3
	}

	if r.X > Mx && r.Y > My {
		return 4
	}

	return 0
}

func main() {
	input, err := os.ReadFile("14.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	robots := make([]Robot, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		ints := GetInts(line)
		robot := Robot{
			ints[0],
			ints[1],
			ints[2],
			ints[3],
		}

		// fmt.Println(robot.X, robot.Y, robot.Vx, robot.Vy)
		robots = append(robots, robot)
	}

	res := TickToTree(robots, 1)
	fmt.Println(res)
}

func TickSeconds(robots []Robot, t int) int {
	for i := 0; i < t; i++ {
		for r := 0; r < len(robots); r++ {
			robots[r].Tick()
		}
	}

	qMap := make(map[int]int)
	for _, robot := range robots {
		qMap[robot.Quadrant()]++
	}

	fmt.Println(qMap[1], qMap[2], qMap[3], qMap[4])

	return qMap[1] * qMap[2] * qMap[3] * qMap[4]
}

// print graph to file and search "########"
func TickToTree(robots []Robot, t int) int {
	file, err := os.OpenFile("./tree.out", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}

	second := 0

	for t > 0 {
		second++
		fmt.Fprintln(file, second)
		fmt.Println(second)

		lMap := make(map[int][]Robot)
		for r := 0; r < len(robots); r++ {
			robots[r].Tick()

			rs, ok := lMap[robots[r].Y]
			if !ok {
				rs = make([]Robot, 0)
			}

			rs = append(rs, robots[r])
			lMap[robots[r].Y] = rs
		}

		ls := make([]int, 0)
		for y := range lMap {
			ls = append(ls, y)
		}

		sort.Ints(ls)

		last := 0
		tree := true
		for i, y := range ls {
			if i == 0 {
				last = y
				continue
			}

			if y-last != 1 {
				tree = false
			}
		}

		for y := 0; y < Ly; y++ {
			xMap := make(map[int]bool)
			rs, ok := lMap[y]
			if ok {
				for _, r := range rs {
					xMap[r.X] = true
				}
			}

			for x := 0; x < Lx; x++ {
				if xMap[x] {
					fmt.Fprint(file, "#")
				} else {
					fmt.Fprint(file, ".")
				}
			}

			fmt.Fprintln(file)
		}

		if tree {
			t--

			fmt.Println(t)
		}
	}

	return second
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
