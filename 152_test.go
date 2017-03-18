package main

import "testing"


func TestMaxAll(t *testing.T) {
	if maxAll(1,2,3,4) != 4 {
		t.Error("1,2,3,4")
	}
}

func TestMaxProduct(t *testing.T) {
	if maxProduct([]int{2,3,-2,4}) != 6 {
		t.Error("2,3,-2,4", maxProduct([]int{2,3,-2,4}))
	}

	if maxProduct([]int{0,0,0,0}) != 0 {
		t.Error("0,0,0,0")
	}

	if maxProduct([]int{3,-1,4}) != 4 {
		t.Error("3,-1,4")
	}

	if n := maxProduct([]int{-2,1,-4}); n != 8 {
		t.Error("-2,1,-4", 8, n)
	}

	if n := maxProduct([]int{2,-5,-2,-4,3}); n != 24 {
		t.Error("2,-5,-2,-4,3", 24, n)
	}

	if n := maxProduct([]int{1,0,-1,2,3,-5,-2}); n != 60 {
		t.Error("1,0,-1,2,3,-5,-2", 60, n)
	}
}