package main

//得到范围 直接遍历
func cellsInRange(s string) []string {
	col1, col2 := rune(s[0]), rune(s[3])
	row1, row2 := rune(s[1]), rune(s[4])
	ans := make([]string, 0)
	for i := col1; i <= col2; i++ {
		for j := row1; j <= row2; j++ {
			temp := string(i) + string(j)
			ans = append(ans, temp)
		}
	}
	return ans
}

//先对数组排序 然后得到前k个缺失数字的和
func minimalKSum(nums []int, k int) int64 {
	ans := int64(0)
	for i := 0; i <= len(nums); i++ {
		if i == 0 && nums[i] > 1 {
			j := min(k, nums[0]-1)
			ans += getRangeSum(1, nums[0]+j+1)
			k -= j
		} else if nums[i] > nums[i-1]+1 {
			j := min(k, nums[i]-nums[i-1]-1)
			ans += getRangeSum(nums[i-1]+1, nums[i-1]+j)
			k -= j
		} else if i == len(nums) {
			ans += getRangeSum(nums[i-1]+1, nums[i-1]+k)
		}
	}
	return ans
}
func getRangeSum(left, right int) int64 {
	if left < right {
		return 0
	}
	return int64(right+left) * int64(right-left+1) / 2
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	m := make(map[int]*TreeNode)
	mFlag := make(map[int]int)
	for _, val := range descriptions {
		var temp *TreeNode
		if _, ok := m[val[0]]; !ok {
			temp := new(TreeNode)
			temp.Val = val[0]
			m[val[0]] = temp
			if _, ok := mFlag[val[0]]; !ok {
				mFlag[val[0]] = 0
			}
		}
		temp = m[val[0]]
		var tempchild *TreeNode
		if _, ok := m[val[1]]; !ok {
			tempchild := new(TreeNode)
			tempchild.Val = val[1]
			m[val[1]] = tempchild
		}
		tempchild = m[val[1]]
		mFlag[val[1]] = 1
		if val[2] == 1 {
			temp.Left = tempchild
		} else {
			temp.Right = tempchild
		}
	}
	var ans *TreeNode
	for key, val := range mFlag {
		if val == 0 {
			ans = m[key]
			break
		}
	}
	return ans
}
