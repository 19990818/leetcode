package main

func possibleBipartition(n int, dislikes [][]int) bool {
	//着色法 我们假设给某一个第一种颜色 那么dislikes以另一种颜色
	m := make(map[int][]int)
	for _, dislike := range dislikes {
		m[dislike[0]] = append(m[dislike[0]], dislike[1])
		m[dislike[1]] = append(m[dislike[1]], dislike[0])
	}
	travel := make(map[int]int)
	color := 1
	bfs := func(start int) bool {
		queue := []int{start}
		travel[start] = color
		for {
			temp := make([]int, 0)
			color = color ^ 3
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				for _, val := range m[cur] {
					if travel[val] != 0 && travel[val] != color {
						return false
					}
					if travel[val] == 0 {
						temp = append(temp, val)
						travel[val] = color
					}
				}
			}
			if len(temp) == 0 {
				break
			}
			queue = temp
		}
		return true
	}
	for i := 1; i <= n; i++ {
		if travel[i] == 0 && !bfs(i) {
			return false
		}
	}
	return true
}
