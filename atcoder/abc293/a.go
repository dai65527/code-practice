package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	ans := []byte(s)
	for i := 1; i < len(ans); i += 2 {
		ans[i], ans[i-1] = ans[i-1], ans[i]
	}
	fmt.Println(string(ans))
}
