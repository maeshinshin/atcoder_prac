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
	var a, b, c int64
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if c%2 == 0 {
		if abs(a) < abs(b) {
			fmt.Println("<")
		} else if abs(a) == abs(b) {
			fmt.Println("=")
		} else {
			fmt.Println(">")
		}
	} else {
		if a < 0 && b < 0 {
			if a < b {
				fmt.Println(">")
			} else if a == b {
				fmt.Println("=")
			} else {
				fmt.Println("<")
			}
		} else {
			if a < b {
				fmt.Println("<")
			} else if a == b {
				fmt.Println("=")
			} else {
				fmt.Println(">")
			}
		}
	}
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
