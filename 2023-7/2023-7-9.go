package main

const mx = int(1e6)

var primes = make([]int, 78498)
var np = make([]bool, mx+1)

func init() {
	pi := 0
	for i := 2; i <= mx; i++ {
		if !np[i] {
			primes[pi] = i
			pi++
			for j := i; j <= mx/i; j++ {
				np[i*j] = true
			}
		}
	}
}
func findPrimePairs(n int) [][]int {
	res := make([][]int, 0)
	if n%2 == 1 {
		if n > 4 && !np[n-2] {
			return [][]int{{2, n - 2}}
		}
		return res
	}
	for _, x := range primes {
		y := n - x
		if x > y {
			break
		}
		if !np[y] {
			res = append(res, []int{x, y})
		}
	}
	return res
}
