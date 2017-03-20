package main

import "testing"
import "fmt"

func TestMerge(t *testing.T) {
	data := map[string][]Interval {
		"1": []Interval{{1,3}, {2,6}, {8,10}, {15,18}},
		"2": []Interval{{4,5}, {1,2}},
		"3": []Interval{{1,4}, {0,2}, {3,5}},
	}


	for k, v := range data {
		fmt.Println(k, v)
		fmt.Println(merge(v))
	}
}