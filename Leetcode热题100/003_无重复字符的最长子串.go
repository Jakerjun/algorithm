package main

import "fmt"

// 滑动窗口 滑右边
func lengthOfLongestSubstring(s string) int {
	cnt := make([]int, 256)
	n := len(s)
	ans := 0
	for l, r := 0, 0; l < n; l++ {
		for r < n && cnt[s[r]] == 0 {
			cnt[s[r]]++
			r++
		}
		ans = max(ans, r-l)
		// r == n
		if r == n {
			break
		}
		cnt[s[l]]--
	}
	return ans
}

// 滑动窗口 滑左边
func lengthOfLongestSubstring2(s string) int {
	cnt := make([]int, 256)
	n := len(s)
	ans := 0
	l, r := 0, 0
	for ; r < n; r++ {
		cnt[s[r]]++
		ans = max(ans, r-l)
		for l < r && cnt[s[r]] > 1 {
			cnt[s[l]]--
			l++
		}
	}
	ans = max(ans, r-l)
	return ans
}

func main() {
	result := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)
}
