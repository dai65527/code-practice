package main

import "fmt"

var h, w int

func rec(i, j int, a [][]int, hash map[int]bool) int {
	if hash[a[i][j]] {
		return 0
	}
	if i == h-1 && j == w-1 {
		return 1
	}
	hash[a[i][j]] = true
	routes := 0
	if i < h-1 {
		routes += rec(i+1, j, a, hash)
	}
	if j < w-1 {
		routes += rec(i, j+1, a, hash)
	}
	delete(hash, a[i][j])
	return routes
}

func main() {
	fmt.Scanf("%d %d", &h, &w)
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	fmt.Println(rec(0, 0, a, map[int]bool{}))
}
