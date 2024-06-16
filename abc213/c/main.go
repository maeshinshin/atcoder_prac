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

type Pos struct {
	x int64
	y int64
}

func main() {
	var h, w, n int64
	fmt.Scanf("%d %d %d", &h, &w, &n)
	sc.Split(bufio.ScanWords)

	var s []Pos
	a := make(map[int64]int64)
	b := make(map[int64]int64)
	for i := int64(0); i < n; i++ {
		s = append(s, Pos{scanInt64(), scanInt64()})
		a[s[i].x] = -1
		b[s[i].y] = -1
	}
	var as, bs []int64
	for k := range a {
		as = append(as, k)
	}
	for k := range b {
		bs = append(bs, k)
	}
	sort.Slice(as, func(i, j int) bool { return as[i] < as[j] })
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	for i, v := range as {
		a[v] = int64(i + 1)
	}
	for i, v := range bs {
		b[v] = int64(i + 1)
	}
	for i := int64(0); i < n; i++ {
		fmt.Println(a[s[i].x], b[s[i].y])
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

// slice
func sum(a []int64) (s int64) {
	for _, v := range a {
		s += v
	}
	return
}

// algorithm
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

func nextPermutation(x sort.Interface) bool {
	n := x.Len()
	if n <= 1 {
		return false
	}

	i := n - 2
	for i >= 0 && !x.Less(i, i+1) {
		i--
	}
	if i < 0 {
		reverse(x, 0, n-1)
		return false
	}

	j := n - 1
	for !x.Less(i, j) {
		j--
	}

	x.Swap(i, j)
	reverse(x, i+1, n-1)
	return true
}

func reverse(x sort.Interface, start, end int) {
	for start < end {
		x.Swap(start, end)
		start++
		end--
	}
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
