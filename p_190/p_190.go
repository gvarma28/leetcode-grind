package main

import (
	"fmt"
)

// https://leetcode.com/problems/reverse-bits/description/
func main() {
	res := reverseBits(43261596)
	fmt.Println(res)
}

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		cur := num%2 == 0
		res = res << 1
		if !cur {
			res++
		}
		num = num >> 1
	}
	return res
}
