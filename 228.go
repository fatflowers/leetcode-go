package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	result := []string{}

	if len(nums) == 0 {
		return result
	}

	makeString := func(start, end int) string {
		if start == end {
			return fmt.Sprintf("%d", start)
		} else {
			return fmt.Sprintf("%d->%d", start, end)
		}
	}

	start := 0
	for end := range nums {
		if end < len(nums)-1 {
			if nums[end+1]-nums[end] != 1 {
				result = append(result, makeString(nums[start], nums[end]))
				start = end + 1
			}
		} else {
			result = append(result, makeString(nums[start], nums[end]))
		}
	}
	return result
}

// func main() {
// 	nums := []int{0}
// 	fmt.Printf("%v\n", summaryRanges(nums))
// }
