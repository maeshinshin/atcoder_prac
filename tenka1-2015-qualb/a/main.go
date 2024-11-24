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

func calc(n int) int {
	if n == 0 {
		return 100
	} else if n < 0 {
		return 0
	}
	return calc(n-1) + calc(n-2) + calc(n-3)
}
func main() {
	fmt.Println(calc(19))

}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func abs(a int64) int64 {
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

// heap queue
type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
