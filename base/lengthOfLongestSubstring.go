package main

func lengthOfLongestSubstring(s string) int {
	res := 0
	cnt := map[byte]int{}
	for l, r := 0, 0; l < len(s); r++ {
		for cnt[s[r]] > 1 {
			cnt[s[l]]--
			l++
		}
		res = Max(res, r-l+1)
	}
	return res
}

func Max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
