package main

import "fmt"
import "testing"

func TestRemoveDuplicates(t *testing.T) {
	l := [][]int{{1,1,2,2,2,3}, {1,2,3}, {1,1,1,1,1,1}, {1,1,1,2}, {1,1,1,2,2,2}, {1, 2,2,2,2,2,2}}

	for i := range l {
		fmt.Println(l[i])
		fmt.Println(removeDuplicates(l[i]))
		fmt.Println(l[i])
		fmt.Println("===============")
	}
	
}