package main

func temperatureTrend(temperatureA []int, temperatureB []int) int {
	n := len(temperatureA)
	ans := 0
	cnt := 0
	for i := 0; i < n-1; i++ {
		diffA := temperatureA[i+1] - temperatureA[i]
		diffB := temperatureB[i+1] - temperatureB[i]
		if diffA*diffB > 0 || (diffA == 0 && diffB == 0) {
			cnt++
		} else {
			ans = max(ans, cnt)
			cnt = 0
		}
	}
	ans = max(ans, cnt)
	return ans
}

func transportationHub(path [][]int) int {
	//直接通往 入度 出度
	out := make(map[int]int)
	in := make(map[int][]int)
	sum := make(map[int]int)
	for _, val := range path {
		out[val[0]]++
		sum[val[0]] = 1
		sum[val[1]] = 1
		in[val[1]] = append(in[val[1]], val[0])
	}
	for k, v := range in {
		m := make(map[int]int)
		for _, u := range v {
			m[u] = 1
		}
		if len(m)+1 == len(sum) && out[k] == 0 {
			return k
		}
	}
	return -1
}

func ballGame(num int, plate []string) [][]int {
	//从右边开始
	ans := make([][]int, 0)
	m, n := len(plate), len(plate[0])
	var run func(status, startx, starty int)
	run = func(status, startx, starty int) {
		if plate[startx][starty] != '.' {
			return
		}
		travel := make(map[int]int)
		curx, cury := startx, starty

		for cnt := 0; cnt <= num; cnt++ {
			//fmt.Println(curx,cury,status)
			if travel[curx*n+cury] == 15 {
				break
			}
			if plate[curx][cury] == 'O' {
				ans = append(ans, []int{startx, starty})
				return
			}
			travel[curx*n+cury] |= (1 << status)
			if plate[curx][cury] == 'W' {
				status = (status + 1) % 4
			} else if plate[curx][cury] == 'E' {
				status = (status + 3) % 4
			}
			switch status {
			case 0:
				cury = cury + 1
			case 1:
				curx = curx - 1
			case 2:
				cury = cury - 1
			case 3:
				curx = curx + 1
			}
			if curx < 0 || curx >= m || cury < 0 || cury >= n {
				break
			}
		}
	}
	for i := 1; i < m-1; i++ {
		run(0, i, 0)
		run(2, i, n-1)
	}
	for j := 1; j < n-1; j++ {
		run(1, m-1, j)
		run(3, 0, j)
	}
	return ans
}
