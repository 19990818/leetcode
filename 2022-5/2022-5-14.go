package main

import (
	"sort"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	ans := 0
	numStr := strconv.Itoa(num)
	for i := 0; i <= len(numStr)-k; i++ {
		temp, _ := strconv.Atoi(numStr[i : i+k])
		//fmt.Println(temp)
		if temp != 0 && num%temp == 0 {
			//fmt.Println(temp)
			ans++
		}
	}
	return ans
}

func waysToSplitArray(nums []int) int {
	ans := 0
	sum := 0
	for _, val := range nums {
		sum += val
	}
	left := 0
	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		right := sum - left
		if left >= right {
			ans++
		}
	}
	return ans
}

type tilesSort [][]int

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Sort(tilesSort(tiles))
	ans := 0
	// if carpetLen==1{
	//     return 1
	// }
	//fmt.Println(tiles)
	pos := make([]int, 0)
	for _, val := range tiles {
		pos = append(pos, val[0])
	}
	sum := make([]int, len(tiles))
	for idx, val := range tiles {
		if idx == 0 {
			sum[idx] = val[1] - val[0] + 1
		} else {
			sum[idx] = sum[idx-1] + (val[1] - val[0] + 1)
		}
	}
	//fmt.Println("sum",sum)
	for idx, val := range tiles {
		end := val[0] + carpetLen - 1
		x := sort.SearchInts(pos, end)
		if x < len(pos) && pos[x] == end {
			x = x + 1
		}
		//fmt.Println("x",x)
		var temp, sumN int
		if idx == 0 {
			temp = 0
		} else {
			temp = sum[idx-1]
		}
		if x >= 2 {
			sumN = sum[x-2]
		} else {
			sumN = 0
		}
		// fmt.Println(sumN,temp,end,tiles[x-1][1],tiles[x-1][0])
		ans = max(max(sumN-temp, 0)+min(end, tiles[x-1][1])-tiles[x-1][0]+1, ans)
	}
	return ans
}
func (m tilesSort) Len() int {
	return len(m)
}
func (m tilesSort) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}
func (m tilesSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
