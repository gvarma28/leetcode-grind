package main

import "fmt"
import "math"

// https://leetcode.com/problems/rearrange-characters-to-make-target-string/description/
func main() {
	var res int = rearrangeCharacters("ilovecodingonleetcode", "code")
	// var res int = rearrangeCharacters("abcba", "abc")
	// var res int = rearrangeCharacters("abbaccaddaeea", "aa")
	fmt.Println(res)
}

func rearrangeCharacters(s string, target string) int {
	// count characters in s
	s_counter := make(map[string]int)
	for _, v := range s {
		s_counter[string(v)]++
	}

	// count characters in target
	t_counter := make(map[string]int)
	for _, v := range target {
		t_counter[string(v)]++
	}

	res := math.MaxInt32
	for k, v := range t_counter {
		if _, exists := s_counter[k]; !exists {
			return 0
		}

		total := s_counter[k] / v
		if res > total {
			res = total
		}
	}
	return res
}
