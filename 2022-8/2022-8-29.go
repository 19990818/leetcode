package main

func removeStones(stones [][]int) int {
	//我们将到达每个点 然后根据每个点进行扩散
	xm, ym := make(map[int][]int), make(map[int][]int)
	for i, val := range stones {
		xm[val[0]] = append(xm[val[0]], i)
		ym[val[1]] = append(ym[val[1]], i)
	}
	m := make(map[int]int)
	var bfs func(now int)
	bfs = func(now int) {
		queue := []int{now}
		m[now] = 1
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			x, y := stones[cur][0], stones[cur][1]
			m[cur] = 1
			for _, subx := range xm[x] {
				if m[subx] == 0 {
					m[subx] = 1
					queue = append(queue, subx)
				}
			}
			for _, suby := range ym[y] {
				if m[suby] == 0 {
					m[suby] = 1
					queue = append(queue, suby)
				}
			}
		}
	}
	ans := 0
	for i := 0; i < len(stones); i++ {
		if m[i] == 0 {
			ans++
			bfs(i)
		}
	}
	return len(stones) - ans
}

func stoneGame(piles []int) bool {
	// return true
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(-dp[i-1][j]+piles[i], -dp[i][j-1]+piles[j])
		}
	}
	return dp[0][n-1] > 0
}
