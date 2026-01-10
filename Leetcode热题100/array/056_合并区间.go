package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	ans := make([][]int, 0)
	n := len(intervals)
	vist := make([]bool, n)
	for i, interval := range intervals {
		// 当前区间被合并过
		if vist[i] {
			continue
		}
		// 当前没有被合并过，往后合并
		j := i + 1
		start, end := interval[0], interval[1]
		for j < n && end >= intervals[j][0] {
			end = max(end, intervals[j][1])
			vist[j] = true
			j++
		}
		ans = append(ans, []int{start, end})
	}
	return ans
}

func main() {
	result := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println(result)
}
