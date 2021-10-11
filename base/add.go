package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	var A, B []int
	for i := len(a) - 1; i >= 0; i -- {
		A = append(A, int(a[i] - '0'))
	}
	for i := len(b) - 1; i >= 0; i -- {
		B = append(B, int(b[i] - '0'))
	}
	res := add(A, B)
	for i := len(res) - 1; i >= 0; i -- {
		fmt.Printf("%d", res[i])
	}
	return
}

func add(a []int, b []int) []int {
	var c []int
	t := 0
	for i := 0; i < len(a) || i < len(b) || t > 0; i++ {
		if i < len(a) {
			t += a[i]
		}
		if i < len(b) {
			t += b[i]
		}
		c = append(c, t%10)
		t /= 10
	}
	return c
}