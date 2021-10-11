package main

import "fmt"

func quick_select(q []int, l, r, k int) int {
	if l >= r {
		return q[l]
	}
	x, i, j := q[l+(r-l)/2], l-1, r+1
	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	len := j - l + 1
	if len >= k {
		return quick_select(q, l, j, k)
	} else {
		return quick_select(q, j+1, r, k-len)
	}
}

func main() {
	var n, k int
	fmt.Scanf("%d%d", &n, &k)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	res := quick_select(q, 0, n-1, k)
	fmt.Println(res)
	return
}
