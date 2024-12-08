package main

import (
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func main() {
	input, err := os.ReadFile("8.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	m := make([][]rune, len(lines))
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		r := []rune(line)
		m[i] = r
	}

	aMap := make(map[rune][]Coordinates)
	for y, line := range m {
		for x, ch := range line {
			if ch == '.' {
				continue
			}

			coordinates, ok := aMap[ch]
			if ok {
				coordinates = append(coordinates, Coordinates{X: x, Y: y})
			} else {
				coordinates = []Coordinates{Coordinates{X: x, Y: y}}
			}
			aMap[ch] = coordinates
		}
	}

	res := CountAntinodesPart2(aMap, len(m[0]))
	fmt.Println(res)
}

// part 1
func CountAntinodes(aMap map[rune][]Coordinates, l int) int {
	// fmt.Println(l)
	nMap := make(map[Coordinates]bool)
	for _, coors := range aMap {
		if len(coors) == 1 {
			continue
		}

		for i := 0; i < len(coors); i++ {
			for j := i + 1; j < len(coors); j++ {
				a := coors[i]
				b := coors[j]

				a1 := Coordinates{
					X: a.X + (a.X - b.X),
					Y: a.Y + (a.Y - b.Y),
				}

				a2 := Coordinates{
					X: b.X + (b.X - a.X),
					Y: b.Y + (b.Y - a.Y),
				}

				if a1.X >= 0 && a1.X < l && a1.Y >= 0 && a1.Y < l {
					// fmt.Println(string(ch), ": ", a1.X, ",", a1.Y)
					nMap[a1] = true
				}

				if a2.X >= 0 && a2.X < l && a2.Y >= 0 && a2.Y < l {
					// fmt.Println(string(ch), ": ", a2.X, ",", a2.Y)
					nMap[a2] = true
				}
			}
		}
	}

	return len(nMap)
}

// part 2
func CountAntinodesPart2(aMap map[rune][]Coordinates, l int) int {
	// fmt.Println(l)
	nMap := make(map[Coordinates]bool)
	for _, coors := range aMap {
		if len(coors) == 1 {
			continue
		}

		for i := 0; i < len(coors); i++ {
			for j := i + 1; j < len(coors); j++ {
				a := coors[i]
				b := coors[j]

				nMap[a] = true
				nMap[b] = true

				a1 := Coordinates{
					X: a.X + (a.X - b.X),
					Y: a.Y + (a.Y - b.Y),
				}

				for a1.X >= 0 && a1.X < l && a1.Y >= 0 && a1.Y < l {
					nMap[a1] = true
					// fmt.Println(a1.X, ",", a1.Y)

					a1.X += (a.X - b.X)
					a1.Y += (a.Y - b.Y)
				}

				a2 := Coordinates{
					X: b.X + (b.X - a.X),
					Y: b.Y + (b.Y - a.Y),
				}


				for a2.X >= 0 && a2.X < l && a2.Y >= 0 && a2.Y < l {
					nMap[a2] = true
					// fmt.Println(a2.X, ",", a2.Y)

					a2.X += (b.X - a.X)
					a2.Y += (b.Y - a.Y)
				}
			}
		}
	}

	return len(nMap)
}
