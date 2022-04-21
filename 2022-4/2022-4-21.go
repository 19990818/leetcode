package main

type MyCalendar struct {
	arr [][]int
}

func ConstructorCal() MyCalendar {
	return MyCalendar{[][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	l, r := 0, len(this.arr)
	for l < r {
		mid := (r-l)>>1 + l
		maxS, minE := max(start, this.arr[mid][0]), min(end, this.arr[mid][1])
		if maxS < minE {
			return false
		}
		if maxS == start {
			l = mid + 1
		} else {
			r = mid
		}
	}
	ans := append([][]int{}, this.arr[0:l]...)
	ans = append(ans, []int{start, end})
	ans = append(ans, this.arr[l:]...)
	this.arr = ans
	return true
}
