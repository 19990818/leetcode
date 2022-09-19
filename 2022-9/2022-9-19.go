package main

func sumPrefixScores(words []string) []int {
	root := &preTree{key: '#', val: 0}
	for _, val := range words {
		root.Insert(val)
	}
	ans := make([]int, len(words))
	for i, val := range words {
		ans[i] = root.Query(val)
	}
	return ans
}

type preTree struct {
	key   byte
	val   int
	child [26]*preTree
}

func (r *preTree) Insert(str string) {
	root := r
	for i := 0; i < len(str); i++ {
		if root.child[int(str[i]-'a')] == nil {
			root.child[int(str[i]-'a')] = &preTree{key: str[i], val: 1}
		} else {
			root.child[int(str[i]-'a')].val += 1
		}
		root = root.child[int(str[i]-'a')]
	}
}

func (r *preTree) Query(str string) int {
	root := r
	ans := 0
	for i := 0; i < len(str); i++ {
		if root.child[int(str[i]-'a')] == nil {
			break
		}
		ans += root.child[int(str[i]-'a')].val
		root = root.child[int(str[i]-'a')]
	}
	return ans
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(height(root.Left), height(root.Right)) + 1
	}
	if height(root.Left) == height(root.Right) {
		return root
	}
	if height(root.Left) > height(root.Right) {
		return subtreeWithAllDeepest(root.Left)
	}
	return subtreeWithAllDeepest(root.Right)
}
