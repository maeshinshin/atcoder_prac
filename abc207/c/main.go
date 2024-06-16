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

type Interval struct {
	from float64
	to   float64
}

func main() {
	var n, ans int64
	fmt.Scanf("%d", &n)
	sc.Split(bufio.ScanWords)

	var v []Interval
	for i := int64(0); i < n; i++ {
		if t := scanInt64(); t == 1 {
			v = append(v, Interval{float64(scanInt64()), float64(scanInt64())})
		} else if t == 2 {
			v = append(v, Interval{float64(scanInt64()), float64(scanInt64()) - 0.1})
		} else if t == 3 {
			v = append(v, Interval{float64(scanInt64()) + 0.1, float64(scanInt64())})
		} else if t == 4 {
			v = append(v, Interval{float64(scanInt64()) + 0.1, float64(scanInt64()) - 0.1})
		}
	}
	for i := int64(0); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (v[i].from <= v[j].from && v[j].from <= v[i].to) || v[j].from <= v[i].from && v[i].from <= v[j].to {
				ans += 1
			}
		}
	}
	fmt.Println(ans)
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

func next_permutation(x sort.Interface) bool {
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
