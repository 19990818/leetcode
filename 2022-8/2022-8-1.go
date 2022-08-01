package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minimumOperations(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		if val != 0 {
			m[val] = 1
		}
	}
	return len(m)
}

func maximumGroups(grades []int) int {
	n := len(grades)
	ans := 0
	cnt := 1
	for n >= cnt {
		ans++
		n -= cnt
		cnt++
	}
	return ans
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	//得到每个节点到各种
	var getDis func(node int) map[int]int
	getDis = func(node int) map[int]int {
		dis1 := make(map[int]int)
		next := edges[node]
		dis1[node] = 0
		dis := 1
		for next != -1 {
			if _, ok := dis1[next]; ok {
				break
			}
			dis1[next] = dis
			dis++
			next = edges[next]
		}
		return dis1
	}
	dis1, dis2 := getDis(node1), getDis(node2)
	ans := -1
	temp := math.MaxInt64
	for i := 0; i < len(edges); i++ {
		_, ok1 := dis1[i]
		_, ok2 := dis2[i]
		if ok1 && ok2 {
			if ans == -1 {
				ans = i
				temp = max(dis1[i], dis2[i])
			} else {
				if temp > max(dis1[i], dis2[i]) {
					ans = i
					temp = max(dis1[i], dis2[i])
				}
			}
		}
	}
	return ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
