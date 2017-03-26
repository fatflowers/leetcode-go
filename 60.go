package main

import "strconv"

func reverseIntList(list []int) []int {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
	}

	return list
}

func nextPerm(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)

	firstPos := -1
	for i := len(result)-1; i >= 1; i-- {
		if result[i] > result[i-1] {
			firstPos = i-1
			break
		}
	}

	if firstPos == -1 {
		return reverseIntList(result)
	}

	for i := len(result)-1; i >= 1; i-- {
		if result[i] > result[firstPos] {
			result[i], result[firstPos] = result[firstPos], result[i]
			break
		}
	}

	reverseIntList(result[firstPos+1:])

	return result
}

func getPermutation(n int, k int) string {
	list := make([]int, n)
    for i := 1; i <= n; i ++ {
    	list[i-1] = i
    }

    for i := 1; i < k ; i++ {
    	list = nextPerm(list)
    }

    result := ""
    for i := range list {
    	result += strconv.Itoa(list[i])
    }

    return result
}