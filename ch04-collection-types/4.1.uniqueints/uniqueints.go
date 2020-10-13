package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <integer1> ...\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	nums := make([]int, 0, len(os.Args[1:]))
	for _, val := range os.Args[1:] {
		if num, err := strconv.Atoi(val); err == nil {
			nums = append(nums, num)
		} else {
			log.Fatal(fmt.Sprintf("'%s' is not a number", val))
		}
	}

	fmt.Println(UniqueInts(nums))
}

func UniqueInts(nums []int) []int {
	uinqueNums := make([]int, 0, len(nums))
	seen := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, found := seen[num]; !found {
			uinqueNums = append(uinqueNums, num)
			seen[num] = true
		}
	}
	return uinqueNums
}
