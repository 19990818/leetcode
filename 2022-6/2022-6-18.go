package main

import "math/rand"

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		temp := new(Node)
		temp.Val = x
		temp.Next = temp
		return temp
	}
	cur := aNode
	flag := 0
	for cur.Next != aNode {
		//在cur处插入
		if x >= cur.Val && x <= cur.Next.Val || (cur.Next.Val < cur.Val && (x >= cur.Val || x <= cur.Next.Val)) {
			inSertSingleNode(cur, x)
			flag = 1
			break
		}
		cur = cur.Next
	}
	if flag == 0 {
		inSertSingleNode(cur, x)

	}
	return aNode
}
func inSertSingleNode(cur *Node, x int) {
	temp := new(Node)
	temp.Val = x
	tempNode := cur.Next
	cur.Next = temp
	temp.Next = tempNode
}

func isScramble(s1 string, s2 string) bool {
	//很容易得到递推公式
	// if len(s1) == 1 {
	// 	return s1 == s2
	// }
	// if s1 == s2 {
	// 	return true
	// }
	// for i := 0; i < len(s1)-1; i++ {
	// 	if isScramble(s1[0:i+1], s2[0:i+1]) && isScramble(s1[i+1:], s2[i+1:]) {
	// 		return true
	// 	}
	// 	if isScramble(s1[i+1:], s2[0:len(s2)-i-1]) && isScramble(s1[0:i+1], s2[len(s2)-i-1:]) {
	// 		return true
	// 	}
	// }
	// return false
	n := len(s1)
	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
		}
	}
	//初始化最小子问题
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = (s1[i] == s2[j])
		}
	}
	//采用长度定义更加直观，而且可以根据长度确定起始位置
	//首先长度不是作为最先维度，i,j范围判断错误，限定k值，实际上会造成情况不符合实际
	//并且速度会更慢

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for m := 2; m <= n; m++ {
				for k := 1; k < m; k++ {
					//找到其中一个正确的就可以说明当前状态可行
					if i <= n-m && j <= n-m {
						if dp[i][j][k] && dp[i+k][j+k][m-k] {
							dp[i][j][m] = true
							break
						}
						if dp[i+k][j][m-k] && dp[i][j+m-k][k] {
							dp[i][j][m] = true
							break
						}
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

func sortArray(nums []int) []int {
	var quickSort func(nums []int)
	quickSort = func(nums []int) {
		if len(nums) == 0 {
			return
		}
		left, right := 0, len(nums)-1
		randPos := rand.Intn(len(nums))
		nums[0], nums[randPos] = nums[randPos], nums[0]
		pior := nums[0]
		//fmt.Println(pior)
		for left < right {
			for left < right && nums[right] >= pior {
				right--
			}
			nums[left] = nums[right]
			for left < right && nums[left] <= pior {
				left++
			}
			nums[right] = nums[left]
		}
		nums[left] = pior
		//fmt.Println(nums)
		quickSort(nums[0:left])
		quickSort(nums[left+1:])
	}
	quickSort(nums)
	return nums
}
