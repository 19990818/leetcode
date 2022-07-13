package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func ConstructorCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res = append(res, strconv.Itoa(root.Val))
	for {
		temp := make([]*TreeNode, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil {
				res = append(res, "null")
			} else {
				temp = append(temp, cur.Left)
				res = append(res, strconv.Itoa(cur.Left.Val))
			}
			if cur.Right == nil {
				res = append(res, "null")
			} else {
				temp = append(temp, cur.Right)
				res = append(res, strconv.Itoa(cur.Right.Val))
			}
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	dataArr := strings.Split(data, ",")
	root := new(TreeNode)
	root.Val, _ = strconv.Atoi(dataArr[0])
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	start := 1
	for start < len(dataArr) {
		tempA := make([]*TreeNode, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if dataArr[start] == "null" {
				cur.Left = nil
			} else {
				temp := new(TreeNode)
				cur.Left = temp
				cur.Left.Val, _ = strconv.Atoi(dataArr[start])
				tempA = append(tempA, cur.Left)
			}
			start++
			if dataArr[start] == "null" {
				cur.Right = nil
			} else {
				temp := new(TreeNode)
				cur.Right = temp
				cur.Right.Val, _ = strconv.Atoi(dataArr[start])
				tempA = append(tempA, cur.Right)
			}
			start++
		}
		if len(tempA) == 0 {
			break
		}
		queue = tempA
	}
	return root
}

func maxCoins(nums []int) int {
	//需要逆向思维 当其中一个气球击破时 可以看成是加入一个气球的逆向
	//将(i,j)作为击破气球的区间,取其中间的一个气球k进行击破 i<k<j
	//我们会将i,j中所有气球进行击破，假设气球k是最后一个被击破的气球
	//那么会存在一个很漂亮的递推关系式子
	//这个需要保证两端气球不被击破，这个我们可以进行预处理一下，在
	//初始数组两端加上1，方便计算的同时可以保证会将所有区间内的气球击破
	//注意因为所有的都为开区间，因此直接是dp[i][k]和dp[k][j]，这两个
	//并不包含k，击破k正好满足要求(i,j)区间
	//dp[i][j]=max(dp[i][j],dp[i][k]+dp[k][j]+nums[i]*nums[j]*nums[k])
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	//fmt.Println(nums)
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[j]*nums[k])

			}
		}
	}
	return dp[0][n-1]
}

func clumsy(n int) int {
	op := []string{"*", "/", "+", "-"}
	i := 0
	ans := n
	for n > 1 {
		switch op[i] {
		case "*":
			ans *= n - 1
		case "/":
			ans /= n - 1
		case "+":
			ans += n - 1
		case "-":
			temp := n - 1
			for cnt := 0; cnt < 2; cnt++ {
				n--
				if n-1 > 0 {
					i = (i + 1) % 4
					if op[i] == "*" {
						temp *= n - 1
					} else {
						temp /= n - 1
					}
				}
			}
			ans -= temp
		}
		i = (i + 1) % 4
		fmt.Println(ans)
		n--
	}
	return ans
}
