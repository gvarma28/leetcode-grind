package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
func main() {
	// res := maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3)
	// res := maximumSubarraySum([]int{4, 4, 4}, 3)
	res := maximumSubarraySum([]int{1, 1, 1, 7, 8, 9}, 3)
	fmt.Println(res)
}

func maximumSubarraySum(nums []int, k int) int64 {

	sumArr := []int64{}
	n := len(nums)

	sumArr = append(sumArr, 0)
	for idx, val := range nums {
		if idx == 0 {
			sumArr = append(sumArr, int64(val))
			continue
		}
		sumArr = append(sumArr, int64(sumArr[idx])+int64(val))
	}

	counter := make(map[int]int)

	p, q := 0, 0
	count := 0
	for {
		if q >= k {
			break
		}
		if _, ok := counter[nums[q]]; !ok {
			count++
		}
		counter[nums[q]]++
		q++
	}

	var max int64 = 0
	if count == k && sumArr[q]-sumArr[p] > max {
		max = sumArr[q] - sumArr[p]
	}
	for {

		if q >= n {
			break
		}

		if counter[p] == 0 {
			delete(counter, counter[p])
		} else {
			counter[p]--
		}
		count--
		p++

		if _, ok := counter[nums[q]]; !ok {
			count++
		}
		counter[nums[q]]++
		q++
		fmt.Println(max, sumArr[q], sumArr[p], count)
		if count == k && sumArr[q]-sumArr[p] > max {
			max = sumArr[q] - sumArr[p]
		}
	}



	fmt.Println(p)
	fmt.Println(count)
	fmt.Println(counter)
	fmt.Println(sumArr)

	return max
}
