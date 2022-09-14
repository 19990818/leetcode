package main

import "sort"

func loudAndRich(richer [][]int, quiet []int) []int {
	in := make(map[int]int)
	ins := make(map[int][]int)
	out := make(map[int][]int)
	for _, val := range richer {
		in[val[1]]++
		ins[val[1]] = append(ins[val[1]], val[0])
		out[val[0]] = append(out[val[0]], val[1])
	}
	n := len(quiet)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}
	ans := make([]int, n)
	for {
		temp := make([]int, 0)
		for len(queue) > 0 {
			cur := queue[0]
			ans[cur] = cur
			for _, i := range ins[cur] {
				if quiet[ans[i]] < quiet[ans[cur]] {
					ans[cur] = ans[i]
				}
			}
			queue = queue[1:]
			for _, val := range out[cur] {
				in[val]--
				if in[val] == 0 {
					temp = append(temp, val)
				}
			}
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
	}
	return ans
}

func carFleet(target int, position []int, speed []int) int {
	type ps struct {
		pos int
		v   int
	}
	pss := make([]ps, len(speed))
	for i := 0; i < len(position); i++ {
		pss[i] = ps{position[i], speed[i]}
	}
	sort.Slice(pss, func(i, j int) bool {
		return pss[i].pos < pss[j].pos
	})
	ans := 1
	cur := len(pss) - 1
	for i := len(pss) - 2; i >= 0; i-- {
		//追不上
		if (target-pss[cur].pos)*pss[i].v < (target-pss[i].pos)*pss[cur].v {
			ans++
			cur = i
		}
	}
	return ans
}
