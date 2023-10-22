package main

import "sort"

func replaceValueInTree(root *TreeNode) *TreeNode {
	// 首先层序遍历一次 得到每层的sum
	sumA := make([]int, 0)
	sumA = append(sumA, root.Val)
	brother := make(map[*TreeNode]int)
	q := []*TreeNode{root}
	for len(q) > 0 {
		temp := q
		q = nil
		sum := 0
		for len(temp) > 0 {
			if temp[0].Left != nil && temp[0].Right != nil {
				brother[temp[0].Left] = temp[0].Right.Val
				brother[temp[0].Right] = temp[0].Left.Val
			}
			if temp[0].Left != nil {
				q = append(q, temp[0].Left)
				sum += temp[0].Left.Val
			}
			if temp[0].Right != nil {
				q = append(q, temp[0].Right)
				sum += temp[0].Right.Val
			}
			temp[0].Val = sumA[len(sumA)-1] - temp[0].Val - brother[temp[0]]
			temp = temp[1:]
		}
		if sum != 0 {
			sumA = append(sumA, sum)
		}
	}
	return root
}

func rowAndMaximumOnes(mat [][]int) []int {
	res := []int{0, 0}
	for i, v := range mat {
		cnt := 0
		for _, v2 := range v {
			if v2 == 1 {
				cnt++
			}
		}
		if cnt > res[1] {
			res = []int{i, cnt}
		}
	}
	return res
}

func maxDivScore(nums []int, divisors []int) int {
	mcnt, res := 0, divisors[0]
	sort.Ints(divisors)
	res = divisors[0]
	for _, d := range divisors {
		cnt := 0
		for _, n := range nums {
			if n%d == 0 {
				cnt++
			}
		}
		if cnt > mcnt {
			res = d
			mcnt = cnt
		}
	}
	return res
}

func addMinimum(word string) int {
	word = "00" + word + "00"
	res := 0
	for i := 2; i < len(word)-2; i++ {
		switch word[i] {
		case 'a':
			res += diff(word[i:i+3], "abc")
			i += 2 - diff(word[i:i+3], "abc")
		case 'b':
			res += 2
			if word[i+1] == 'c' {
				i++
				res--
			}
		case 'c':
			res += 2
		}
	}
	return res
}
func diff(a, b string) int {
	cnt := 0
	i := 0
	for _, v := range a {
		for i < len(b) && byte(v) != b[i] {
			i++
		}
		if i < len(b) {
			cnt++
			i++
		}
	}
	return len(a) - cnt
}

func minHeightShelves(books [][]int, shelfWidth int) int {
	// fmt.Println(books)
	dp := make([]int, len(books)+1)
	for i := range dp {
		dp[i] = 1e9
	}
	dp[0] = 0
	for i, v := range books {
		h := v[1]
		cst := 0
		for j := i; j >= 0; j-- {
			cst += books[j][0]
			h = max(books[j][1], h)
			if cst > shelfWidth {
				break
			}
			dp[i+1] = min(dp[i+1], dp[j]+h)
		}
	}
	return dp[len(books)]
}
