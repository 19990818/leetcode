package main

import (
	"sort"
	"strconv"
)

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, 10)
	for i := 0; i < 10; i++ {
		if i == 0 {
			dp[i] = 10
		} else if i == 1 {
			dp[i] = 9 * 9
		} else {
			dp[i] = dp[i-1] * (10 - i)
		}
	}
	ans := 0
	for i := 0; i < n && i < 10; i++ {
		ans += dp[i]
	}
	return ans
}

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if targetCapacity > jug1Capacity+jug2Capacity {
		return false
	}
	for i := -jug2Capacity; i <= jug2Capacity; i++ {
		if (targetCapacity-i*jug1Capacity)%jug2Capacity == 0 {
			return true
		}
	}
	return false
}

func mostFrequent(nums []int, key int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			m[nums[i+1]]++
		}
	}
	maxCount := 0
	ans := 0
	for k, val := range m {
		if val > maxCount {
			maxCount = val
			ans = k
		}
	}
	return ans
}

func sortJumbled(mapping []int, nums []int) []int {
	a := arrNum{nums, mapping}
	sort.Sort(a)
	return a.arr
}
func fnum(mapping []int, num int) int {
	ans := 0
	s := strconv.Itoa(num)
	for i := 0; i < len(s); i++ {
		ans = ans*10 + mapping[s[i]-'0']
	}
	return ans
}

type arrNum struct {
	arr     []int
	mapping []int
}

func (m arrNum) Len() int {
	return len(m.arr)
}

func (m arrNum) Less(i, j int) bool {
	return fnum(m.mapping, m.arr[i]) < fnum(m.mapping, m.arr[j])
}

func (m arrNum) Swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func getAncestors(n int, edges [][]int) [][]int {
	m := make(map[int][]int)
	for _, val := range edges {
		m[val[1]] = append(m[val[1]], val[0])
	}
	ans := make([][]int, n)

	for i := 0; i < n; i++ {
		temp := make([]int, 0)
		tempArr := append([]int{}, m[i]...)
		visted := make(map[int]bool)
		for len(tempArr) > 0 {
			if !visted[tempArr[0]] {
				temp = append(temp, tempArr[0])
				visted[tempArr[0]] = true
			}
			for _, val := range m[tempArr[0]] {
				//fmt.Println(visted[val])
				if !visted[val] {
					tempArr = append(tempArr, val)
				}
			}
			tempArr = tempArr[1:]
			//fmt.Println(tempArr)
		}
		sort.Ints(temp)
		ans[i] = temp
	}
	return ans
}

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
