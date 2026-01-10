package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	deque := make([]int, n)
	h, t := 0, 0
	for i := range k - 1 {
		for h < t && nums[deque[t-1]] < nums[i] {
			t--
		}
		deque[t] = i
		t++
	}
	m := n - k + 1
	ans := make([]int, m)
	for l, r := 0, k-1; r < n; l, r = l+1, r+1 {
		for h < t && nums[deque[t-1]] < nums[r] {
			t--
		}
		deque[t] = r
		t++
		for deque[h] < l {
			h++
		}
		ans[l] = nums[deque[h]]
	}
	return ans
}

func main() {
	result := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(result)
}
