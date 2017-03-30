package main

import "testing"

func TestFindMin(t *testing.T) {
    println(findMin([]int{4,5,6,7,0,1,2}))
    println(findMin([]int{4,0,0,0,0,0,2}))
}