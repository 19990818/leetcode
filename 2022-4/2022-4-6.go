package main

import "math/rand"

type SolutionPickIndex struct {
	wEnd []int
}

func ConstructorPickIndex(w []int) SolutionPickIndex {
	wEnd := make([]int, 0)
	end := 0
	for i := 0; i < len(w); i++ {
		end = end + w[i] - 1
		wEnd = append(wEnd, end)
	}
	return SolutionPickIndex{wEnd: wEnd}
}

func (this *SolutionPickIndex) PickIndex() int {
	total := this.wEnd[len(this.wEnd)-1]
	x := rand.Intn(total)
	// fmt.Println(x,total)
	for i := 0; i < len(this.wEnd); i++ {
		if i == 0 && x >= 0 && x < this.wEnd[0] {
			return 0
		}
		if i < len(this.wEnd)-1 && x >= this.wEnd[i] && x < this.wEnd[i+1] {
			return i + 1
		}
	}
	return -1
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	var bfs func(queue [][]int)
	m := len(board)
	n := len(board[0])
	var getM func(board [][]byte, i, j int) int
	getM = func(board [][]byte, col, row int) int {
		ans := 0
		for i := col - 1; i <= col+1; i++ {
			for j := row - 1; j <= row+1; j++ {
				if i != col || j != row {
					if i >= 0 && i < m && j >= 0 && j < n && board[i][j] == 'M' {
						ans++
					}
				}
			}
		}
		return ans
	}
	flag := make([][]int, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]int, n)
	}
	bfs = func(queue [][]int) {
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for i := cur[0] - 1; i <= cur[0]+1; i++ {
				for j := cur[1] - 1; j <= cur[1]+1; j++ {
					if i != cur[0] || j != cur[1] {
						if i >= 0 && i < m && j >= 0 && j < n && flag[i][j] == 0 {
							flag[i][j] = 1
							count := getM(board, i, j)
							//fmt.Println(count)
							if count == 0 {
								// fmt.Println(i,j,cur)
								board[i][j] = 'B'
								queue = append(queue, []int{i, j})
							} else {
								board[i][j] = byte(count + '0')
							}
						}
					}
				}
			}
			//fmt.Println(queue)
		}
	}
	flag[click[0]][click[1]] = 1
	count := getM(board, click[0], click[1])
	if count != 0 {
		board[click[0]][click[1]] = byte(count + '0')
		return board
	}
	board[click[0]][click[1]] = 'B'
	queue := make([][]int, 0)
	queue = append(queue, click)
	bfs(queue)
	return board
}

func findPairs(nums []int, k int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	ans := 0
	// fmt.Println(m)
	for key := range m {
		if k == 0 {
			if m[key] > 1 {
				ans++
				delete(m, key)
			}
			continue
		}
		if _, ok := m[key-k]; ok {
			ans++
			delete(m, key)
		}
		if _, ok := m[key+k]; ok {
			ans++
			delete(m, key)
		}
	}
	return ans
}
