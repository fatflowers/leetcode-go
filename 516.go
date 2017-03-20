package main

import "fmt"

/*func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}*/

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

/* 理解错题意了，没有“连续”这两个字，感觉像是个0 1背包问题了。
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



func longestPalindromeSubseq(s string) int {
    r := []rune(s)
    if len(r) ==0 || len(r) == 1 {
    	return len(r)
    }

    subSeq := [][]int{}
    for i := range r {
    	subSeq = append(subSeq, make([]int, len(r)-i))
    	subSeq[i][0] = 1
    }

    for i := 1; i < len(r); i++ {
    	for j := range r {
    		if i-j < len(subSeq[j]) {
    					if j ==3 {
	    					println(i,j)
	    				}
    			for k := j; k < i; k ++ {

    				if r[k] == r[i] {
    					if i == k {
    						subSeq[j][i-j] = maxInt(subSeq[j][i-j], 1)
    					} else if i-1==k {
    						subSeq[j][i-j] = maxInt(subSeq[j][i-j], 2)
    					} else if k+1 == i-j-1 {
    						subSeq[j][i-j] = maxInt(subSeq[j][i-j], 3)
    					} else {
    						subSeq[j][i-j] = maxInt(subSeq[j][i-j], 2 + subSeq[k+1][i-k-2])
    					}
    				}
    			}

    			subSeq[j][i] = maxInt(subSeq[j][i], subSeq[j][i-1])
    		} else {
    			break
    		}
    	}
    }
    fmt.Printf("%v\n", subSeq)

    return subSeq[0][len(r)-1]
}





