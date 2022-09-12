package main

// #1551 使数组中所有元素相等的最小操作数
func minOperations2(n int) int {
	if n%2 == 1 {
		//说明为奇数
		return (n*n - 1) / 4
	}
	return n * n / 4
}
