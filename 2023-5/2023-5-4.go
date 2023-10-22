package main

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	dppre := make([]int, k+1)
	dpsuf := make([]int, k+1)
	m := make(map[int]int)
	for _, f := range fruits {
		m[f[0]] = f[1]
	}
	dppre[0] = m[startPos]
	for i := 1; i <= k; i++ {
		dpsuf[i] = dpsuf[i-1] + m[startPos+i]
		dppre[i] = dppre[i-1] + m[startPos-i]
	}
	res := 0
	for i := 0; i <= k; i++ {
		res = max(res, dpsuf[i]+dppre[max(k-2*i, 0)])
		res = max(res, dppre[i]+dpsuf[max(k-2*i, 0)])
	}
	return res
}
