package main

import (
	"fmt"
	"os"
	// "sort"
	"strconv"
	"strings"
	// "sync"
	// "runtime"
)

// part 1
func main() {
	input, err := os.ReadFile("17.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n\n")
	program := GetInts(lines[1])

	registers := strings.Split(lines[0], "\n")
	rA := GetInts(registers[0])
	rB := GetInts(registers[1])
	rC := GetInts(registers[2])

	fmt.Println(program)
	fmt.Println(rA)
	fmt.Println(rB)
	fmt.Println(rC)

	ops := []Operation{
		Op0, Op1, Op2, Op3, Op4, Op5, Op6, Op7,
	}

	pointer := 0
	for {
		if pointer+1 >= len(program) {
			break
		}

		p := program[pointer]
		oprand := program[pointer+1]
		op := ops[p]
		jump := op(rA, rB, rC, oprand)
		if jump == -1 || p == 5 {
			pointer += 2
		} else {
			pointer = jump
		}

		if p == 5 {
			fmt.Print(jump, ",")
		}
	}

	fmt.Println(rA)
	fmt.Println(rB)
	fmt.Println(rC)
}

// part 2
// func main() {
// 	cpuNum := runtime.NumCPU()
// 	runtime.GOMAXPROCS(cpuNum)

// 	input, err := os.ReadFile("17.in")
// 	if err != nil {
// 		panic(err)
// 	}

// 	lines := strings.Split(string(input), "\n\n")
// 	program := GetInts(lines[1])

// 	registers := strings.Split(lines[0], "\n")
// 	rA := GetInts(registers[0])
// 	rB := GetInts(registers[1])
// 	rC := GetInts(registers[2])

// 	fmt.Println(program)
// 	fmt.Println(rA)
// 	fmt.Println(rB)
// 	fmt.Println(rC)

// 	ops := []Operation{
// 		Op0, Op1, Op2, Op3, Op4, Op5, Op6, Op7,
// 	}

// 	g := cpuNum
// 	close := 0
// 	trueA := 0

// 	start := 0
// 	var wg sync.WaitGroup
// 	for i := 1; i <= g; i++ {
// 		wg.Add(1)
// 		go func(startPlus int) {
// 			rightA := start + startPlus
// 			for {
// 				if close == 1 {
// 					break
// 				}

// 				newA := []int{rightA}
// 				newB := make([]int, 1)
// 				newC := make([]int, 1)

// 				fmt.Println(rightA)
// 				compareLen := -1
// 				pointer := 0
// 				for {
// 					if pointer+1 >= len(program) {
// 						break
// 					}

// 					p := program[pointer]
// 					oprand := program[pointer+1]

// 					op := ops[p]
// 					jump := op(newA, newB, newC, oprand)
// 					if jump == -1 || p == 5 {
// 						pointer += 2
// 					} else {
// 						pointer = jump
// 					}

// 					if p == 5 {
// 						compareLen++
// 						if jump != program[compareLen] {
// 							break
// 						} else {
// 							if compareLen == len(program)-1 {
// 								close = 1
// 								trueA = rightA
// 							}
// 						}
// 					}
// 				}

// 				rightA += g + 1
// 			}
// 			wg.Done()
// 		}(i)
// 	}

// 	wg.Wait()
// 	fmt.Println(trueA)
// }

func ComboOprand(A []int, B []int, C []int, oprand int) int {
	if oprand <= 3 {
		return oprand
	} else if oprand == 4 {
		return A[0]
	} else if oprand == 5 {
		return B[0]
	} else if oprand == 6 {
		return C[0]
	}

	panic(fmt.Sprintf("Invalid combo oprand: %d", oprand))
}

type Operation func([]int, []int, []int, int) int

func Op0(A []int, B []int, C []int, oprand int) int {
	num := A[0]
	input := ComboOprand(A, B, C, oprand)
	devide := 1 << input
	res := num / devide
	A[0] = res

	return -1
}

func Op1(A []int, B []int, C []int, oprand int) int {
	num := B[0]
	res := num ^ oprand
	B[0] = res

	return -1
}

func Op2(A []int, B []int, C []int, oprand int) int {
	input := ComboOprand(A, B, C, oprand)
	res := input % 8
	B[0] = res

	return -1
}

func Op3(A []int, B []int, C []int, oprand int) int {
	numA := A[0]
	if numA == 0 {
		return -1
	} else {
		return oprand
	}
}

func Op4(A []int, B []int, C []int, oprand int) int {
	num1 := B[0]
	num2 := C[0]
	res := num1 ^ num2
	B[0] = res

	return -1
}

func Op5(A []int, B []int, C []int, oprand int) int {
	input := ComboOprand(A, B, C, oprand)
	res := input % 8
	return res
}

func Op6(A []int, B []int, C []int, oprand int) int {
	num := A[0]
	input := ComboOprand(A, B, C, oprand)
	devide := 1 << input
	res := num / devide
	B[0] = res

	return -1
}

func Op7(A []int, B []int, C []int, oprand int) int {
	num := A[0]
	input := ComboOprand(A, B, C, oprand)
	devide := 1 << input
	res := num / devide
	C[0] = res

	return -1
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
