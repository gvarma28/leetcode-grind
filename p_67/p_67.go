package main

import (
	"fmt"
	// "strconv"
	"strings"
)

// https://leetcode.com/problems/add-binary/description/
func main() {
	res := addBinary("10100", "1111")
	fmt.Println(res)
}

func addBinary(a string, b string) string {
	var res strings.Builder

	lenA := len(a)
	lenB := len(b)

	carry := 0

	p := lenA - 1
	q := lenB - 1

	for {
		if p < 0 && q < 0 {
			break
		}

		sum := carry
		if p >= 0 && a[p] == '1' {
			sum += 1
		}
		if q >= 0 && b[q] == '1' {
			sum += 1
		}

		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
		if sum == 1 || sum == 3 {
			res.WriteString("1")
		} else {
			res.WriteString("0")
		}
		p--
		q--
	}

	if carry > 0 {
		res.WriteString("1")
	}

	return reverse(res.String())
}

func reverse(s string) string {
	runes := []rune(s)
	p, q := 0, len(runes)-1

	for {
		if p >= q {
			break
		}
		runes[p], runes[q] = runes[q], runes[p]
		p++
		q--
	}
	return string(runes)
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