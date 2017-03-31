package main
import "fmt"
/*

 A, C, G, T
Given s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",

Return:
["AAAAACCCCC", "CCCCCAAAAA"].
*/
func findRepeatedDnaSequences(s string) []string {
    if len(s) < 20 {
        return []string{}
    }

    tmp := [][]int{}

    for i := 0; i < len(s) - 10; i ++ {
        tmpEle := []int{i}
        for j := i+1; j < len(s)-10; j++ {
            if s[i:i+10] == s[j:j+10] {
                tmpEle = append(tmpEle, j)
            }
        }
        if len(tmpEle) > 1 {
            tmp = append(tmp, tmpEle)
        }
    }

    fmt.Println(tmp)
    // [[0 10] [5 16]]

    tmp2 := [][]int{}
    length := 10
    for i := range tmp {
        result[s[tmp[i][0]:tmp[i][0]+length]] = struct{}{}
        
    }


    return []string{}

}