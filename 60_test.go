package main

import "testing"

func TestGetPermutation(t *testing.T) {
	t.Log(getPermutation(4, 1))
	t.Log(getPermutation(4, 2))
	t.Log(getPermutation(4, 3))
	t.Log(getPermutation(4, 4))
	t.Log(getPermutation(4, 5))
	t.Log(getPermutation(4, 6))
}