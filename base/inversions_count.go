package main

import "fmt"

var n int
var q, temp []int

func merge_sort(q []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := l + (r-l)/2
	res := merge_sort(q, l, mid) + merge_sort(q, mid+1, r)
	i, j, k := l, mid+1, 0
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			temp[k] = q[i]
			k++
			i++
		} else {
			temp[k] = q[j]
			k++
			j++
			res += mid - i + 1
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
	return res
}

func main() {
	fmt.Scanf("%d", &n)
	q = make([]int, n)
	temp = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	res := merge_sort(q, 0, n-1)
	fmt.Println(res)
	return
}
