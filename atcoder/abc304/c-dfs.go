package main

import (
	"fmt"
)

type Node struct {
	x int
	y int
}

func dist2(p1, p2 Node) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return dx*dx + dy*dy
}

func dfs(current int, ans []bool, graph [][]bool) {
	for i := range graph[current] {
		if !ans[i] && graph[current][i] {
			ans[i] = true
			dfs(i, ans, graph)
		}
	}
}

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)
	nodes := make([]Node, n)
	for i := range nodes {
		fmt.Scanf("%d %d", &nodes[i].x, &nodes[i].y)
	}
	// 点i,jがつながっていれば(=距離がd以内)ならtrue
	graph := make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			graph[i][j] = d*d >= dist2(nodes[i], nodes[j])
		}
	}
	// 0のノードを起点にして辿れる場合は感染しているはず
	// 深さ優先探索する
	ans := make([]bool, n)
	ans[0] = true
	dfs(0, ans, graph)
	for _, a := range ans {
		if a {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
