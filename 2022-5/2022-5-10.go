package main

type trie struct {
	val  int
	next [26]*trie
}
type MapSum struct {
	m *trie
}

func ConstructorM() MapSum {
	return MapSum{new(trie)}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.m
	for i := 0; i < len(key); i++ {
		if cur.next[key[i]-'a'] == nil {
			cur.next[key[i]-'a'] = new(trie)
		}
		cur = cur.next[key[i]-'a']
	}
	cur.val = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.m
	var sumVal func(root *trie) int
	sumVal = func(root *trie) int {
		temp := 0
		if root == nil {
			return 0
		}
		temp += root.val
		for i := 0; i < 26; i++ {
			temp += sumVal(root.next[i])
		}
		return temp
	}
	for i := 0; i < len(prefix); i++ {
		if cur.next[prefix[i]-'a'] == nil {
			return 0
		}
		cur = cur.next[prefix[i]-'a']
	}
	return sumVal(cur)
}

func checkValidString(s string) bool {
	joker := make([]int, 0)
	stack := make([]int, 0)
	for i, val := range s {
		if val == ')' && len(joker) == 0 && len(stack) == 0 {
			return false
		}
		if val == '(' {
			stack = append(stack, i)
		} else if val == '*' {
			joker = append(joker, i)
		} else {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			} else {
				joker = joker[1:len(joker)]
			}
		}
	}
	// fmt.Println(joker,stack)
	if len(stack) == 0 {
		return true
	}
	if len(joker) < len(stack) {
		return false
	}
	right := 0
	count := 0
	for _, val := range stack {
		for right < len(joker) && joker[right] <= val {
			right++
		}
		if right != len(joker) {
			count++
		}
		right++
	}
	return count == len(stack)
}
