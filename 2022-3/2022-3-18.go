package main

import "bytes"

func countBattleships(board [][]byte) int {
	//只需要判断为x的左边和上边是否是x 如果是不需要加
	ans := 0
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				if !(i-1 >= 0 && board[i-1][j] == 'X') && !(j-1 >= 0 && board[i][j-1] == 'X') {
					ans++
				}
			}
		}
	}
	return ans
}

//通过字典树进行求解 通过字典来进行存储和查找 能够有效减小空间时间复杂度
type trie struct {
	left, right *trie
}

func (t *trie) addNum(num int) {
	cur := t
	for i := 30; i >= 0; i-- {
		temp := num >> i & 1
		if temp == 0 {
			if cur.left == nil {
				cur.left = &trie{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &trie{}
			}
			cur = cur.right
		}
	}
}
func (t *trie) checkNum(num int) int {
	cur := t
	ans := 0
	for i := 30; i >= 0; i-- {
		temp := num >> i & 1
		if temp == 0 {
			if cur.right != nil {
				ans = ans*2 + 1
				cur = cur.right
			} else {
				ans *= 2
				cur = cur.left
			}
		} else {
			if cur.left != nil {
				ans = ans*2 + 1
				cur = cur.left
			} else {
				ans *= 2
				cur = cur.right
			}
		}
	}
	return ans
}
func findMaximumXOR(nums []int) int {
	t := &trie{}
	ans := 0
	for _, val := range nums {
		t.addNum(val)
	}
	for _, val := range nums {
		ans = max(ans, t.checkNum(val))
	}
	return ans
}

//本质上是解方程的思想，高阶方程，在此处有一些特有元素 可以直接确定元素个数
func originalDigits(s string) string {
	m := make(map[rune]int)
	for _, val := range s {
		m[val]++
	}
	cnt := make([]int, 10)
	cnt[0] = m['z']
	cnt[2] = m['w']
	cnt[4] = m['u']
	cnt[6] = m['x']
	cnt[8] = m['g']
	cnt[5] = m['f'] - cnt[4]
	cnt[3] = m['h'] - cnt[8]
	cnt[7] = m['s'] - cnt[6]
	cnt[9] = m['i'] - cnt[5] - cnt[6] - cnt[8]
	cnt[1] = m['o'] - cnt[0] - cnt[2] - cnt[4]
	ans := []byte{}
	for i, c := range cnt {
		ans = append(ans, bytes.Repeat([]byte{byte('0' + i)}, c)...)
	}
	return string(ans)
}
