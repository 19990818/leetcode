package main

type Trie struct {
	wordMap   map[string]bool
	prefixMap map[string]bool
}

func ConstructorTrie() Trie {
	return Trie{make(map[string]bool), make(map[string]bool)}
}

func (this *Trie) Insert(word string) {
	this.wordMap[word] = true
	temp := ""
	for _, val := range word {
		temp += string(val)
		this.prefixMap[temp] = true
	}
}

func (this *Trie) Search(word string) bool {
	return this.wordMap[word]
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.prefixMap[prefix]
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func (t *TrieNode) insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

type WordDictionary struct {
	trieRoot *TrieNode
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{&TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.trieRoot.insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(i int, trie *TrieNode) bool
	dfs = func(i int, trie *TrieNode) bool {
		if i == len(word) {
			return trie.isEnd
		}
		ch := word[i]
		if ch != '.' {
			ch -= 'a'
			child := trie.children[ch]
			if child != nil && dfs(i+1, child) {
				return true
			}
		} else {
			for index := range trie.children {
				child := trie.children[index]
				if child != nil && dfs(i+1, child) {
					return true
				}
			}
			return false
		}
		return false
	}
	return dfs(0, this.trieRoot)
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	length1 := ax2 - ax1
	height1 := ay2 - ay1
	length2 := bx2 - bx1
	height2 := by2 - by1
	ans := length1*height1 + length2*height2
	crossX1 := min(ax2, bx2)
	crossY1 := min(ay2, by2)
	crossX2 := max(ax1, bx1)
	crossY2 := max(ay1, by1)
    if (crossX1 - crossX2)>0&&(crossY1 - crossY2)>0{
        ans -= (crossX1 - crossX2) * (crossY1 - crossY2)
    }
	
	return ans
}
