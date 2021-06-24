// https://atcoder.jp/contests/joi2007ho/tasks/joi2007ho_c
// https://www.ioi-jp.org/joi/2006/2007-ho-prob_and_sol/2007-ho.pdf#page=5
// AC :)

package main

import (
	"fmt"
	"os"
	"sort"
)

type piller struct {
	x int
	y int
}

func main() {
	input := os.Stdin

	var n int
	var a [5001][5001]bool

	fmt.Fscan(input, &n)
	p := make([]piller, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &p[i].x, &p[i].y)
		a[p[i].x][p[i].y] = true
	}

	sort.Slice(p, func(i, j int) bool {
		if p[i].x < p[j].x {
			return true
		} else if p[i].x > p[j].x {
			return false
		}
		if p[i].y < p[j].y {
			return true
		}
		return false

	})

	maxarea := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if p[i].x < p[j].x {
				maxarea = max(maxarea, findarea(&a, p[i], p[j]))
			}
		}
	}

	outfile := os.Stdout
	fmt.Fprintln(outfile, maxarea)
}

func findarea(a *[5001][5001]bool, p1, p2 piller) int {
	dx := p2.x - p1.x
	dy := p2.y - p1.y

	if isInside(p1.x+dx-dy, p1.y+dx+dy) && isInside(p1.x-dy, p1.y+dx) {
		if (*a)[p1.x+dx-dy][p1.y+dx+dy] && (*a)[p1.x-dy][p1.y+dx] {
			return dx*dx + dy*dy
		}
	}
	return 0
}

func isInside(x, y int) bool {
	return (x >= 0 && x <= 5000 && y >= 0 && y <= 5000)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
