package main
/*
func minTwoInt(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func findMin(nums []int) int {
    if len(nums) == 0 {
        return -1
    } else if len(nums) == 1 {
        return nums[0]
    } else if len(nums) == 2 {
        return minTwoInt(nums[0], nums[1])
    }

    half := len(nums)/2

    if nums[half] > nums[half+1] {
        return nums[half+1]
    } else if nums[half] < nums[half-1] {
        return nums[half]
    } else {
        a := findMin(nums[:half])
        b := findMin(nums[half+1:])
        if a != -1 {
            if b != -1 {
                return minTwoInt(a, b)
            } else {
                return a
            }
        } else {
            return b
        }
    }
}*/