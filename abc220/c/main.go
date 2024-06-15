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
	var n, x int64
	fmt.Scanf("%d", &n)
	sc.Split(bufio.ScanWords)

	var s []int64
	for i := int64(0); i < n; i++ {
		s = append(s, scanInt64())
	}
	x = scanInt64()

	sum := sum(s)
	ans := x / sum
	ans *= n
	r := x % sum
	for _, v := range s {
		ans++
		r -= v
		if r < 0 {
			fmt.Println(ans)
			return
		}
	}
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

func sum(a []int64) (s int64) {
	for _, v := range a {
		s += v
	}
	return
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

func scanString() (s string) {
	sc.Scan()
	s = sc.Text()
	return
}
