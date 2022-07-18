package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

//go使用fmt标准化输入输出
// func getMaxBottles(a int) int {
// 	return a / 2
// }
// func main() {
// 	var input int
// 	for {
// 		fmt.Scanln(&input)
// 		if input == 0 {
// 			break
// 		}
// 		fmt.Println(getMaxBottles(input))
// 	}
// 	//fmt.Println(err)
// }

//动态规划不能重复需要从大的开始 从小的开始会在后面重复选择较小值
func main1() {
	var n, m int
	fmt.Scan(&n, &m)
	vs := make([][]int, 0)
	ps := make(map[int]int)
	cs := make(map[int][]int)
	for i := 0; i < m; i++ {
		var v, p, q int
		fmt.Scan(&v, &p, &q)
		if q != 0 {
			ps[i] = q
			cs[q-1] = append(cs[q-1], i)
		}
		vs = append(vs, []int{v, v * p})
	}
	//fmt.Println(vs)
	fmt.Println(dpSolve(vs, n, ps, cs))
}

//vs表示花费和满意度，total为最大消费，ps为子件和主件的关系
func dpSolve(vs [][]int, total int, ps map[int]int, cs map[int][]int) int {
	dp := make([]int, total+1)
	dp[0] = 0
	for idx, val := range vs {
		//为主件
		if ps[idx] != 0 {
			continue
		}

		//还是重复选取了
		for i := total; i > 0; i-- {
			if i >= val[0] {
				dp[i] = max(dp[i], dp[i-val[0]]+val[1])
				//fmt.Println(dp[i])
			}
			if len(cs[idx]) > 0 && i >= val[0]+vs[cs[idx][0]][0] {
				dp[i] = max(dp[i], dp[i-val[0]-vs[cs[idx][0]][0]]+val[1]+vs[cs[idx][0]][1])
				//fmt.Println(dp[i])
			}
			if len(cs[idx]) > 1 && i >= val[0]+vs[cs[idx][1]][0] {
				dp[i] = max(dp[i], dp[i-val[0]-vs[cs[idx][1]][0]]+val[1]+vs[cs[idx][1]][1])
				//fmt.Println(dp[i])
			}
			if len(cs[idx]) > 1 && i >= val[0]+vs[cs[idx][0]][0]+vs[cs[idx][1]][0] {
				dp[i] = max(dp[i], dp[i-val[0]-vs[cs[idx][1]][0]-vs[cs[idx][0]][0]]+val[1]+vs[cs[idx][1]][1]+vs[cs[idx][0]][1])
				//fmt.Println(dp[i])
			}
		}
		//fmt.Println(dp)
	}
	return dp[total]
}

func main2() {
	var s string
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		s = input.Text()
		var temp strings.Builder
		for _, val := range s {
			if unicode.IsLetter(val) {
				temp.WriteRune(val)
			} else {
				temp.WriteRune(' ')
			}
		}
		s = temp.String()
		sArr := strings.Split(s, " ")
		s2Arr := make([]string, 0)
		for i := len(sArr) - 1; i >= 0; i-- {
			if sArr[i] != "" {
				s2Arr = append(s2Arr, sArr[i])
			}
		}
		fmt.Println(strings.Join(s2Arr, " "))
	}
}
