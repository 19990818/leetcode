package main

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	m := make(map[int][]int)
	for _, v := range edges {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}
	travel := make(map[int]int)
	q := []int{1}
	res := make(map[int]float64)
	travel[1] = 1
	res[1] = 1.0
	for i := 0; i < t && len(q) > 0; i++ {
		temp := q
		q = nil
		for _, v := range temp {
			cnt := 0
			for _, v2 := range m[v] {
				if travel[v2] == 0 {
					cnt++
					q = append(q, v2)
				}
			}
			for _, v2 := range m[v] {
				if travel[v2] == 0 {
					travel[v2] = 1
					res[v2] = res[v] / float64(cnt)
				}
			}
			if cnt != 0 {
				delete(res, v)
			}
		}

	}
	return res[target]
}
