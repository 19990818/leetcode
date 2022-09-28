package main

func getKthMagicNumber(k int) int {
	if k == 1 {
		return 1
	}
	queue := []int{1}
	tp, fp, sp := 0, 0, 0
	for k > 1 {
		temp := min(3*queue[tp], 5*queue[fp])
		temp = min(temp, 7*queue[sp])
		queue = append(queue, temp)
		if 3*queue[tp] == temp {
			tp++
		}
		if 5*queue[fp] == temp {
			fp++
		}
		if 7*queue[sp] == temp {
			sp++
		}
		k--
	}
	return queue[len(queue)-1]
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	i := 1
	for ; i < len(preorder); i++ {
		if preorder[i] == postorder[len(postorder)-2] {
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:i], postorder[0:i-1])
	root.Right = constructFromPrePost(preorder[i:], postorder[i-1:len(postorder)-1])
	return root
}

//898 子数组按位或操作 和按位或最大的最小子数组长度题相似
//将arr[j]标记为一个集合 然后进行遍历i，如果不在集合中，就是或操作不相等
//将当前元素加入集合中
func subarrayBitwiseORs(arr []int) int {
	m := make(map[int]int)
	m[arr[0]] = 1
	for i := 1; i < len(arr); i++ {
		m[arr[i]] = 1
		for j := i - 1; j >= 0 && arr[j]|arr[i] != arr[j]; j-- {
			arr[j] |= arr[i]
			m[arr[j]] = 1
		}
	}
	return len(m)
}
