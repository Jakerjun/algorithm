package main

import "fmt"

func findAnagrams(s string, p string) []int {
	cnt := make([]int, 26)
	for _, val := range p {
		cnt[val-'a']--
	}

	ans := make([]int, 0)
	n := len(s)
	m := len(p)

	// 题目用例 存在 len(s) < len(p) 的情况 直接特判掉就行
	if n < m {
		return []int{}
	}

	// 思路： 先构造一个长度为 m - 1 的窗口
	// 每次遍历的时候 将当前位置的值加入窗口
	// 最左侧值去掉窗口，然后每次都检测
	for i := range m - 1 {
		cnt[s[i]-'a']++
	}

	// 主逻辑 进入滑窗 退出滑窗
	for left, right := 0, m-1; right < n; left, right = left+1, right+1 {
		// 最右侧值进入窗口
		cnt[s[right]-'a']++
		// 检测窗口是否满足
		if isCheck(cnt) {
			ans = append(ans, left)
		}
		// 当前值为最右值的检测结束 可以将最左的值出滑窗了
		cnt[s[left]-'a']--
	}
	return ans
}

func isCheck(cnt []int) bool {
	for _, x := range cnt {
		if x != 0 {
			return false
		}
	}
	return true
}

func main() {
	result := findAnagrams("cbaebabacd", "abc")
	fmt.Println(result)
}
