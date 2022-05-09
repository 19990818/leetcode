package main

func flipLights(n int, presses int) int {
	//对编号为k,2k,2k+1,3k+1的灯泡进行反转
	//那么找出最大公倍数 之后的状态均是之前的操作能确定下来的
	//将问题规模缩小到6
	shift := max(0, 6-n)
	//枚举每种操作的可能 每种操作都只能是0或1 因为重复两次相当于没有效果
	var getCount1 func(num int) int
	getCount1 = func(num int) int {
		ans := 0
		for num > 0 {
			if num%2 == 1 {
				ans++
			}
			num /= 2
		}
		return ans
	}
	m := make(map[int]int)
	for i := 0; i < 16; i++ {
		light := 0b111111
		if presses%2 == getCount1(i)%2 && getCount1(i) <= presses {
			if i&1 > 0 {
				light ^= 0b111111 >> shift
			}
			if i>>1&1 > 0 {
				light ^= 0b101010 >> shift
			}
			if i>>2&1 > 0 {
				light ^= 0b010101 >> shift
			}
			if i>>3&1 > 0 {
				light ^= 0b100100 >> shift
			}
			m[light] = 1
		}
	}
	return len(m)
}

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dpNum := make([]int, len(nums))
	dp[0] = 1
	dpNum[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] == dp[j]+1 {
				dpNum[i] += dpNum[j]
			}
		}
		if dpNum[i] == 0 {
			dpNum[i] = 1
		}
	}
	ans := 0
	maxNum := dp[0]
	for i := 1; i < len(nums); i++ {
		maxNum = max(maxNum, dp[i])
	}
	for i := 0; i < len(nums); i++ {
		if dp[i] == maxNum {
			ans += dpNum[i]
		}
	}
	return ans
}

type MagicDictionary struct {
	strArr []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.strArr = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	var isSuccess func(src, target string) bool
	isSuccess = func(src, target string) bool {
		if len(src) != len(target) || src == target {
			return false
		}
		for i := 0; i < len(src); i++ {
			if src[i] != target[i] {
				return src[i+1:] == target[i+1:]
			}
		}
		return true
	}
	for _, val := range this.strArr {
		if isSuccess(val, searchWord) {
			return true
		}
	}
	return false
}
