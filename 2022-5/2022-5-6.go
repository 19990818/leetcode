package main

func isPossible(nums []int) bool {
	m1 := make(map[int]int)
	for _, val := range nums {
		m1[val]++
	}
	mEnd := make(map[int]int)
	for _, val := range nums {
		if m1[val] == 0 {
			continue
		}
		//fmt.Println(val)
		if mEnd[val-1] <= 0 {
			//此种情况为当前元素需要作为开始
			if m1[val+1] > 0 && m1[val+2] > 0 {
				m1[val]--
				m1[val+1]--
				m1[val+2]--
				mEnd[val+2]++
			} else {
				return false
			}
		} else {
			m1[val]--
			mEnd[val-1]--
			mEnd[val]++
		}
		//fmt.Println(m1,mEnd)
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	number := make([]int, 0)
	number = append(number, 0)
	for {
		temp := make([]*TreeNode, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			num := number[0]
			number = number[1:]
			if cur.Left != nil {
				temp = append(temp, cur.Left)
				number = append(number, num*2+1)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
				number = append(number, num*2+2)
			}
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
		ans = max(ans, number[len(number)-1]-number[0]+1)
	}
	return ans
}

func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	stack := make([]int, 0)
	count := 0
	stack = append(stack, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] >= stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else {
			count++
			if count > 1 {
				return false
			}
			if len(stack) == 1 || nums[i] >= stack[len(stack)-2] {
				stack = stack[0 : len(stack)-1]
				stack = append(stack, nums[i])
			}
		}
	}
	return true
}
