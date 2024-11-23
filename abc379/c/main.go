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

type pair struct {
	x, a int
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	sc.Split(bufio.ScanWords)

	var xa []pair
	for i := 0; i < m; i++ {
		xa = append(xa, pair{scanInt(), 0})
	}
	for i := 0; i < m; i++ {
		xa[i].a = scanInt()
	}
	sort.Slice(xa, func(i, j int) bool { return xa[i].x < xa[j].x })
	{
		tmp := 0
		for _, v := range xa {
			if tmp+1 < v.x {
				fmt.Println(-1)
				return
			}
			tmp += v.a
		}
	}
	sort.Slice(xa, func(i, j int) bool { return xa[i].x > xa[j].x })
	ans := 0
	{
		top := n + 1
		for _, v := range xa {
			if v.x < top-v.a {
				ans += v.a*(v.a+1)/2 + v.a*(top-v.a-v.x-1)
				top = top - v.a
			} else if v.x > top-v.a {
				fmt.Println(-1)
				return
			} else {
				ans += (top - v.x - 1) * (top - v.x) / 2
				top = v.x
			}
		}
	}
	fmt.Println(ans)
	return

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
func scanInt() int {
	var err error
	sc.Scan()
	i, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		log.Fatal("Failed to input", err)
	}
	return int(i)
}

func scanString() (s string) {
	sc.Scan()
	s = sc.Text()
	return
}
