package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	X, Y int
}

func main() {
	input, err := os.ReadFile("13.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n\n")

	games := make([][]Move, 0)
	for _, group := range lines {
		gameLines := strings.Split(group, "\n")
		if len(gameLines) < 3 {
			continue
		}

		game := make([]Move, 3)
		for i, line := range gameLines {
			if len(line) == 0 {
				continue
			}

			ints := GetInts(line)
			move := Move{
				X: ints[0],
				Y: ints[1],
			}

			game[i] = move
		}

		games = append(games, game)
	}

	res := 0
	for _, game := range games {
		res += PlayPart2(game)
	}

	fmt.Println(res)
}

// part 1
func PlayPart1(game []Move) int {
	aX, aY, bX, bY, tX, tY := game[0].X, game[0].Y, game[1].X, game[1].Y, game[2].X, game[2].Y

	min := 0

	ab := make([]int, 2)

	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			if a*aX+b*bX == tX && a*aY+b*bY == tY {
				cost := a*3 + b
				if cost < min || min == 0 {
					min = cost
					ab[0] = a
					ab[1] = b
				}
			}
		}
	}

	if ab[0] != 0 || ab[1] != 0 {
		fmt.Println("a: ", ab[0], " b: ", ab[1])
	}
	return min
}

// part 2
// math
// question like
// 69a + 27b = 18641 && 23a + 71b = 10279
// if a,b is int, then it is ok
func PlayPart2(game []Move) int {
	aX, aY, bX, bY, tX, tY := game[0].X, game[0].Y, game[1].X, game[1].Y, game[2].X, game[2].Y
	tX += 10000000000000
	tY += 10000000000000

	aXf, aYf, bXf, bYf, tXf, tYf := float64(aX), float64(aY), float64(bX), float64(bY), float64(tX), float64(tY)

	// aXf * a + bXf * b = tXf
	// aYf * a + bYf * b = tYf
	// aYf * a * (aXf / aYf) + bYf * b * (aXf / aYf) = tYf * (aXf / aYf)
	// bXf * b - bYf * b * (aXf / aYf) = tXf- tYf * (aXf / aYf)
	// (bXf - bYf * (aXf / aYf)) * b = tXf - tYf * (aXf / aYf)
	b := (tXf - tYf*(aXf/aYf)) / (bXf - bYf*(aXf/aYf))
	a := (tXf - b*bXf) / aXf

	bI := int(roundFloat(b, 0))
	aI := int(roundFloat(a, 0))

	if aI < 0 || bI < 0 || aX*aI+bX*bI != tX || aY*aI+bY*bI != tY {
		// fmt.Println("a: ", a, " b: ", b)
		return 0
	}

	fmt.Println("a: ", aI, " b: ", bI)
	return 3*aI + bI
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

// func Float2Int(f flot64) int {
// 	temp := fmt.Sprintf("%.2f", f)

// }
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
