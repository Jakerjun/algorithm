package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := map[int]struct{}{}
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	longestStreak := 0
	for num := range numSet {
		// 只从开头节点往后遍历
		if _, ok := numSet[num-1]; !ok {
			currentNum := num
			currentStreak := 1
			for {
				if _, ok := numSet[currentNum+1]; !ok {
					break
				}
				currentNum++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}

		}
	}
	return longestStreak

}

func main() {
	result := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Println(result)
}
