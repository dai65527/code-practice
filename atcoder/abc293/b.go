package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	called := make([]bool, n)
	for i := 0; i < n; i++ {
		var ai int
		fmt.Scanf("%d", &ai)
		if !called[i] {
			called[ai-1] = true
		}
	}
	ans := make([]string, 0, n)
	for i := range called {
		if !called[i] {
			ans = append(ans, strconv.Itoa(i+1))
		}
	}
	fmt.Println(len(ans))
	fmt.Println(strings.Join(ans, " "))
}
