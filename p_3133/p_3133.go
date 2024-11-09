package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-array-end/description/
func main() {
	res := minEnd(2, 7)
	fmt.Println(res)
}

func minEnd(n int, x int) int64 {
	var res int64 = 0

	xbits := []int{}

	for i := 0; i <= 64; i++ {
		if (x>>i)&1 == 1 {
			xbits = append(xbits, 1)
		} else {
			xbits = append(xbits, 0)
		}
	}

	fmt.Println(xbits)
	idx := 0
	for i := 0; i <= 64; i++ {
		if xbits[i] == 1 {
			res += (1 << i)
			continue
		}
		if ((n-1)>>idx)&1 == 1 {
			res += (1 << i)
			idx++
		}
	}

	// for i := 0; i < n; i++ {
	// 	for j := 0; i <= 64; j++ {
	// 		if xbits[i] == 1 {
	// 			continue
	// 		}
	// 	}
	// }


	return res

}
