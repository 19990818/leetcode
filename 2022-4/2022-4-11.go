package main

import (
	"math"
	"sort"
	"strconv"
)

func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	var total int
	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j < len(wall[i]); j++ {
			m[sum+wall[i][j]]++
			sum += wall[i][j]
		}
		//fmt.Println(m)
		total = sum
	}
	ans := len(wall)
	for key, val := range m {
		if key == 0 || key == total {
			continue
		}
		ans = min(ans, len(wall)-val)
	}
	return ans
}

func nextGreaterElement(n int) int {
	nStr := strconv.Itoa(n)
	ans := []byte(nStr)
	stack := make([]int, 0)
	stack = append(stack, len(nStr)-1)
	flag := 0
	for i := len(nStr) - 2; i >= 0; i-- {
		swapR := i
		for len(stack) > 0 && nStr[i] < nStr[stack[len(stack)-1]] {
			swapR = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
		if swapR != i {
			ans[i], ans[swapR] = ans[swapR], ans[i]
			flag = 1
			temp := make([]int, 0)
			for j := i + 1; j < len(ans); j++ {
				temp = append(temp, int(ans[j]-'0'))
			}
			sort.Ints(temp)
			//fmt.Println(temp)
			for j := 0; j < len(temp); j++ {
				ans[i+1+j] = byte(temp[j] + '0')
			}
			break
		}
		if nStr[i] == nStr[stack[len(stack)-1]] {
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}
	if flag == 0 {
		return -1
	}
	digit, err := strconv.Atoi(string(ans))
	if err != nil {
		return -1
	}
	if digit > math.MaxInt32 {
		return -1
	}
	return digit
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	if quadTree1.IsLeaf && quadTree1.Val || quadTree2.IsLeaf && !quadTree2.Val {
		return quadTree1
	}
	if quadTree1.IsLeaf && !quadTree1.Val || quadTree2.IsLeaf && quadTree2.Val {
		return quadTree2
	}
	l1 := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	r1 := intersect(quadTree1.TopRight, quadTree2.TopRight)
	l2 := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	r2 := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if l1.IsLeaf && r1.IsLeaf && l2.IsLeaf && r2.IsLeaf && l1.Val == l2.Val && l2.Val == r1.Val && r1.Val == r2.Val {
		return &Node{l1.Val, true, nil, nil, nil, nil}
	}
	return &Node{true, false, l1, r1, l2, r2}
}
