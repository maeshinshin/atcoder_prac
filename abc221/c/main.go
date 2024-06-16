package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n []int64
	var ans int64
	ans = -math.MaxInt64
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		n = append(n, int64(s[i]-'0'))
	}

	sort.Slice(n, func(i, j int) bool { return n[i] > n[j] })
	for bi := int64(0); bi < 1<<uint(len(n)); bi++ {
		var a, b []int64
		for i := int64(0); i < int64(len(n)); i++ {
			if bi&(1<<i) == 0 {
				a = append(a, n[i])
			} else {
				b = append(b, n[i])
			}
		}
		if !((len(a) > 0) && (len(b) > 0)) {
			continue
		}
		var asum, bsum int64
		for i := 0; i < len(a); i++ {
			asum = asum*10 + a[i]
		}
		for i := 0; i < len(b); i++ {
			bsum = bsum*10 + b[i]
		}
		if ans < asum*bsum {
			ans = asum * bsum
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
