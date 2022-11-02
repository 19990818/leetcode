package main

import "math"

func bestCoordinate(towers [][]int, radius int) []int {
	maxSum := 0
	res := make([]int, 2)
	for x := 0; x <= 50; x++ {
		for y := 0; y <= 50; y++ {
			temp := 0
			for j := 0; j < len(towers); j++ {
				dis := pow2(y-towers[j][1]) + pow2(x-towers[j][0])
				if dis <= pow2(radius) {
					temp += int(float64(towers[j][2]) / (math.Sqrt(float64(dis)) + 1))
				}
			}
			if temp > maxSum {
				maxSum = temp
				res[0], res[1] = x, y
			}
		}

	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func pow2(a int) int {
	return a * a
}
