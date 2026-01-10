package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 将小于0的数标记
	for i := range n {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 用哈希表 但是是以正负来代替bool类型的哈希表
	// 有可能一个数字会出现多次 所以以abs设置该数值的符号最保险
	for i := range n {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	result := firstMissingPositive([]int{1, 2, 0})
	fmt.Println(result)
}
