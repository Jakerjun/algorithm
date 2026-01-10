package main

import "fmt"

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	// 暂时记录一下第一行和第一列是否会被刷成0
	row0, col0 := false, false
	// 判断第一行
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	// 判断第一列
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}

	// 用标记变量法 如果某一行列上为0 那么就吧这一个行列开头的元素都设置为0
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 开始判断能不能刷成0
	if col0 {
		for i := range n {
			matrix[i][0] = 0
		}
	}
	if row0 {
		for j := range m {
			matrix[0][j] = 0
		}
	}

}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
