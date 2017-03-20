package main

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	data := map[string]int{
		//"": 0,
		//"abaa": 3,
		//"aaaaa": 5,
		//"abcd": 1,
		//"bbbab": 4,
		//"abab": 3,
		//"abcabcabcabc": 3,
		"aabaaba": 6,
		}

	for k, v := range data {
		if n := longestPalindromeSubseq(k); n != data[k] {
			t.Errorf("s: %s, right: %d, yours: %d", k, v, n)
		}
	}
}