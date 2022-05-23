package main

import (
	"sort"
)

func partitionLabels(s string) []int {
	letter := make([][]int, 26)
	for i := 0; i < 26; i++ {
		letter[i] = make([]int, 2)
		letter[i][0] = -1
		letter[i][1] = -1
	}
	for index, val := range s {
		letter[val-'a'][1] = index
	}
	for i := len(s) - 1; i >= 0; i-- {
		letter[s[i]-'a'][0] = i
	}
	temp := make([][]int, 0)
	for _, val := range letter {
		if val[0] != -1 && val[1] != -1 {
			temp = append(temp, val)
		}
	}
	letter = temp
	sort.Sort(towDemion(letter))
	// fmt.Println(letter)
	cur := make([][]int, 0)
	cur = append(cur, letter[0])
	for i := 1; i < len(letter); i++ {
		if letter[i][0] < cur[len(cur)-1][1] {
			cur[len(cur)-1][1] = max(letter[i][1], cur[len(cur)-1][1])
		} else {
			cur = append(cur, letter[i])
		}
	}
	//fmt.Println(cur)
	ans := make([]int, 0)
	for _, val := range cur {
		ans = append(ans, val[1]-val[0]+1)
	}
	return ans
}

type towDemion [][]int

func (m towDemion) Len() int {
	return len(m)
}

func (m towDemion) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func (m towDemion) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	//最大为中心位置
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}
	for _, val := range mines {
		grid[val[0]][val[1]] = 1
	}
	var travel func(x, y int) int
	travel = func(x, y int) int {
		if grid[x][y] == 1 {
			return 0
		}
		i := 1
		for x-i >= 0 && x+i < n && y-i >= 0 && y+i < n {
			if grid[x-i][y] == 0 && grid[x+i][y] == 0 && grid[x][y-i] == 0 && grid[x][y+i] == 0 {
				i++
			} else {
				break
			}
		}
		return i
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				ans = max(ans, travel(i, j))
			}
		}
	}
	return ans
}

func reorganizeString(s string) string {
	letter := make([]int, 26)
	for _, val := range s {
		letter[val-'a']++
	}
	maxCount := 0
	var maxLetter byte
	for i := 0; i < 26; i++ {
		if maxCount < letter[i] {
			maxCount = letter[i]
			maxLetter = byte(i + 'a')
		}
	}
	parts := make([][]byte, maxCount)
	for i := 0; i < maxCount; i++ {
		parts[i] = append(parts[i], maxLetter)
	}
	cur := 0
	for i := 0; i < 26; i++ {
		if byte(i+'a') != maxLetter {
			for letter[i] > 0 {
				parts[cur] = append(parts[cur], byte(i+'a'))
				cur = (cur + 1) % maxCount
				letter[i]--
			}
		}
	}
	ans := make([]byte, 0)
	for index, val := range parts {
		if len(val) == 1 && index != len(parts)-1 {
			return ""
		}
		ans = append(ans, val...)
	}
	return string(ans)
}
