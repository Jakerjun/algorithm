package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	x, y := 0, 0
	n, m = len(matrix), len(matrix[0])
	ans := make([]int, n*m)
	idx := 0
	for {
		// 往右走
		for check(matrix, x, y) {
			ans[idx] = matrix[x][y]
			idx++
			matrix[x][y] = flag
			y++

		}
		// 归位
		x++
		y--
		// 往下走
		for check(matrix, x, y) {
			ans[idx] = matrix[x][y]
			idx++
			matrix[x][y] = flag
			x++
		}
		// 归位
		x--
		y--
		// 往左走
		for check(matrix, x, y) {
			ans[idx] = matrix[x][y]
			idx++
			matrix[x][y] = flag
			y--
		}
		// 归位
		x--
		y++
		// 往上走
		for check(matrix, x, y) {
			ans[idx] = matrix[x][y]
			idx++
			matrix[x][y] = flag
			x--
		}
		x++
		y++
		if !check(matrix, x, y) {
			break
		}
	}
	return ans
}

var (
	n, m int
)

const (
	flag = 0x3f3f3f3f
)

// 能走就true 不能就false
func check(matrix [][]int, i, j int) bool {
	if i < 0 || i >= n || j < 0 || j >= m {
		return false
	}
	return matrix[i][j] != flag
}

func main() {
	result := spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	fmt.Println(result)
}
