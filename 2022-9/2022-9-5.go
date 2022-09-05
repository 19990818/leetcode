package main

func checkDistances(s string, distance []int) bool {
	actual := make([]int, 26)
	m := make(map[rune]int)
	for i, val := range s {
		if _, ok := m[val]; !ok {
			m[val] = i
		} else {
			actual[val-'a'] = i - m[val] - 1
		}
	}
	//fmt.Println(actual)
	for i := range distance {
		_, ok := m[rune(i+'a')]
		if distance[i] != actual[i] && ok {
			return false
		}
	}
	return true
}

func numberOfWays(startPos int, endPos int, k int) int {
	if endPos < startPos {
		endPos, startPos = startPos, endPos
	}
	dis := endPos - startPos
	if (dis+k)%2 == 1 {
		return 0
	}
	n := (k - dis) / 2
	mod := int(1e9 + 7)
	var fib func(x, y int) int
	dp := make([][]int, 1001)
	for i := range dp {
		dp[i] = make([]int, 1001)
	}
	fib = func(x, y int) int {
		if y == 0 || y == x {
			return 1
		}
		if dp[x-1][y-1] == 0 {
			dp[x-1][y-1] = fib(x-1, y-1)
		}
		if dp[x-1][y] == 0 {
			dp[x-1][y] = fib(x-1, y)
		}
		return (dp[x-1][y-1] + dp[x-1][y]) % mod
	}
	fib(k, n)
	return dp[k][n]
}

//定下右边指针后 左边指针移动
func longestNiceSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	ans := 1
	i, j := 0, 0
	sum := 0
	for ; j < len(nums); j++ {
		for (nums[j] & sum) > 0 {
			sum = sum ^ nums[i]
			//fmt.Println(sum,nums[i],i)
			i++
		}
		sum = sum ^ nums[j]
		//fmt.Println(i,j,sum)/*  */
		if ans < j-i+1 {
			ans = j - i + 1
		}
	}
	return ans
}
