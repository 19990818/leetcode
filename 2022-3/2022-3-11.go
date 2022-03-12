package main

func lexicalOrder(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := 0; i < n; i++ {
		ans[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			if num > n {
				num /= 10
			}
			num += 1
		}
	}
	return ans
}
