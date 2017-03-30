package main

func findPeakElement(nums []int) int {
    if len(nums) == 0 {
        return -1
    } else if len(nums) == 1 {
        return 0
    }

    for i := range nums {
        if i < len(nums)-1 && nums[i] > nums[i+1] {
            return i
        }
    }

    return len(nums)-1
}