package main

import "sort"

func minimumMoney(transactions [][]int) int64 {
	backMax := 0
	ans := int64(0)
	costMax := -1
	for _, val := range transactions {
		if val[0]-val[1] > 0 {
			ans += int64(val[0] - val[1])
			backMax = max(backMax, val[1])
		} else {
			costMax = max(costMax, val[0])
		}
	}
	return ans + int64(max(backMax, costMax))
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	ans := make([][]int, 0)
	curx, cury := rStart, cStart
	ans = append(ans, []int{rStart, cStart})
	cnt := 1
	for len(ans) != rows*cols {
		if cnt%2 == 1 {
			for j := cury + 1; j <= cury+cnt; j++ {
				if j < cols && j >= 0 && curx < rows && curx >= 0 {
					ans = append(ans, []int{curx, j})
				}
			}
			cury = cury + cnt
			for i := curx + 1; i <= curx+cnt; i++ {
				if i < rows && i >= 0 && cury < cols && cury >= 0 {
					ans = append(ans, []int{i, cury})
				}
			}
			curx = curx + cnt
		} else {
			for j := cury - 1; j >= cury-cnt; j-- {
				if j < cols && j >= 0 && curx < rows && curx >= 0 {
					ans = append(ans, []int{curx, j})
				}
			}
			cury = cury - cnt
			for i := curx - 1; i >= curx-cnt; i-- {
				if i < rows && i >= 0 && cury < cols && cury >= 0 {
					ans = append(ans, []int{i, cury})
				}
			}
			curx = curx - cnt
		}
		//fmt.Println(curx,cury,len(ans))
		cnt++
	}
	return ans
}

func possibleBipartition(n int, dislikes [][]int) bool {
	m1, m2 := make(map[int]int), make(map[int]int)
	sort.Slice(dislikes, func(i, j int) bool {
		return dislikes[i][0] < dislikes[j][0]
	})
	for _, dislike := range dislikes {
		if m1[dislike[0]] == 1 && m2[dislike[0]] == 1 {
			return false
		}
		if m1[dislike[1]] == 1 && m2[dislike[1]] == 1 {
			return false
		}
		if m1[dislike[0]] == 1 && m1[dislike[1]] == 1 {
			return false
		}
		if m2[dislike[1]] == 1 && m2[dislike[0]] == 1 {
			return false
		}
		if m1[dislike[0]] == 1 {
			m2[dislike[1]] = 1
		} else if m2[dislike[0]] == 1 {
			m1[dislike[1]] = 1
		} else if m1[dislike[1]] == 1 {
			m2[dislike[0]] = 1
		} else if m2[dislike[1]] == 1 {
			m1[dislike[0]] = 1
		} else {
			m1[dislike[0]] = 1
			m2[dislike[1]] = 1
		}
	}
	return true
}
