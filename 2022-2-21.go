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
	nums []int
	tree []int
}

func ConstructorNumArray(nums []int) NumArray {
	temp := NumArray{nums: nums}
	tree := make([]int, 2*len(nums))
	for i := len(nums); i < 2*len(nums); i++ {
		tree[i] = nums[i-len(nums)]
	}
	for i := len(nums) - 1; i > 0; i-- {
		tree[i] = tree[2*i] + tree[2*i+1]
	}
	return temp
}

func (this *NumArray) Update(index int, val int) {
	pos := index + len(this.nums)
	this.tree[pos] -= (this.nums[index] - val)
	for pos > 1 {
		left, right := pos, pos
		if pos%2 == 0 {
			right++
		} else {
			left--
		}
		this.tree[pos/2] = this.tree[left] + this.tree[right]
		pos /= 2
	}
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for left <= right {
		if left%2 == 1 {
			sum += this.tree[left]
			left++
		}
		if right%2 == 0 {
			sum += this.tree[right]
			right--
		}
		left /= 2
		right /= 2
	}
	return sum
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
