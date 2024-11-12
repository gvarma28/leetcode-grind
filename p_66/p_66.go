package main

import (
	"fmt"
)

// problem_url
func main() {
	res := plusOne([]int{1, 2, 3})
	fmt.Println(res)
}

func plusOne(digits []int) []int {
	n := len(digits)

	// reverse the array
	digits = reverse(digits)

	broken := false
	for i := 0; i < n; i++ {
		if digits[i] < 9 {
			digits[i]++
			broken = true
			break
		} else {
			digits[i] = 0
		}
	}

	if !broken {
		digits = append(digits, 1)
	}

	// reverse the array
	digits = reverse(digits)

	return digits
}

func reverse(arr []int) []int {
	n := len(arr)
	p, q := 0, n-1
	for {
		if p >= q {
			break
		}
		arr[p], arr[q] = arr[q], arr[p]
		p++
		q--
	}
	return arr
}
