package main

import "testing"
import "fmt"

func TestBit(t *testing.T) {
	println(15 << 1)
}

func TestGrayCode(t *testing.T) {
	fmt.Println(grayCode(1))
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(3))
	fmt.Println(grayCode(4))
}