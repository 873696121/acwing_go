package main

import (
	"fmt"
	"strconv"
)

func cmp(a string, b string) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}

func sub(a string, b string) string {
	if !cmp(a, b) {
		return fmt.Sprintf("-%s", sub(b, a))
	}
	t := 0
	var A, B, c []int
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}
	for i := len(b) - 1; i >= 0; i-- {
		B = append(B, int(b[i]-'0'))
	}
	for i := 0; i < len(A); i++ {
		t += A[i]
		if i < len(B) {
			t -= B[i]
		}
		c = append(c, (t+10)%10)
		if t < 0 {
			t = -1
		} else {
			t = 0
		}
	}
	var res string
	for i := len(c); i >= 0; i -- {
		res += strconv.Itoa(c[i])
	}
	for len(c) > 1 && c[len(c) - 1:][0] == 0 {

	}
	return res
}

func main() {
	var a, b string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	res := sub(a, b)
	fmt.Println(res)
	return
}
