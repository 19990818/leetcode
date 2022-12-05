package main

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	n := len(toppingCosts)
	var dfs func(cur, idx int)
	res := baseCosts[0]
	dfs = func(cur, idx int) {
		res = getClosestNum(res, cur, target)
		if idx == n || cur > target {
			return
		}
		dfs(cur, idx+1)
		dfs(cur+toppingCosts[idx], idx+1)
		dfs(cur+toppingCosts[idx]*2, idx+1)
	}
	for _, v := range baseCosts {
		dfs(v, 0)
	}
	return res
}
func getClosestNum(a, b, target int) int {
	if abs(a-target) < abs(b-target) {
		return a
	}
	if abs(a-target) == abs(b-target) && a < b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
