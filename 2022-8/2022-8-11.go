package main

func minimumReplacement(nums []int) int64 {
	//从最后面那个得到 每次最后面不需要再拆解
	if len(nums) == 1 {
		return 0
	}
	cur := nums[len(nums)-1]
	ans := int64(0)
	for i := len(nums) - 2; i >= 0; i-- {

		cnt := nums[i]/cur + 1
		if nums[i]%cur == 0 {
			cnt -= 1
		}
		ans += int64(cnt - 1)
		if cnt != 0 {
			cur = nums[i] / cnt
		}
		//fmt.Println(cur)
	}
	return ans
}

func longestIdealString(s string, k int) int {
	//记录下每个字母结束的最大值
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		for cur, cnt := byte(s[i]), 0; cur >= 'a' && cnt <= k; cur, cnt = byte(cur-1), cnt+1 {
			m[s[i]] = max(m[s[i]], m[cur]+1)
		}
		for cur, cnt := byte(s[i]+1), 0; cur <= 'z' && cnt < k; cur, cnt = byte(cur+1), cnt+1 {
			m[s[i]] = max(m[s[i]], m[cur]+1)
		}
	}
	ans := 0
	for _, val := range m {
		ans = max(ans, val)
	}
	return ans
}

func findBall(grid [][]int) []int {
	//实际上就是看确定的轨道
	n := len(grid[0])
	m := len(grid)
	to := make([][]int, m)
	for i := range to {
		to[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && j+1 < n && grid[i][j+1] == 1 {
				to[i][j] = 1
			} else if grid[i][j] == -1 && j > 0 && grid[i][j-1] == -1 {
				to[i][j] = -1
			} else {
				to[i][j] = 0
			}
		}
	}
	//fmt.Println(to)
	ans := make([]int, n)
	for j := 0; j < n; j++ {
		left := j
		for i := 0; i < m; i++ {
			if to[i][left] == 0 {
				ans[j] = -1
				break
			}
			left += to[i][left]
		}
		if ans[j] != -1 {
			ans[j] = left
		}

	}
	return ans
}
