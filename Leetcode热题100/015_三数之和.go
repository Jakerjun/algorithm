package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	// 排序，方便双指针
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	for i := range n {
		// 如果当前的i的值和之前用过的 i 相同那么没必要重复使用
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		j, k := i+1, n-1
		for j < k {
			sum := nums[j] + nums[k]
			if sum < target {
				j++
			} else if sum > target {
				k--
			} else {
				// 当前是一个答案点
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return ans
}

func main() {
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)
}
