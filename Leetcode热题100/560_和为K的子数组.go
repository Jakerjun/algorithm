package main

import "fmt"

func subarraySum(nums []int, k int) int {
	/*
		公式推演
		1. 先对数组求前缀和
		pre[i] = pre[i-1] + nums[i]
		2. 遍历每个下标值 可以知道当前的 k 其实是 pre[j] - pre[i] == k, i < j
		pre[j] - pre[i] = k    (1)
		3. 但是我们目前遍历下标的时候 只能快速知道pre[j]的值 有多少个 pre[i] 我们是不确定的
		4. 所以我们需要构造一个求pre[i]的式子 也就是对 (1) 式进行变形 : pre[i] = pre[j] - k (2)
		5. 通过变形的式子我们可以快速知道pre[i]是否存在，但是此时遍历我们其实还需要求pre[i]的个数
		6. 所以通过hash表存储每一个遍历过的pre的值
	*/
	visit := make(map[int]int)
	visit[0] = 1
	ans := 0
	pre := 0
	for _, val := range nums {
		pre += val
		ans += visit[pre-k]
		visit[pre]++
	}
	return ans
}

func main() {
	result := subarraySum([]int{1, 1, 1}, 2)
	fmt.Println(result)
}
