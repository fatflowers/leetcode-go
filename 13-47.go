package main

import "strconv"
import "reflect"

func reverseSlice(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	return nums
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

func getNextPermutation(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
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
		return nums
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[firstPos] {
			nums[i], nums[firstPos] = nums[firstPos], nums[i]
			break
		}
	}

	reverseSlice(nums[firstPos+1:])

	return nums
}

func nextPermutation(nums []int) {
	println(joinNum(getNextPermutation(nums)))
}

func permuteUnique(nums []int) [][]int {
	start := make([]int, len(nums))
	copy(start, nums)

	result := [][]int{nums}
	for {
		nums = getNextPermutation(nums)
		if reflect.DeepEqual(start, nums) {
			break
		}
		result = append(result, make([]int, len(nums)))
		copy(result[len(result) - 1], nums)
	}

	return result
}


//func main() {
//	//nextPermutation([]int{1, 2, 3})
//	//nextPermutation([]int{3, 2, 1})
//	//nextPermutation([]int{1, 1, 5})
//
//	fmt.Printf("%v\n", permuteUnique([]int{1,1,2}))
//}
