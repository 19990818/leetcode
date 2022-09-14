package main

func mirrorReflection(p int, q int) int {
	i := 1
	for {
		if (i*p)%q == 0 {
			if (i*p)%(2*q) == 0 {
				return 2
			}
			if i%2 == 0 {
				return 0
			}
			return 1
		}
		i++
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	left, right, parent := make(map[int]*TreeNode), make(map[int]*TreeNode), make(map[int]*TreeNode)
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		left[r.Val] = r.Left
		right[r.Val] = r.Right
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)
	if k == 0 {
		return []int{target.Val}
	}
	ans := make([]int, 0)
	queue := []*TreeNode{target}
	travel := make(map[*TreeNode]int)
	for k > 0 && len(queue) > 0 {
		temp := make([]*TreeNode, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			travel[cur] = 1
			if left[cur.Val] != nil && travel[left[cur.Val]] == 0 {
				temp = append(temp, left[cur.Val])
			}
			if right[cur.Val] != nil && travel[right[cur.Val]] == 0 {
				temp = append(temp, right[cur.Val])
			}
			if parent[cur.Val] != nil && travel[parent[cur.Val]] == 0 {
				temp = append(temp, parent[cur.Val])
			}
		}
		k--
		if len(temp) > 0 {
			if k == 0 {
				for _, node := range temp {
					ans = append(ans, node.Val)
				}
			}
		}
		queue = temp
	}
	return ans
}
