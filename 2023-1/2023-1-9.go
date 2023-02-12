package main

func reinitializePermutation(n int) int {
	start, cur, cnt := n-2, n-2, 0
	for cur != start || cnt == 0 {
		if cur%2 == 0 {
			cur = cur / 2
		} else {
			cur = n/2 + cur/2
		}
		cnt++
	}
	return cnt
}
