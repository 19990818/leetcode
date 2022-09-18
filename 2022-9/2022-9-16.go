package main

import "sort"

// #850 矩形面积II
// 解题思路 降维+散点化+扫描线
// 我们可以将所有高度降维到y轴上 然后判断每个高度中是否有x对应
// 如果这个高度在某个x上存在说明这部分需要面积
// 同时为了减少重复计算 我们选择开始那边进行计算后，结束
// 不再进行计算，采用差分法，可以标记这一段宽度是否结束
// 其中对所有相同x进行处理的时候 需要将所有的这一段的x的差分均进行计算
// 因为实际上还是需要将所有矩形进行讨论 这一段相同的x只是特殊化放在一起处理
func rectangleArea(rectangles [][]int) int {
	mod := int(1e9 + 7)
	ans := 0
	// 首先降维得到散点的高度值阵列 然后将高度值排序
	// 得到高度阵列
	n := len(rectangles) * 2
	yvs := make([]int, 0, n)
	// 差分过程需要对有序化的过程 差分实际上是在x轴上进行
	// 将x轴上的点进行排序 d是差分使用的标记
	type xpair struct {
		idx int
		x   int
		d   int
	}
	xvs := make([]xpair, 0, n)
	for i, recrectangle := range rectangles {
		yvs = append(yvs, recrectangle[1], recrectangle[3])
		xvs = append(xvs, xpair{i, recrectangle[0], 1}, xpair{i, recrectangle[2], -1})
	}
	//去重高度值
	sort.Ints(yvs)
	m := 0
	for _, yv := range yvs[1:] {
		if yv != yvs[m] {
			m++
			yvs[m] = yv
		}
	}
	yvs = yvs[0 : m+1]
	sort.Slice(xvs, func(i, j int) bool {
		return xvs[i].x < xvs[j].x
	})
	seg := make([]int, m)
	for i := 0; i < n; i++ {
		j := i
		// 找到第一个不为x的值 可以将x值相同的一起处理
		for j+1 < n && xvs[j+1].x == xvs[i].x {
			j++
		}
		if j+1 == n {
			break
		}
		// 开始差分
		cover := 0
		for x := i; x <= j; x++ {
			curHigh, curLow := rectangles[xvs[x].idx][3], rectangles[xvs[x].idx][1]
			for k := 0; k < m; k++ {
				if curLow <= yvs[k] && yvs[k+1] <= curHigh {
					seg[k] += xvs[x].d
				}
			}
		}
		for k := 0; k < m; k++ {
			if seg[k] > 0 {
				cover += (yvs[k+1] - yvs[k])
			}
		}
		ans += cover * (xvs[j+1].x - xvs[i].x)
		i = j
	}
	return ans % mod
}
