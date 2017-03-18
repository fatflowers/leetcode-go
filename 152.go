package main

func maxAll(a ...int) int {
	max := a[0]
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}

	return max
}

func minAll(a ...int) int {
	min := a[0]
	for i := range a {
		if a[i] < min {
			min = a[i]
		}
	}

	return min
}

func maxProduct(nums []int) int {
    if len(nums) == 0 {
    	return 0
    }

    previous := nums[0]
    previousMin := nums[0]
    result := nums[0]
    tmp := 0
    for i := 1; i < len(nums); i++ {
    	tmp = maxAll(previous*nums[i], previousMin*nums[i], nums[i])
    	previousMin = minAll(previous*nums[i], previousMin*nums[i], nums[i])
    	previous = tmp
    	result = maxAll(previous, result)
    }

    return result
}
