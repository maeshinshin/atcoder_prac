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

type Point struct {
	x int64
	y int64
	z int64
}

func getVertices(p1, p2 Point) []Point {
	return []Point{
		{p1.x, p1.y, p1.z},
		{p1.x, p1.y, p2.z},
		{p1.x, p2.y, p1.z},
		{p1.x, p2.y, p2.z},
		{p2.x, p1.y, p1.z},
		{p2.x, p1.y, p2.z},
		{p2.x, p2.y, p1.z},
		{p2.x, p2.y, p2.z},
	}
}

func main() {
	var v1 []Point
	var v2 []Point
	sc.Split(bufio.ScanWords)
	for i := int64(0); i < 2; i++ {
		v1 = append(v1, Point{scanInt64(), scanInt64(), scanInt64()})
	}
	for i := int64(0); i < 2; i++ {
		v2 = append(v2, Point{scanInt64(), scanInt64(), scanInt64()})
	}
	if v1[0].x > v1[1].x {
		tmp := v1[1].x
		v1[1].x = v1[0].x
		v1[0].x = tmp
	}
	if v1[0].y > v1[1].y {
		tmp := v1[1].y
		v1[1].y = v1[0].y
		v1[0].y = tmp
	}
	if v1[0].z > v1[1].z {
		tmp := v1[1].z
		v1[1].z = v1[0].z
		v1[0].z = tmp
	}
	if v2[0].x > v2[1].x {
		tmp := v2[1].x
		v2[1].x = v2[0].x
		v2[0].x = tmp
	}
	if v2[0].y > v2[1].y {
		tmp := v2[1].y
		v2[1].y = v2[0].y
		v2[0].y = tmp
	}
	if v2[0].z > v2[1].z {
		tmp := v2[1].z
		v2[1].z = v2[0].z
		v2[0].z = tmp
	}

	v11 := getVertices(v1[0], v1[1])
	v22 := getVertices(v1[0], v2[1])
	for _, p := range v11 {
		if (v2[0].x <= p.x) && (v2[0].y <= p.y) && (v2[0].z <= p.z) && (p.x <= v2[1].x) && (p.y <= v2[1].y) && (p.z <= v2[1].z) {
			count := make(map[bool]int)
			count[v2[0].x == p.x]++
			count[v2[0].y == p.y]++
			count[v2[0].z == p.z]++
			count[v2[1].x == p.x]++
			count[v2[1].y == p.y]++
			count[v2[1].z == p.z]++
			if count[true] == 0 {
				fmt.Println("Yes")
				return
			}
		}
	}
	for _, p := range v22 {
		if (v1[0].x <= p.x) && (v1[0].y <= p.y) && (v1[0].z <= p.z) && (p.x <= v1[1].x) && (p.y <= v1[1].y) && (p.z <= v1[1].z) {
			count := make(map[bool]int)
			count[v1[0].x == p.x]++
			count[v1[0].y == p.y]++
			count[v1[0].z == p.z]++
			count[v1[1].x == p.x]++
			count[v1[1].y == p.y]++
			count[v1[1].z == p.z]++
			if count[true] == 0 {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
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
