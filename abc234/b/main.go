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

type pos struct {
	x int
	y int
}

func main() {
	var n int64
	fmt.Scanf("%d", &n)
	sc.Split(bufio.ScanWords)

	var v []pos
	for i := int64(0); i < n; i++ {
		v = append(v, pos{int(scanInt64()), int(scanInt64())})
	}

	max := -math.MaxFloat64
	for i := int64(0); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			tmp := math.Sqrt(float64((v[i].x-v[j].x)*(v[i].x-v[j].x) + (v[i].y-v[j].y)*(v[i].y-v[j].y)))
			if max < tmp {
				max = tmp
			}
		}
	}
	fmt.Println(max)
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

func scanString() (s string) {
	sc.Scan()
	s = sc.Text()
	return
}
