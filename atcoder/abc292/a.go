package main

import "fmt"

func main() {
	var s string
	offset := 'A' - 'a'
	fmt.Scanf("%s", &s)
	ans := make([]byte, 0, len(s))
	for _, b := range []byte(s) {
		ans = append(ans, b+byte(offset))
	}
	fmt.Println(string(ans))
}
