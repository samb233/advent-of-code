package main

import (
	"fmt"
	"os"

	// "strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("12.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	m := make([][]rune, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		r := []rune(line)
		m = append(m, r)
	}

	res := PricePart2(m)
	fmt.Println(res)
}

type Plant struct {
	T rune
	X int
	Y int
}

// part 1
func Price(m [][]rune) int {
	plantMap := make(map[Plant]struct{})
	plantQueue := make([]Plant, 0)
	ml := len(m)

	adj4 := [4][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	var price func(Plant) (int, int)
	price = func(p Plant) (int, int) {
		_, ok := plantMap[p]
		if ok {
			return 0, 0
		}
		plantMap[p] = struct{}{}

		perimeter := 0
		area := 1

		for _, adj := range adj4 {
			newX, newY := p.X+adj[0], p.Y+adj[1]

			available := newX >= 0 && newX < ml && newY >= 0 && newY < ml
			if !available {
				perimeter += 1
				continue
			}

			newP := Plant{
				T: m[newY][newX],
				X: newX,
				Y: newY,
			}

			if newP.T == p.T {
				addP, addA := price(newP)
				perimeter += addP
				area += addA
			} else {
				perimeter += 1
				plantQueue = append(plantQueue, newP)
			}
		}

		return perimeter, area
	}

	firstP := Plant{
		T: m[0][0],
		X: 0,
		Y: 0,
	}

	plantQueue = append(plantQueue, firstP)
	res := 0
	for len(plantQueue) > 0 {
		p := plantQueue[0]
		plantQueue = plantQueue[1:]

		perimeter, area := price(p)
		res += perimeter * area

		if perimeter > 0 && area > 0 {
			fmt.Println("T: ", string(p.T), "P: ", perimeter, "A: ", area)
		}
	}

	return res
}

// part 2
type Side struct {
	X, Y, D int
}

func PricePart2(m [][]rune) int {
	plantMap := make(map[Plant]struct{})
	plantQueue := make([]Plant, 0)
	ml := len(m)

	adj4 := [4][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	sideMap := make(map[Side]struct{})
	var Area func(Plant) int
	Area = func(p Plant) int {
		_, ok := plantMap[p]
		if ok {
			return 0
		}
		plantMap[p] = struct{}{}

		area := 1

		for d, adj := range adj4 {
			side := Side{p.X, p.Y, d}
			newX, newY := p.X+adj[0], p.Y+adj[1]

			available := newX >= 0 && newX < ml && newY >= 0 && newY < ml
			if !available {
				sideMap[side] = struct{}{}
				continue
			}

			newP := Plant{
				T: m[newY][newX],
				X: newX,
				Y: newY,
			}

			if newP.T == p.T {
				area += Area(newP)
			} else {
				sideMap[side] = struct{}{}
				plantQueue = append(plantQueue, newP)
			}
		}

		return area
	}

	firstP := Plant{
		T: m[0][0],
		X: 0,
		Y: 0,
	}

	adj2 := [][]int{
		{1, 0},
		{0, 1},
	}

	plantQueue = append(plantQueue, firstP)
	res := 0
	for len(plantQueue) > 0 {
		p := plantQueue[0]
		plantQueue = plantQueue[1:]

		area := Area(p)
		if area == 0 {
			continue
		}

		trueSideMap := make(map[Side]struct{})
		for s := range sideMap {

			cnt := 0
			for _, adj := range adj2 {

				newX := s.X + adj[0]
				newY := s.Y + adj[1]

				newS := Side{
					X: newX,
					Y: newY,
					D: s.D,
				}

				if _, ok := sideMap[newS]; !ok {
					cnt++
				}
			}

			if cnt == 2 {
				trueSideMap[s] = struct{}{}
			}
		}

		side := len(trueSideMap)
		res += side * area
		sideMap = make(map[Side]struct{})

		if side > 0 && area > 0 {
			fmt.Println("T: ", string(p.T), "S: ", side, "A: ", area)
			for s := range trueSideMap {
				fmt.Print(s.X, ",")
				fmt.Print(s.Y, ",")
				fmt.Print(s.D, "\n")
			}
		}
	}

	return res
}
