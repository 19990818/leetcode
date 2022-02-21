package main

type NumMatrix struct {
	matrix    [][]int
	sumMatrix [][]int
}

func ConstructorNumMatrix(matrix [][]int) NumMatrix {
	NumMatrixEx := NumMatrix{matrix: matrix}
	sum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sum[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 && j == 0 {
				sum[i][j] = matrix[0][0]
			} else if j == 0 {
				sum[i][j] = sum[i-1][j] + matrix[i][j]
			} else if i == 0 {
				sum[i][j] = sum[i][j-1] + matrix[i][j]
			} else {
				sum[i][j] = sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1] + matrix[i][j]
			}
		}
	}
	NumMatrixEx.sumMatrix = sum
	return NumMatrixEx
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	//fmt.Println(this.sumMatrix)
	ans := this.sumMatrix[row2][col2]
	if row1 > 0 {
		ans -= this.sumMatrix[row1-1][col2]
	}
	if col1 > 0 {
		ans -= this.sumMatrix[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		ans += this.sumMatrix[row1-1][col1-1]
	}
	return ans
}

type NumArray struct {
	nums   []int
	posSum []int
	differ map[int]int
}

func ConstructorNumArray(nums []int) NumArray {
	temp := NumArray{nums: nums}
	sum := make([]int, len(nums))
	differ := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sum[i] = nums[i]
		} else {
			sum[i] = sum[i-1] + nums[i]
		}
		differ[i] = 0
	}
	temp.posSum = sum
	temp.differ = differ
	return temp
}

func (this *NumArray) Update(index int, val int) {
	this.differ[index] += val - this.nums[index]
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	ans := this.posSum[right] - this.posSum[left] + this.nums[left]
	for _, val := range this.differ {
		ans += val
	}
	return ans
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	ans := 0
	profit := make([]int, 0)
	for i := 1; i < len(prices); i++ {
		profit = append(profit, prices[i]-prices[i-1])
	}
	//fmt.Println(profit)
	dp := make([]int, len(profit))
	dp[0] = profit[0]
	ans = max(ans, dp[0])
	for i := 1; i < len(profit); i++ {
		dp[i] = max(dp[i-1]+profit[i], profit[i])
		for j := i - 3; j >= 0; j-- {
			dp[i] = max(dp[j]+profit[i], dp[i])
		}
		ans = max(dp[i], ans)
	}
	//fmt.Println(dp)
	return ans
}

func timeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	for i := 0; i < len(tickets); i++ {
		if i < k {
			ans += min(tickets[i], tickets[k])
		} else {
			ans += min(tickets[i], tickets[k]-1)
		}
	}
	return ans
}
