package main

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordsLen := len(words)
	wM := make(map[string]int)
	for _, val := range words {
		wM[val]++
	}
	ans := make([]int, 0)
	//开始位置
	for i := 0; i < wordLen; i++ {
		wordsM := make(map[string]int)
		for k, v := range wM {
			wordsM[k] = v
		}
		//fmt.Println(wordsM)
		for k := i; k <= i+wordsLen*wordLen-wordLen && k <= len(s)-wordLen; k = k + wordLen {
			wordsM[s[k:k+wordLen]]--
			if wordsM[s[k:k+wordLen]] == 0 {
				delete(wordsM, s[k:k+wordLen])
			}
		}
		if len(wordsM) == 0 {
			ans = append(ans, i)
		}
		for j := i + wordLen; j <= len(s)-wordLen*wordsLen; j = j + wordLen {
			//fmt.Println(j,wordsM)
			wordsM[s[j-wordLen:j]]++
			if wordsM[s[j-wordLen:j]] == 0 {
				delete(wordsM, s[j-wordLen:j])
			}
			wordsM[s[j+wordLen*wordsLen-wordLen:j+wordLen*wordsLen]]--
			if wordsM[s[j+wordLen*wordsLen-wordLen:j+wordLen*wordsLen]] == 0 {
				delete(wordsM, s[j+wordLen*wordsLen-wordLen:j+wordLen*wordsLen])
			}
			if len(wordsM) == 0 {
				ans = append(ans, j)
			}
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//树一般直接使用递归方式求解，查找的节点可以用其子树的叶子节点等效替代
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	} else {
		suc := root.Right
		for suc.Left != nil {
			suc = suc.Left
		}
		suc.Right = deleteNode(root.Right, suc.Val)
		suc.Left = root.Left
		return suc
	}
	return root
}

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	for i, val := range s {
		if val == ')' && len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}
	if len(stack) == 0 {
		return len(s)
	}
	ans := stack[0]
	stack = append(stack, len(s))
	for i := 1; i < len(stack); i++ {
		ans = max(ans, stack[i]-stack[i-1]-1)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
