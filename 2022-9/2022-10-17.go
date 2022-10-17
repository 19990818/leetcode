package main

func totalFruit(fruits []int) int {
	//记录以每个位置开始的可能数目
	type pair struct {
		m   map[int]int
		cnt int
	}
	dp := make([][]pair, len(fruits))
	for i := range dp {
		dp[i] = make([]pair, 2)
	}
	dp[len(fruits)-1][0] = pair{map[int]int{fruits[len(fruits)-1]: 1}, 1}
	for i := len(fruits) - 2; i >= 0; i-- {
		if dp[i+1][0].m[fruits[i]] == 1 {
			dp[i][0].m = dp[i+1][0].m
			dp[i][0].cnt = dp[i+1][0].cnt + 1
		} else if dp[i+1][0].m[fruits[i]] == 0 {
			dp[i][1].m = dp[i+1][0].m
			dp[i][1].m[fruits[i]] = 1
			dp[i][1].cnt = dp[i+1][0].cnt + 1
		}
		if dp[i+1][1].m[fruits[i]] == 1 {
			dp[i][1].m = dp[i+1][1].m
			dp[i][1].cnt = dp[i+1][1].cnt + 1
		} else {
			dp[i][0].cnt = 1
			dp[i][0].m[fruits[i]] = 1
		}
	}
	ans := 2
	for _, val := range dp {
		ans = max(ans, val[0].cnt)
		ans = max(ans, val[1].cnt)
	}
	return ans
}

func totalFruit2(fruits []int) int {
	//记录以每个位置开始的可能数目
	m := make(map[int]int)
	i, j := 0, 0
	ans := 0
	for i <= j && j < len(fruits) {
		for {
			m[fruits[j]]++
			j++
			if len(m) > 2 {
				break
			}
			if j >= len(fruits) {
				j++
				break
			}
		}
		//fmt.Println(i,j,m)
		ans = max(ans, j-i-1)
		for len(m) > 2 {
			m[fruits[i]]--
			if m[fruits[i]] == 0 {
				delete(m, fruits[i])
			}
			i++
		}
	}
	return ans
}
