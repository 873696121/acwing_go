package main

import "fmt"

var n int
var q, temp []int

func merge_sort(q []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	merge_sort(q, l, mid)
	merge_sort(q, mid+1, r)
	i, j, k := l, mid+1, 0
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			temp[k] = q[i]
			i++
			k++
		} else {
			temp[k] = q[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = q[j]
		k++
		j++
	}
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = temp[j]
	}
}

func main() {
	fmt.Scanf("%d", &n)
	q = make([]int, n)
	temp = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	merge_sort(q, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
	return
}
