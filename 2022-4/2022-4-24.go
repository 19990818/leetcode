package main

import "sort"

func intersection(nums [][]int) []int {
	m := make(map[int]int)
	for _, val1 := range nums {
		for _, val2 := range val1 {
			m[val2]++
		}
	}
	ans := make([]int, 0)
	for key, val := range m {
		if val == len(nums) {
			ans = append(ans, key)
		}
	}
	sort.Ints(ans)
	return ans
}

func countLatticePoints(circles [][]int) int {
	var inCircle func(i, j, x, y, r int) bool
	inCircle = func(i, j, x, y, r int) bool {
		return (i-x)*(i-x)+(j-y)*(j-y) <= r*r
	}
	minx, miny := 200, 200
	maxx, maxy := 0, 0
	for _, val := range circles {
		minx = min(minx, val[0]-val[2])
		miny = min(miny, val[1]-val[2])
		maxx = max(maxx, val[0]+val[2])
		maxy = max(maxy, val[1]+val[2])
	}
	ans := 0
	for i := minx; i <= maxx; i++ {
		for j := miny; j <= maxy; j++ {
			for _, val := range circles {
				if inCircle(i, j, val[0], val[1], val[2]) {
					ans++
					break
				}
			}
		}
	}
	return ans
}

func countRectangles(rectangles [][]int, points [][]int) []int {
	ans := make([]int, 0)
	m := make(map[int][]int)
	for _, val := range rectangles {
		m[val[1]] = append(m[val[1]], val[0])
	}
	for key := range m {
		sort.Ints(m[key])
	}
	for _, val := range points {
		count := 0
		for i := 1; i <= 100; i++ {
			if i >= val[1] {
				count += len(m[i]) - sort.SearchInts(m[i], val[0])
			}
		}
		ans = append(ans, count)
	}
	return ans
}
