package main

import (
	"fmt"
)

// https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/description/
func main() {
	res := takeCharacters("aabaaaacaabc", 2)
	fmt.Println(res)
}

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	counter := [3]int{}
	n := len(s)

	for i := 0; i < n; i++ {
		idx := s[i] - 'a'
		counter[idx]++
	}

	for _, val := range counter {
		if val < k {
			return -1
		}
	}

	curCount := [3]int{}
	p, q := 0, 0

	max := 0
	for {
		if p >= n || q >= n {
			break
		}
		idx := s[q] - 'a'
		curCount[idx]++
		if counter[idx] - curCount[idx] < 2 {
			if max < q - p {
				max = q-p
			}
			p = q+1
			curCount = [3]int{}
		}
		q++
	}

	fmt.Println(counter, curCount)
	return n-max
}
