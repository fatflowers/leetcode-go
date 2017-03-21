package main

func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
    	return len(nums)
    }

    validP, curN, curTimes := 0, nums[0], 1

    for i := range nums {
    	if i == 0 {
    		continue
    	}
    	if nums[i] == curN {
    		curTimes ++
    		if curTimes > 2 && validP == 0 {
    			validP = i
    		} else if curTimes <= 2 && validP != 0 {
    			nums[validP] = nums[i]
    			validP ++
    		}
    	} else {
    		if validP != 0{
	    		nums[validP] = nums[i]
				validP ++
	    	}
	    	curN = nums[i]
	    	curTimes=1
    	} 	
    }

    if validP == 0 {
    	return len(nums)
    }

    return validP
}