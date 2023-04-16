package main

func findSolution(customFunction func(int, int) int, z int) [][]int {
	ans := make([][]int, 0)
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 1000; j++ {
			if customFunction(i, j) == z {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
