package main

type cordinate struct {
	x      int
	y      int
	height int
}

func pacificAtlantic(heights [][]int) [][]int {
	atlantic := make([]cordinate, 0)
	pacific := make([]cordinate, 0)
	m, n := len(heights), len(heights[0])
	//fmt.Println(m, n)
	var newArr func(m, n int) [][]int
	newArr = func(m, n int) [][]int {
		flag := make([][]int, m)
		for i := 0; i < m; i++ {
			flag[i] = make([]int, n)
		}
		return flag
	}
	flag := newArr(m, n)
	isShowup1, isShowup2 := newArr(m, n), newArr(m, n)
	//fmt.Println(isShowup1,isShowup2)
	//fmt.Println(flag)
	for j := 0; j < n; j++ {
		atlantic = append(atlantic, cordinate{m - 1, j, heights[m-1][j]})
		flag[m-1][j] += 1
		//fmt.Println(flag[m-1][j])
		isShowup1[m-1][j] = 1
		pacific = append(pacific, cordinate{0, j, heights[0][j]})
		flag[0][j] += 1
		isShowup2[0][j] = 1
	}
	for i := 0; i < m-1; i++ {
		atlantic = append(atlantic, cordinate{i, n - 1, heights[i][n-1]})
		flag[i][n-1]++
		isShowup1[i][n-1] = 1
	}
	for i := 1; i < m; i++ {
		pacific = append(pacific, cordinate{i, 0, heights[i][0]})
		flag[i][0]++
		isShowup2[i][0] = 1
	}
	//fmt.Println(flag)
	queque := append([]cordinate{}, atlantic...)
	for len(queque) > 0 {
		cur := queque[0]
		queque = queque[1:]
		temp := getAroundLand(cur.x, cur.y, m, n, heights)
		//fmt.Println("temp", temp)
		for _, val := range temp {
			//fmt.Println(isShowup1[val.x][val.y])
			if isShowup1[val.x][val.y] == 0 {
				isShowup1[val.x][val.y] = 1
				queque = append(queque, val)
				flag[val.x][val.y]++
			}
		}
		//fmt.Println(queque)
	}

	queque = append([]cordinate{}, pacific...)
	for len(queque) > 0 {
		cur := queque[0]
		queque = queque[1:]
		temp := getAroundLand(cur.x, cur.y, m, n, heights)

		for _, val := range temp {
			if isShowup2[val.x][val.y] == 0 {
				isShowup2[val.x][val.y] = 1
				queque = append(queque, val)
				flag[val.x][val.y]++
			}
		}
	}
	//fmt.Println(flag)
	ans := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if flag[i][j] == 2 {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
func getAroundLand(i, j, m, n int, heights [][]int) []cordinate {
	ans := make([]cordinate, 0)
	if i+1 < m {
		if heights[i+1][j] >= heights[i][j] {
			ans = append(ans, cordinate{i + 1, j, heights[i+1][j]})
		}
	}
	if i-1 >= 0 {
		if heights[i-1][j] >= heights[i][j] {
			ans = append(ans, cordinate{i - 1, j, heights[i-1][j]})
		}
	}
	if j-1 >= 0 && heights[i][j-1] >= heights[i][j] {
		ans = append(ans, cordinate{i, j - 1, heights[i][j-1]})
	}
	if j+1 < n && heights[i][j+1] >= heights[i][j] {
		ans = append(ans, cordinate{i, j + 1, heights[i][j+1]})
	}
	return ans
}
