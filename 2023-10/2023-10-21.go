package main

func countPairs(n int, edges [][]int) int64 {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for _, v := range edges {
		merge(v[0], v[1], parent)
	}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[find(i, parent)]++
	}
	res := 0
	for i := 0; i < n; i++ {
		res += (n - m[i]) * m[i]
	}
	return int64(res / 2)
}
func find(a int, parent []int) int {
	if a == parent[a] {
		return a
	}
	parent[a] = find(parent[a], parent)
	return parent[a]
}
func merge(a, b int, parent []int) {
	a = find(a, parent)
	b = find(b, parent)
	parent[a] = b
}

func minSizeSubarray(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	c := target / sum
	part := target % sum
	m := make(map[int]int)
	m[0] = -1
	flag := false
	minCnt := len(nums)
	if part == 0 {
		flag = true
		minCnt = 0
	}
	prefix := 0
	for i := 0; i < len(nums)*2; i++ {
		prefix += nums[i%len(nums)]
		if m[prefix-part] != 0 {
			flag = true
			minCnt = min(minCnt, i-m[prefix-part])
		}
		m[prefix] = i
	}
	if flag {
		return c*len(nums) + minCnt
	}
	return -1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func differenceOfSums(n int, m int) int {
	cntm := n / m
	end := cntm * m
	return n*(n+1)/2 - (m+end)*cntm
}

func minOperations2(s1 string, s2 string, x int) int {
	diff := make([]int, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}
	if len(diff)%2 != 0 {
		return -1
	}
	dp := make([]int, len(diff)/2+1)
	for i := 1; i < len(diff)/2+1; i++ {
		dp[i] = 1e5
		for j := 0; j < i; j++ {
			dp[i] = min(dp[i], dp[j]+min(cnt1(diff, x, j, i), cnt2(diff, x, j, i)))
		}
	}
	return dp[len(diff)/2]
}
func cnt1(diff []int, x, start, end int) int {
	res := 0
	for i := 2 * start; i < 2*end; i += 2 {
		res += min(diff[i+1]-diff[i], x)
	}
	return res
}
func cnt2(diff []int, x, start, end int) int {
	res := 0
	for i := 2*start + 1; i < 2*end-1; i += 2 {
		res += min(diff[i+1]-diff[i], x)
	}
	res += min(diff[2*end-1]-diff[2*start], x)
	return res
}
