package main

import "fmt"

func main() {
	cols := 3
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(Make2D(cols, nums))
}

func Make2D(cols int, nums []int) [][]int {
	rows := len(nums) / cols
	if len(nums)%cols != 0 {
		rows++
	}
	num2d := make([][]int, rows)
	for i, num := range nums {
		r := i / cols
		c := i % cols
		if num2d[r] == nil {
			num2d[r] = make([]int, cols)
		}
		num2d[r][c] = num
	}
	return num2d
}
