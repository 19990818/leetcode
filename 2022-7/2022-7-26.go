package main

type CBTInserter struct {
	root *TreeNode
}

func Constructor2(root *TreeNode) CBTInserter {
	return CBTInserter{root: root}
}

func (this *CBTInserter) Insert(val int) int {
	queue := []*TreeNode{this.root}
	for {
		temp := make([]*TreeNode, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: val}
				return cur.Val
			}
			if cur.Right == nil {
				cur.Right = &TreeNode{Val: val}
				return cur.Val
			}
			temp = append(temp, cur.Left, cur.Right)
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
	}
	return 0
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
