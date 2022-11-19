package main

func convertTemperature(celsius float64) []float64 {
	res := make([]float64, 2)
	res[0] = celsius + 273.15
	res[1] = celsius*1.80 + 32
	return res
}

func subarrayLCM(nums []int, k int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			res++
		}
		j := i + 1
		temp := nums[i]
		for j < len(nums) {
			//fmt.Println(temp,nums[j],gcd(temp,nums[j]))
			temp = (temp * nums[j]) / gcd(temp, nums[j])
			//fmt.Println(i,j,temp)
			if temp == k {
				res++
			}
			j++
		}
	}
	return res
}
func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func order(root *TreeNode) [][]int {
	res := make([][]int, 0)
	q := []*TreeNode{root}
	for {
		temp := make([]*TreeNode, 0)
		tempV := make([]int, 0)
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				temp = append(temp, cur.Left)
				tempV = append(tempV, cur.Left.Val)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
				tempV = append(tempV, cur.Right.Val)
			}
		}
		if len(temp) == 0 {
			break
		}
		res = append(res, tempV)
		q = temp
	}
	return res
}

func maxPalindromes(s string, k int) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = true
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}
	res := 0
	start := 0
	for j := k - 1; j < len(s); j++ {
		flag := 0
		for i := j - k + 1; i >= start; i-- {
			if dp[i][j] {
				res++
				flag = 1
				break
			}
		}
		if flag == 1 {
			start = j + 1
		}
	}
	return res
}
