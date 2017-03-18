package main

import "testing"
import "reflect"

func TestProductExceptSelf(t *testing.T) {
	if !reflect.DeepEqual(productExceptSelf([]int{1,2,3,4}), []int{24,12,8,6}) {
		t.Error([]int{1,2,3,4}, productExceptSelf([]int{1,2,3,4}))
	}

	if !reflect.DeepEqual(productExceptSelf([]int{1}), []int{1}) {
		t.Error([]int{1})
	}

	if !reflect.DeepEqual(productExceptSelf([]int{0, 0}), []int{0, 0}) {
		t.Error([]int{0, 0})
	}
}