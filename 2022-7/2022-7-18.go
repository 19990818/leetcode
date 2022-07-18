package main

import "sort"

func containVirus(isInfected [][]int) int {
	// bfs中间增加一些细节处理 进行简化
	// dic表示扩散方向
	ans := 0
	dic := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	// 还需要进行记录旁边被扩散的位置，即0的位置
	// 此题是需要进行确定建立防火墙的数量，这个防火墙的边界
	// 实际上就是10之间的边界。可以将这个值进行一个记录处理
	// 因为只有当两个位置完全相同时候，他们之间的防火墙才是一样的
	// 因此碰到10边界只要完成不重复遍历，防火墙数量不会重复计算
	type pair struct {
		x int
		y int
	}
	m, n := len(isInfected), len(isInfected[0])
	for {
		// 每轮需要重新计算
		neighbors := make([]map[pair]int, 0)
		// 表示每轮中需要隔离每个区域防火墙的数量
		firewalls := make([]int, 0)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				// 不是感染或者已经完成遍历 直接跳过
				if isInfected[i][j] != 1 {
					continue
				}
				q := make([]pair, 0)
				q = append(q, pair{i, j})
				// 进行bfs的点有一个标记表示当前区域
				idx := len(neighbors) + 1
				firewall := 0
				neighbor := make(map[pair]int)
				isInfected[i][j] = -idx
				for len(q) > 0 {
					cur := q[0]
					q = q[1:]
					for _, val := range dic {
						//满足边界条件
						if cur.x+val[0] >= 0 && cur.x+val[0] < m && cur.y+val[1] >= 0 && cur.y+val[1] < n {
							if isInfected[cur.x+val[0]][cur.y+val[1]] == 1 {
								q = append(q, pair{cur.x + val[0], cur.y + val[1]})
								isInfected[cur.x+val[0]][cur.y+val[1]] = -idx
							} else if isInfected[cur.x+val[0]][cur.y+val[1]] == 0 {
								firewall++
								neighbor[pair{cur.x + val[0], cur.y + val[1]}] = 1
							}
						}
					}
				}
				// 将当前区域需要的防火墙数量加入防火墙数组当中
				firewalls = append(firewalls, firewall)
				neighbors = append(neighbors, neighbor)
			}
		}
		if len(neighbors) == 0 {
			break
		}
		idx := 0
		//不是选择需要防火墙最多 而是选择扩散最多的 即邻居最多的
		// for k := 1; k < len(firewalls); k++ {
		// 	if firewalls[k] > firewalls[idx] {
		// 		idx = k
		// 	}
		// }
		for k := 1; k < len(neighbors); k++ {
			if len(neighbors[k]) > len(neighbors[idx]) {
				idx = k
			}
		}
		ans += firewalls[idx]
		if len(neighbors) == 1 {
			break
		}
		//将进行特殊处理的进行还原 被防火墙围起来的标记为2
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if isInfected[i][j] < 0 {
					//是被围起来的区域
					if isInfected[i][j] == -idx-1 {
						isInfected[i][j] = 2
					} else {
						isInfected[i][j] = 1
					}
				}
			}
		}
		// 同时需要进行扩散 选择的下标不进行扩散
		for i, val := range neighbors {
			//需要进行扩散
			if i != idx {
				for key := range val {
					isInfected[key.x][key.y] = 1
				}
			}
		}
	}
	return ans
}

func minOperations(nums []int, numsDivide []int) int {
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	temp := numsDivide[0]
	if len(numsDivide) > 1 {
		temp = gcd((numsDivide[0]), numsDivide[1])
	}
	for k := 2; k < len(numsDivide); k++ {
		temp = gcd(temp, numsDivide[k])
	}
	sort.Ints(nums)
	for i, val := range nums {
		if temp%val == 0 {
			return i
		}
	}
	return -1
}

func longestArithSeqLength(nums []int) int {
	// 等差数组需要两个才能确定，值比较小 直接确定差值大小
	// dp[nums[i]][d]表示以nums[i]结束 差值为d的状态
	dp := make([][]int, 501)
	for i := range dp {
		dp[i] = make([]int, 1001)
	}
	ans := 2
	for i := 0; i < len(nums); i++ {
		for d := -500; d <= 500; d++ {
			if nums[i]+d >= 0 && nums[i]+d <= 500 {
				dp[nums[i]][d+500] = dp[nums[i]+d][d+500] + 1
			} else {
				dp[nums[i]][d+500] = 1
			}
			ans = max(ans, dp[nums[i]][d+500])
		}
	}
	return ans
}
