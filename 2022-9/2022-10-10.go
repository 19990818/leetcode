package main

import "math"

func minSwap(nums1 []int, nums2 []int) int {
	//对于每个位置我们有交换和不交换两种选择 初始化一个元素
	//为什么可以这么做，我们设想当子问题为i-1时候我们最终结束的值只能为两个数组中的一个
	//将以最后一个值视为一个整体 然后我们就可以得到一个
	// a   b
	// c   d这样形式的例子，那么我们将进行ac或者bd的交换 只有这两种状态
	//假设我们当前位置位于bd 我们可能将ac交换了 也可能没交换 得到bd不进行交换得到最小解
	//为ac交换或者不进行交换中的较小值 同理bd进行交换也是其中的较小值
	//动态规划将对所有情况都进行讨论
	dp1, dp2 := 0, 1
	for i := 1; i < len(nums1); i++ {
		//将a,b分别置为当前的选择的不进行交换和进行交换的答案
		a, b := math.MaxInt64, math.MaxInt64
		//主要针对nums1[i-1],nums1[i],nums2[i-1],nums2[i]进行讨论
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			a = min(a, dp1)
			b = min(b, dp2+1)
		}
		if nums2[i] > nums1[i-1] && nums1[i] > nums2[i-1] {
			a = min(a, dp2)
			b = min(b, dp1+1)
		}
		dp1, dp2 = a, b
	}
	return min(dp1, dp2)
}
