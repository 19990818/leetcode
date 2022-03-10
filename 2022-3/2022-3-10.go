package main

import (
	"math/rand"
)

type RandomizedSet struct {
	m    map[int]int
	list []int
}

func ConstructorRandom() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.m[val] = len(this.list)
		this.list = append(this.list, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; ok {
		temp := this.list[len(this.list)-1]
		this.list[this.m[val]] = temp
		this.list = this.list[0 : len(this.list)-1]
		this.m[temp] = this.m[val]
		delete(this.m, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

type Solution struct {
	list []int
}

func ConstructorSolution(head *ListNode) Solution {
	cur := head
	list := make([]int, 0)
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}
	return Solution{list}
}

func (this *Solution) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

type Solution2 struct {
	origin []int
	m      map[int]int
}

func Constructor2(nums []int) Solution2 {
	return Solution2{nums, make(map[int]int)}
}

func (this *Solution2) Reset() []int {
	this.m = make(map[int]int)
	return this.origin
}

func (this *Solution2) Shuffle() []int {
	temp := make([]int, 0)
	for len(temp) != len(this.origin) {
		for {
			randNum := rand.Intn(len(this.origin))
			//fmt.Println(randNum)
			if _, ok := this.m[randNum]; !ok {
				this.m[randNum] = 1
				temp = append(temp, this.origin[randNum])
				break
			}
		}
		//fmt.Println(temp)
	}
	this.m = make(map[int]int)
	return temp
}

type Solution3 struct {
	origin []int
	cur    []int
}

func Constructor3(nums []int) Solution3 {
	return Solution3{nums, append([]int{}, nums...)}
}

func (this *Solution3) Reset() []int {
	copy(this.cur, this.origin)
	return this.origin
}

func (this *Solution3) Shuffle() []int {
	for i := range this.cur {
		j := i + rand.Intn(len(this.cur)-i)
		this.cur[i], this.cur[j] = this.cur[j], this.cur[i]
	}
	return this.cur
}
