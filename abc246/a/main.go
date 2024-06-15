package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	x := make(map[int]int)
	y := make(map[int]int)
	for i := int64(0); i < 3; i++ {
		var tx, ty int
		fmt.Scanf("%d %d", &tx, &ty)
		x[tx] += 1
		y[ty] += 1
	}
	for k, v := range x {
		if v == 1 {
			fmt.Print(k)
		}
	}
	fmt.Print(" ")
	for k, v := range y {
		if v == 1 {
			fmt.Print(k)
		}
	}
	fmt.Println()
}

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

// scan
func scanInt64() (i int64) {
	var err error
	sc.Scan()
	i, err = strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		log.Fatal("Failed to input", err)
	}
	return
}
