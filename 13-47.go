package main

import "strconv"

func reverseSlice(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func joinNum(nums []int) string {
	result := ""
	for i := range nums {
		result += strconv.Itoa(nums[i])
		if i != len(nums)-1 {
			result += ","
		}
	}
	return result
}

func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		println(joinNum(nums))
		return
	}

	firstPos := -1

	for i := len(nums) - 1; i > 0; i-- {
		if i-1 >= 0 && nums[i] > nums[i-1] {
			firstPos = i - 1
			break
		}
	}

	if firstPos == -1 {
		reverseSlice(nums)
		println(joinNum(nums))
		return
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[firstPos] {
			nums[i], nums[firstPos] = nums[firstPos], nums[i]
			break
		}
	}

	reverseSlice(nums[firstPos+1:])

	println(joinNum(nums))
}

//func main() {
//	nextPermutation([]int{1, 2, 3})
//	nextPermutation([]int{3, 2, 1})
//	nextPermutation([]int{1, 1, 5})
//}
