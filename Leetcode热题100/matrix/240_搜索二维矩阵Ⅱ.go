package main

import "fmt"

// plan1 暴力
func searchMatrix1(matrix [][]int, target int) bool {
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// n*log(m)
func searchMatrix2(matrix [][]int, target int) bool {
	m := len(matrix[0])
	for row := range matrix {
		start, end := 0, m-1
		for start <= end {
			mid := (start + end) >> 1
			if target < matrix[row][mid] {
				end = mid - 1
			} else if target > matrix[row][mid] {
				start = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

// O(n)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 5))
}
