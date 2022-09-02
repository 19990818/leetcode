package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numSpecialEquivGroups(words []string) int {
	m := make(map[string]int)
	ans := 0
	for _, val := range words {
		odd, even := []byte{}, []byte{}
		for i := range val {
			if i%2 == 0 {
				odd = append(odd, val[i])
			} else {
				even = append(even, val[i])
			}
		}
		sort.Slice(odd, func(i, j int) bool {
			return odd[i] < odd[j]
		})
		sort.Slice(even, func(i, j int) bool {
			return even[i] < even[j]
		})
		if m[string(odd)+string(even)] == 0 {
			m[string(odd)+string(even)] = 1
			ans++
		}
	}
	return ans
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{new(TreeNode)}
	}
	m := make(map[int][]*TreeNode)
	ans := make([]*TreeNode, 0)
	for i := 1; i < n-1; i++ {
		if m[i] == nil {
			m[i] = allPossibleFBT(i)
		}
		for _, left := range m[i] {
			if m[n-1-i] == nil {
				m[n-1-i] = allPossibleFBT(n - 1 - i)
			}
			for _, right := range m[n-1-i] {
				temp := new(TreeNode)
				temp.Left = left
				temp.Right = right
				ans = append(ans, temp)
			}
		}
	}
	return ans
}

type FreqStack struct {
	freM map[int]int
	freS [][]int
}

func Constructor() FreqStack {
	return FreqStack{
		freM: make(map[int]int),
		freS: make([][]int, 0),
	}
}

func (this *FreqStack) Push(val int) {
	this.freM[val]++
	if len(this.freS) < this.freM[val] {
		this.freS = append(this.freS, []int{val})
	} else {
		this.freS[this.freM[val]-1] = append(this.freS[this.freM[val]-1], val)
	}
}

func (this *FreqStack) Pop() int {
	//fmt.Println(this.freM,this.freS)
	lastArr := this.freS[len(this.freS)-1]
	res := lastArr[len(lastArr)-1]
	this.freS[len(this.freS)-1] = lastArr[0 : len(lastArr)-1]
	if len(this.freS[len(this.freS)-1]) == 0 {
		this.freS = this.freS[0 : len(this.freS)-1]
	}
	this.freM[res]--
	return res
}
