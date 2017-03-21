package main

func grayCode(n int) []int {
    if n == 0 {
    	return []int{0}
    }

	result := []int{0, 1}
    
    for i := uint32(1); i < uint32(n); i++ {
    	tmp := []int{}
    	for j := range result {
    		tmp = append(tmp, result[len(result)-j-1]|1<<i)
    	}
    	result = append(result, tmp...)
    }


    return result
}

/*
0 000
1 001
3 011
2 010
6 110
7 111
5 101
4 100
*/