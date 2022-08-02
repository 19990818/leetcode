package main

func longestCycle(edges []int) int {
	// 每个点只有一个出度
	// 遍历过的节点不再需要遍历 如果再次遍历会出现多次重复遍历环
	// 从而导致超时
	ans := -1
	m := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		if m[i] == 0 {
			m2 := make(map[int]int)
			cnt := 1
			next := i
			for next != -1 {
				m[next] = cnt
				m2[next] = cnt
				cnt++
				next = edges[next]
				if m2[next] != 0 {
					ans = max(ans, cnt-m2[next])
					break
				}
				if m[next] != 0 {
					break
				}
			}
		}
	}
	return ans
}
