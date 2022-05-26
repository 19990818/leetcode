package main

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	type travel struct {
		dis int
		out int
	}
	m := make([][]travel, n)
	out := make([][]int, n)
	in := make([]int, 0)
	for _, val := range flights {
		m[val[0]] = append(m[val[0]], travel{val[2], val[1]})
		out[val[0]] = append(out[val[0]], val[1])
		in[val[1]]++
	}
	topoArr := topoSort(out, in)
	dis := make([]int, n)
	i := 0
	for i < len(topoArr) && topoArr[i] != src {
		dis[topoArr[i]] = -1
		i++
	}
	i = i - 1
	dis[i] = 0
	for i < len(topoArr) {
		val := topoArr[i]
		for _, v := range m[val] {
			if dis[v.out] == -1 {
				dis[v.out] = dis[val] + v.dis
			}
			dis[v.out] = min(dis[v.out], dis[val]+v.dis)
		}
	}
	return dis[dst]
}
func topoSort(out [][]int, in []int) []int {
	ans := make([]int, 0)
	zero := make([]int, 0)
	for idx, val := range in {
		if val == 0 {
			zero = append(zero, idx)
		}
	}
	for len(zero) > 0 {
		cur := zero[0]
		ans = append(ans, cur)
		zero = zero[1:]
		for _, val := range out[cur] {
			in[val]--
			if in[val] == 0 {
				zero = append(zero, val)
			}
		}
	}
	return ans
}

func rotatedDigits(n int) int {
	//2,5,6,9  0,1,8组合
	//最多四位
	ans := 0
	for i := 2; i <= n; i++ {
		num := split(i)
		if num != 0 && num != i {
			//fmt.Println(num,i)
			ans++
		}
	}
	return ans
}

func split(n int) int {
	if n == 0 {
		return 0
	}
	m := map[int]int{0: 0, 1: 1, 2: 5, 5: 2, 6: 9, 8: 8, 9: 6}
	ans := 0
	arrT := make([]int, 0)
	for n > 0 {
		temp := n % 10
		if temp == 3 || temp == 4 || temp == 7 {
			return 0
		}
		arrT = append(arrT, temp)
		n /= 10
	}
	for i := len(arrT) - 1; i >= 0; i-- {
		ans = ans*10 + m[arrT[i]]
	}
	return ans
}
