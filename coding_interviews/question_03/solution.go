package main

import "fmt"

func Find(matrix []int, rows, columns, target int) bool {
	if matrix == nil || rows <= 0 || columns <= 0 {
		return false
	}
	row := 0
	column := columns - 1
	for row < rows && column >= 0 {
		index := row*columns + column
		if matrix[index] == target {
			return true
		} else if matrix[index] > target {
			column -= 1
		} else {
			row += 1
		}
	}

	return false
}

func main() {
	matrix := []int{
		1, 2, 8, 9,
		2, 4, 9, 12,
		4, 7, 10, 13,
		6, 8, 11, 15,
	}

	res := Find(matrix, 4, 4, 16)
	fmt.Println(res)
	str := "chenxingyu"
	for i, s := range str {
		fmt.Println(i, s)
	}
}
