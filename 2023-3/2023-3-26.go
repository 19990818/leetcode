package main

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if numOnes >= k {
		return k
	}
	if numOnes+numZeros >= k {
		return numOnes
	}
	return numOnes - (k - numOnes - numZeros)
}

func primeSubOperation(nums []int) bool {
	primes := getPrimes()
	//fmt.Println(primes)
	t := 0
	for _, v := range nums {
		if v <= t {
			return false
		}
		for j := len(primes) - 1; j >= 0; j-- {
			if v-primes[j] > t {
				t = v - primes[j]
				break
			}
		}

	}
	return true
}
func getPrimes() []int {
	res := make([]int, 0)
	res = append(res, 0)
	for i := 2; i <= 1000; i++ {
		flag := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}
	return res
}
