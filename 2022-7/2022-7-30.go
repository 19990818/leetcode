package main

type unitFind struct {
	parent, rank []int
}

// 并查集的合并过程是不可避免的 实际上仍然是操作parent数组
// 但是通过merge可以很直观的进行一个展示
func newUnitFind(n int) unitFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return unitFind{parent: parent, rank: make([]int, n)}
}
func (this unitFind) find(x int) int {
	if x == this.parent[x] {
		return x
	}
	this.parent[x] = this.find(this.parent[x])
	return this.parent[x]
}

func (this unitFind) merge(a, b int) {
	a, b = this.find(a), this.find(b)
	if a == b {
		return
	}
	if this.rank[a] > this.rank[b] {
		this.parent[b] = a
	} else if this.rank[a] < this.rank[b] {
		this.parent[a] = b
	} else {
		this.rank[a]++
		this.parent[b] = a
	}
}

func largestComponentSize(nums []int) (ans int) {
	n := 0
	for _, val := range nums {
		n = max(val, n)
	}
	f := newUnitFind(n + 1)
	for _, val := range nums {
		for i := 2; i*i <= val; i++ {
			if val%i == 0 {
				f.merge(val, i)
				f.merge(val, val/i)
			}
		}
	}
	cnt := make([]int, n+1)
	for _, val := range nums {
		idx := f.find(val)
		cnt[idx]++
		ans = max(ans, cnt[idx])
	}
	return ans
}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	var dfs func(sum, cur int)
	mid := make([]int, 0)
	dfs = func(sum, cur int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := append([]int{}, mid...)
			ans = append(ans, temp)
			return
		}
		for i := cur; i < len(candidates); i++ {
			val := candidates[i]
			mid = append(mid, val)
			dfs(sum+val, i)
			mid = mid[0 : len(mid)-1]
		}
	}
	dfs(0, 0)
	return ans
}
