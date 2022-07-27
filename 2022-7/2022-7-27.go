package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

type Skiplist struct {
	m map[int]int
}

func Constructor3() Skiplist {
	return Skiplist{make(map[int]int)}
}

func (this *Skiplist) Search(target int) bool {
	return this.m[target] != 0

}

func (this *Skiplist) Add(num int) {
	this.m[num]++
}

func (this *Skiplist) Erase(num int) bool {
	if this.m[num] == 0 {
		return false
	}
	this.m[num]--
	return true
}

type pair struct {
	rate int
	s    string
}
type FoodRatings struct {
	fs map[string]pair
	cs map[string]*redblacktree.Tree
}

func Constructor4(foods []string, cuisines []string, ratings []int) FoodRatings {
	fs := map[string]pair{}
	cs := map[string]*redblacktree.Tree{}
	for i, food := range foods {
		rate, cuisine := ratings[i], cuisines[i]
		if cs[cuisine] == nil {
			cs[cuisine] = redblacktree.NewWith(func(a, b interface{}) int {
				x, y := a.(pair), b.(pair)
				if x.rate != y.rate {
					return utils.IntComparator(y.rate, x.rate)
				}
				return utils.StringComparator(x.s, y.s)
			})
		}
		fs[food] = pair{rate, cuisine}
		cs[cuisine].Put(pair{rate, food}, nil)
	}
	return FoodRatings{fs, cs}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	p := this.fs[food]
	t := this.cs[p.s]
	t.Remove(pair{p.rate, food})
	t.Put(pair{newRating, food}, nil)
	p.rate = newRating
	this.fs[food] = p
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.cs[cuisine].Left().Key.(pair).s
}

type NumberContainers struct {
	idxM map[int]int
	numM map[int]*redblacktree.Tree
}

func Constructor5() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*redblacktree.Tree{}}
}

func (this *NumberContainers) Change(index int, number int) {
	if _, ok := this.idxM[index]; ok {
		oldNum := this.idxM[index]
		this.numM[oldNum].Remove(index)
	}
	this.idxM[index] = number
	if this.numM[number] == nil {
		this.numM[number] = redblacktree.NewWith(func(a, b interface{}) int {
			return utils.IntComparator(a.(int), b.(int))
		})
	}
	this.numM[number].Put(index, nil)
}

func (this *NumberContainers) Find(number int) int {
	if this.numM[number] == nil || this.numM[number].Empty() {
		return -1
	}
	return this.numM[number].Left().Key.(int)
}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		m[val] = i
	}
	for i, val := range nums {
		if _, ok := m[target-val]; ok {
			return []int{i, m[target-val]}
		}
	}
	return []int{-1, -1}
}
