package main

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	exp := 0
	egy := 0
	for _, val := range energy {
		egy += val
	}
	egy = max(egy+1-initialEnergy, 0)
	for _, val := range experience {
		if initialExperience > val {
			initialExperience += val
		} else {
			exp += val + 1 - initialExperience
			initialExperience = 2*val + 1
		}
	}
	return egy + exp
}

func largestPalindromic(num string) string {
	nums := make([]int, 10)
	for _, val := range num {
		nums[int(val-'0')]++
	}
	ans := make([]byte, 0)
	for i := 9; i >= 0; i-- {
		for nums[i] >= 2 {
			if i == 0 && len(ans) == 0 {
				break
			}
			nums[i] -= 2
			ans = append(ans, byte(i+'0'))
		}
	}
	flag := 0
	for i := 9; i >= 0; i-- {
		if nums[i] > 0 {
			ans = append(ans, byte(i+'0'))
			flag = 1
			break
		}
	}
	var reverseArr func(arr []byte) []byte
	reverseArr = func(arr []byte) []byte {
		res := make([]byte, 0)
		for i := len(arr) - 1; i >= 0; i-- {
			res = append(res, arr[i])
		}
		return res
	}
	if flag == 0 {

		ans = append(ans, reverseArr(ans)...)
	} else {
		ans = append(ans, reverseArr(ans[0:len(ans)-1])...)
	}
	return string(ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	//记录每个点的深度
	parent := make(map[int]int)
	lf := make(map[int]int)
	rg := make(map[int]int)
	var treeDepth func(r *TreeNode)
	treeDepth = func(r *TreeNode) {
		if r == nil {
			return
		}

		if r.Left != nil {
			lf[r.Val] = r.Left.Val
			parent[r.Left.Val] = r.Val
			treeDepth(r.Left)
		}
		if r.Right != nil {
			rg[r.Val] = r.Right.Val
			parent[r.Right.Val] = r.Val
			treeDepth(r.Right)
		}
	}
	treeDepth(root)
	travel := make(map[int]int)
	var help func(s int) int
	help = func(s int) int {
		travel[s] = 1
		ans := 0
		if travel[lf[s]] == 0 && lf[s] != 0 {
			travel[lf[s]] = 1
			ans = max(ans, help(lf[s])+1)
		}
		if travel[rg[s]] == 0 && rg[s] != 0 {
			travel[rg[s]] = 1

			ans = max(ans, help(rg[s])+1)
		}
		if travel[parent[s]] == 0 && parent[s] != 0 {
			travel[parent[s]] = 1
			ans = max(help(parent[s])+1, ans)
		}
		return ans
	}
	return help(start)
}
