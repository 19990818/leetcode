package main

func maxChunksToSorted(arr []int) int {
	ans := 1
	cur := arr[0]
	count := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < cur {
			count--
		} else {
			if count == 0 {
				ans++
			}
			count += arr[i] - cur - 1
			cur = arr[i]
		}
	}
	return ans
}

func isIdealPermutation(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if abs(i-nums[i]) > 1 {
			return false
		}
	}
	return true
}

//双指针循环需要注意没有达到边界的那个指针是否需要继续
//将问题进行归化，同时可以将开始状态和结束状态的变化进行对比
func canTransform(start string, end string) bool {
	n := len(start)
	i, j := 0, 0
	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}
		if i >= n || j >= n {
			return i == j
		}
		if start[i] != end[j] {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i++
		j++
	}
	for i < n {
		if start[i] != 'X' {
			return false
		}
		i++
	}
	for j < n {
		if end[j] != 'X' {
			return false
		}
		j++
	}
	return true
}
