package main

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// 欠债表
	cnt := make([]int, 256)
	for _, val := range t {
		cnt[val]--
	}
	n, m := len(s), len(t)
	start, size := 0, n+1
	for l, r, dept := 0, 0, m; r < n; r++ {
		cnt[s[r]]++
		if cnt[s[r]] <= 0 {
			dept--
		}
		if dept == 0 {
			// s[l] 这个字符不欠债 而且还有多的
			// cnt[s[l]] == 0 的时候 表示这个字符刚好还清
			for cnt[s[l]] > 0 {
				cnt[s[l]]--
				l++
			}
			if r-l+1 < size {
				size = r - l + 1
				start = l
			}
		}
	}

	if size == n+1 {
		return ""
	}

	return s[start : start+size]
}

func main() {
	result := minWindow("ADOBECODEBANC", "ABC")
	println(result)
}
