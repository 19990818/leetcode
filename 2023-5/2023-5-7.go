package main

func numPairsDivisibleBy60(time []int) int {
	m := make(map[int]int)
	ans := 0
	for _, v := range time {
		if m[60-v%60] > 0 {
			ans++
			m[60-v%60]--
		} else {
			m[v%60]++
		}
	}
	return ans
}

func distinctDifferenceArray(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	res := make([]int, len(nums))
	m2 := make(map[int]int)
	for i, v := range nums {
		m2[v]++
		m[v]--
		if m[v] == 0 {
			delete(m, v)
		}
		res[i] = len(m2) - len(m)
	}
	return res
}

type FrequencyTracker struct {
	m   map[int]int
	cnt map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		m:   make(map[int]int),
		cnt: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	this.m[number]++
	this.cnt[this.m[number]]++
	this.cnt[this.m[number]-1]--
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.m[number] == 0 {
		return
	}
	this.m[number]--
	this.cnt[this.m[number]]++
	this.cnt[this.m[number]+1]--
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.cnt[frequency] > 0
}

func colorTheArray(n int, queries [][]int) []int {
	// 进行修改只会影响前后两个元素
	t := make([]int, n)
	res := make([]int, len(queries))
	t[queries[0][0]] = queries[0][1]
	for i := 1; i < len(queries); i++ {
		res[i] = res[i-1]
		idx := queries[i][0]
		v := queries[i][1]
		if idx+1 < n {
			if t[idx] == t[idx+1] && t[idx] != 0 {
				res[i]--
			}
			if v == t[idx+1] {
				res[i]++
			}
		}
		if idx > 0 {
			if t[idx] == t[idx-1] && t[idx] != 0 {
				res[i]--
			}
			if v == t[idx-1] {
				res[i]++
			}
		}
		t[idx] = v
	}
	return res
}

func minIncrements(n int, cost []int) int {
	var dfs func(i int) (int, int)
	dfs = func(i int) (int, int) {
		if 2*i >= n {
			return 0, cost[i-1]
		}
		ops1, sum1 := dfs(2 * i)
		ops2, sum2 := dfs(2*i + 1)
		return ops1 + ops2 + abs(sum1-sum2), max(sum1, sum2) + cost[i-1]
	}
	res, _ := dfs(1)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isWinner(player1 []int, player2 []int) int {
	cnt := func(player []int) int {
		s := 0
		for i, v := range player {
			if (i > 0 && player[i-1] == 10) || (i > 1 && player[i-2] == 10) {
				s += 2 * v
			} else {
				s += v
			}
		}
		return s
	}
	if cnt(player1) > cnt(player2) {
		return 1
	}
	if cnt(player1) < cnt(player2) {
		return 2
	}
	return 0
}
