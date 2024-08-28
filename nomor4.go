package main

import (
	"fmt"
	"strconv"
	"strings"
)

func permute(nums []string) [][]string {
	var res [][]string
	var helper func([]string, int)
	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				} else {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				}
			}
		}
	}
	helper(nums, len(nums))
	return res
}

func evalEquation(equation string, mapping map[rune]rune) (bool, string) {
	translated := []rune{}
	for _, ch := range equation {
		if val, found := mapping[ch]; found {
			translated = append(translated, val)
		} else {
			translated = append(translated, ch)
		}
	}
	equationStr := string(translated)
	parts := strings.Fields(equationStr)

	left, err1 := strconv.Atoi(parts[0])
	right, err2 := strconv.Atoi(parts[2])
	result, err3 := strconv.Atoi(parts[4])

	if err1 != nil || err2 != nil || err3 != nil {
		return false, ""
	}

	switch parts[1] {
	case "+":
		return left+right == result, equationStr
	case "-":
		return left-right == result, equationStr
	}
	return false, ""
}

// Fungsi untuk menyelesaikan teka-teki
func solvePuzzle(equation string) string {
	uniqueCharsMap := map[rune]struct{}{}
	for _, ch := range equation {
		if ch >= 'A' && ch <= 'Z' {
			uniqueCharsMap[ch] = struct{}{}
		}
	}

	uniqueChars := []string{}
	for k := range uniqueCharsMap {
		uniqueChars = append(uniqueChars, string(k))
	}

	for _, perm := range permute([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
		mapping := make(map[rune]rune)
		for i, ch := range uniqueChars {
			mapping[rune(ch[0])] = rune(perm[i][0])
		}
		if valid, solution := evalEquation(equation, mapping); valid {
			return solution
		}
	}
	return "No solution"
}

func main() {
	// Contoh penggunaan
	input1 := "II + II = HIU"
	output1 := solvePuzzle(input1)
	fmt.Println("Example 1:")
	fmt.Printf("Input : %s\n", input1)
	fmt.Printf("Output : %s\n\n", output1)

	input2 := "ABD - AD = DKL"
	output2 := solvePuzzle(input2)
	fmt.Println("Example 2:")
	fmt.Printf("Input : %s\n", input2)
	fmt.Printf("Output : %s\n", output2)
}
