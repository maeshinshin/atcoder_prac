package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func main() {
	var x int
	fmt.Scan(&x)
	if x < 40 {
		fmt.Println(40 - x)
	} else if x < 70 {
		fmt.Println(70 - x)
	} else if x < 90 {
		fmt.Println(90 - x)
	} else {
		fmt.Println("expert")
	}
}
