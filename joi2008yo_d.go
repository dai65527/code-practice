// https://atcoder.jp/contests/joi2008yo/tasks/joi2008yo_d

package main

import "fmt"

type star struct {
	x int
	y int
}

func findstar(x int, y int, stars []star) bool {
	for _, s := range stars {
		if x == s.x && y == s.y {
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	pic := make([]star, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&pic[i].x, &pic[i].y)
	}

	var m int
	fmt.Scan(&m)
	stars := make([]star, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&stars[i].x, &stars[i].y)
	}

	for _, s := range stars {
		for _, p := range pic {
			dx := s.x - p.x
			dy := s.y - p.y
			found := true
			for _, pp := range pic {
				if !findstar(pp.x+dx, pp.y+dy, stars) {
					found = false
					break
				}
			}
			if found {
				fmt.Println(dx, dy)
				return
			}
		}
	}
}
