package main

func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func isPalindrome(r []rune) bool {
	if len(r) == 0 || len(r) == 1 {
		return true
	}

	for i := range r {
		if r[i] != r[len(r)-i-1] {
			return false
		}
	}

	return true
}

func searchPalindrome(r []rune, start int) int {
	minI := start
	for i := start-1; i >=0; i-- {
		if isPalindrome(r[i:start+1]) {
			minI = i
		}
	}

	return minI
}

/* 理解错题意了，没有“连续字符串”这两个字，感觉像是个0 1背包问题了。
func longestPalindromeSubseq(s string) int {
    result := 0 
    r := []rune(s)
    if len(r) ==0 || len(r) == 1 {
    	return len(r)
    }

    prePos := 0
    result = 1
    for i := range r {
    	if i == 0 {
    		continue
    	}
    	if prePos > 0 && r[i] == r[prePos]-1 {
    		prePos -= 1
    	} else {
    		prePos = searchPalindrome(r, i)
    	}
		result = maxInt(result, i-prePos+1)
    }

    return result
} */