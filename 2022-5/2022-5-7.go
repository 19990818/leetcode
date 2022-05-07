package main

import "strconv"

func constructArray(n int, k int) []int {
	ans := make([]int, 0)
	for i := n; i > k; i-- {
		ans = append(ans, i)
	}
	flag := -1
	for len(ans) < n {
		temp := ans[len(ans)-1] + flag*k
		ans = append(ans, temp)
		k--
		flag *= -1
	}
	return ans
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		root = trimBST(root.Right, low, high)
		return root
	}
	if root.Val > high {
		root = trimBST(root.Left, low, high)
		return root
	}
	root.Right = trimBST(root.Right, low, high)
	root.Left = trimBST(root.Left, low, high)
	return root
}

func maximumSwap(num int) int {
	str := strconv.Itoa(num)
	for i := 0; i < len(str)-1; i++ {
		pos := -1
		maxByte := str[i]
		j := i + 1
		ans := 0
		for ; j < len(str); j++ {
			if str[j] >= maxByte && str[j] > str[i] {
				pos = j
				maxByte = str[j]
				strArr := []byte(str)
				strArr[i], strArr[pos] = strArr[pos], strArr[i]
				temp, _ := strconv.Atoi(string(strArr))
				ans = max(ans, temp)
				//fmt.Println(ans)
			}
		}
		if pos != -1 {
			return ans
		}
	}
	return num
}
