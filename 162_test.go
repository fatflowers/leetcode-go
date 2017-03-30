package main

import "testing"

func TestFindPeakElement(t *testing.T) {
    println(findPeakElement([]int{1,2,3,1}))
    println(findPeakElement([]int{1,2,3,4}))
    println(findPeakElement([]int{5,2,3,4}))
}