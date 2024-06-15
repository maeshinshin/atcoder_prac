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

type cheese struct {
	Deliciousness int64
	Amount        int64
}

func main() {
	var n, w int64
	fmt.Scanf("%d %d", &n, &w)
	sc.Split(bufio.ScanWords)

	var c []cheese
	for i := int64(0); i < n; i++ {
		c = append(c, cheese{scanInt64(), scanInt64()})
	}
	sort.Slice(c, func(i, j int) bool { return c[i].Deliciousness > c[j].Deliciousness })
	ans := int64(0)
	for i := int64(0); i < n; i++ {
		if w >= c[i].Amount {
			ans += c[i].Deliciousness * c[i].Amount
			w -= c[i].Amount
		} else {
			ans += w * c[i].Deliciousness
			break
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
