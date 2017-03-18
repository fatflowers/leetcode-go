package main

func isPalindrom(s string, start, end int) bool {
	if start < 0 {
		start = 0
	}
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// f(n) = max(f(n-1), Pal(len(s) - f(n-1), n)),
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	last := 1
	lastIndex := 0
	for i := 1; i < len(s); i ++ {
		if isPalindrom(s, i - last, i) {
			last = last + 1
			lastIndex = i
		} else if isPalindrom(s, i - last - 1, i) {
			last = last + 2
			lastIndex = i
		}
	}

	return s[lastIndex-last+1:lastIndex+1]
}

//func main() {
//	println(longestPalindrome("a"))
//	println(longestPalindrome("abbaaaaa"))
//	println(longestPalindrome("babad"))
//	println(longestPalindrome("cbbd"))
//}