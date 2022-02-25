package main

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for x1 := 0; x1 < len(nums)-1; x1++ {
		for x2 := x1 + 1; x2 <= x1+k; x2++ {
			if abs(nums[x1]-nums[x2]) <= t {
				return true
			}
		}
	}
	return false
}

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	//通过hash实现 实际上是对区域范围进行判断 使用桶排序 桶可以用来表示一个范围
	//然后通过hash表达式可以唯一表示一个数 x=(t+1)*a+b(0<=b<=t)
	m := make(map[int]int)
	for i, val := range nums {
		id := getId(val, t+1)
		if _, ok := m[id]; ok {
			return true
		} else {
			m[id] = val
		}
		if _, ok := m[id-1]; ok && abs(m[id]-m[id-1]) <= t {
			return true
		}
		if _, ok := m[id+1]; ok && abs(m[id]-m[id+1]) <= t {
			return true
		}
		if i >= k {
			delete(m, getId(nums[i-k], t+1))
		}
	}
	return false
}

func getId(val, t int) int {
	if val >= 0 {
		return val / t
	}
	return (val+1)/t - 1
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' && i+ans < m && j+ans < n {
				ans = max(ans, getMaxSquare(matrix, i, j))
			}
		}
	}
	return ans * ans
}
func getMaxSquare(matrix [][]byte, i, j int) int {
	ans := 0
	//fmt.Println(i,j)
	for i+ans < len(matrix) && j+ans < len(matrix[0]) {
		for k := j; k <= j+ans; k++ {
			if matrix[i+ans][k] == '0' {
				return ans
			}
		}
		for k := i + ans; k >= i; k-- {
			if matrix[k][j+ans] == '0' {
				return ans
			}
		}
		ans += 1
	}
	return ans
}

func maximalSquare2(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min3(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}
func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func countWords(words1 []string, words2 []string) int {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	ans := 0
	for _, val := range words1 {
		m1[val]++
	}
	for _, val := range words2 {
		m2[val]++
	}
	for _, val := range words1 {
		if m1[val] == 1 && m2[val] == 1 {
			ans++
		}
	}
	return ans
}
