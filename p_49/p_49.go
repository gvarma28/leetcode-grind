package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/group-anagrams/
func main() {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	// res := groupAnagrams([]string{"a"})
	fmt.Println(res)
}

type Store struct {
	original string
	sorted   string
}

// CFOT
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	storeArr := []Store{}

	for _, val := range strs {
		// sort strings
		runes := []rune(val)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		// instantiating an struct with original and sorted strings
		curRecord := Store{
			original: val,
			sorted:   string(runes),
		}
		storeArr = append(storeArr, curRecord)
	}

	// sorting the storeArr based on the sorted strings
	sort.Slice(storeArr, func(i, j int) bool {
		return storeArr[i].sorted < storeArr[j].sorted
	})

	n := len(storeArr)
	curAnagramGroup := []string{storeArr[0].original}
	for i := 1; i < n; i++ {
		if storeArr[i].sorted == storeArr[i-1].sorted {
			curAnagramGroup = append(curAnagramGroup, storeArr[i].original)
			continue
		}
		res = append(res, curAnagramGroup)
		curAnagramGroup = []string{storeArr[i].original}
	}
	res = append(res, curAnagramGroup)

	return res
}
