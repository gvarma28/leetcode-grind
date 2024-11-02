package main

import (
	"fmt"
	"strings"
)

// problem_url
func main() {
	res := addBinary("11", "1")
	fmt.Println(res)
}

func addBinary(a string, b string) string {
	res := ""

	lenA := len(a)
	lenB := len(b)

	carry := 0

	for i := 0; i < max(lenA, lenB); i++ {
		curA := ""
		curB := ""
		
		if i < lenA && string(a[i]) == "1" {
			curA = string(a[i])
		}
		if i < lenB {
			curB = string(b[i])
		}

		


	}


	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b;
}

// func binaryToDecimal(a string) int64 {
// 	var res int64 = 0
// 	n := len(a)
//     var temp int64 = 1
// 	for i := 0; i < n; i++ {
// 		if string(a[n-i-1]) == "1" {
// 			res += temp << i
// 		}
// 	}
// 	return res
// }

// func decimalToBinary(a int64) string {
// 	var res strings.Builder
// 	var temp int64 = 1
// 	for i := 62; i >= 0; i-- {
//         fmt.Println(temp << i)
// 		if (temp << i) <= a {
// 			res.WriteString("1")
// 			a -= (temp << i)
// 			continue
// 		}
// 		res.WriteString("0")
// 	}
// 	return strings.TrimLeft(res.String(), "0")
// }