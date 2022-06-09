package main

import (
	"math/rand"
)

type rect struct {
	startx int
	starty int
	width  int
	height int
	points int
}
type Solution struct {
	rectArr []rect
}

func Constructor(rects [][]int) Solution {
	tempArr := make([]rect, 0)
	for i := range rects {
		rectins := rect{}
		rectins.startx = rects[i][0]
		rectins.starty = rects[i][1]
		rectins.width = rects[i][2] - rects[i][0] + 1
		rectins.height = rects[i][3] - rects[i][1] + 1
		rectins.points = rectins.width * rectins.height
		tempArr = append(tempArr, rectins)
	}
	return Solution{tempArr}
}

func (this *Solution) Pick() []int {
	choseRec := this.PickRec()
	x := rand.Intn(this.rectArr[choseRec].width) + this.rectArr[choseRec].startx
	y := rand.Intn(this.rectArr[choseRec].height) + this.rectArr[choseRec].starty
	return []int{x, y}
}

func (this *Solution) PickRec() int {
	sum := 0
	for _, val := range this.rectArr {
		sum += val.points
	}
	randomArea := rand.Intn(sum)
	temp := 0
	for i := range this.rectArr {
		temp += this.rectArr[i].points
		if temp > randomArea {
			return i
		}
	}
	return len(this.rectArr) - 1
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	left, right := 0, buckets
	//二进制思想，将猪的死亡视为一位二进制
	//其中某个二进制表示的值为有毒，那么使用猪能够组成的最大状态数量
	//需要大于水和毒药的总数量,设总共为t轮，因为每个猪最多死一轮,那么每个猪
	//所能表示的最大状态数为(t+1),设总共有x只猪
	var isFinish func(x, t int, total int) bool
	isFinish = func(x, t, total int) bool {
		sum := 1
		for i := x; i > 0; i-- {
			sum *= (t + 1)
			if sum >= total {
				return true
			}
		}
		return sum >= total
	}
	t := minutesToTest / minutesToDie
	for left < right {
		mid := (right-left)>>1 + left
		if isFinish(mid, t, buckets) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	i, j := 0, 0
	s1count, s2count := 0, 0
	for {
		if i == 0 {
			s1count++
		}
		if j == 0 {
			s2count++
		}
		if s1[i] == s2[j] {
			j = (j + 1) % len(s2)
		}
		i = (i + 1) % len(s1)
		if i == 0 && j == 0 {
			break
		}
	}
	return n1 * s1count / n2 * s2count
}
"acb"
4
"ab"
2
 "acb"
1
 "acb"
1