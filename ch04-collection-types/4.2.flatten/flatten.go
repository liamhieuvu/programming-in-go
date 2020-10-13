package main

import "fmt"

func main() {
	num2d := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	fmt.Println(Flatten(num2d))
}

func Flatten(num2d [][]int) []int {
	flat := []int{}
	for _, nums := range num2d {
		for _, num := range nums {
			flat = append(flat, num)
		}
	}
	return flat
}
