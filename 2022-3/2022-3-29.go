package main

import (
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if validIpv4(queryIP) {
		return "IPv4"
	}
	if validIpv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}
func validIpv4(queryIP string) bool {
	ipArr := strings.Split(queryIP, ".")
	if len(ipArr) != 4 {
		return false
	}
	for _, val := range ipArr {
		t, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		if len(val) > 1 && val[0] == '0' {
			return false
		}
		if t < 0 || t > 255 {
			return false
		}
	}
	return true
}
func validIpv6(queryIP string) bool {
	ipArr := strings.Split(queryIP, ":")
	if len(ipArr) != 8 {
		return false
	}
	for _, val := range ipArr {
		if len(val) < 1 || len(val) > 4 {
			return false
		}
		for _, val2 := range val {
			if !(val2 >= '0' && val2 <= '9') && !(val2 >= 'a' && val2 <= 'f') && !(val2 >= 'A' && val2 <= 'F') {
				return false
			}
		}
	}
	return true
}

func findRadius1(houses []int, heaters []int) int {
	ans := 0
	for _, house := range houses {
		temp := math.MaxInt32
		for _, heater := range heaters {
			temp = min(temp, abs(house-heater))
		}
		ans = max(temp, ans)
	}
	return ans
}

func findRadius(houses []int, heaters []int) int {
	ans := 0
	sort.Ints(houses)
	sort.Ints(heaters)
	j := 0
	for _, val := range houses {
		for j < len(heaters) && heaters[j] < val {
			j++
		}
		if j == 0 {
			ans = max(ans, heaters[j]-val)
		} else if j == len(heaters) {
			ans = max(ans, val-heaters[j-1])
		} else {
			ans = max(ans, min(heaters[j]-val, val-heaters[j-1]))
		}
	}
	return ans
}

type SolutionCir struct {
	radius   float64
	x_center float64
	y_center float64
}

func ConstructorCir(radius float64, x_center float64, y_center float64) SolutionCir {
	return SolutionCir{radius, x_center, y_center}
}

func (this *SolutionCir) RandPoint() []float64 {
	t := 2 * math.Pi * (rand.Float64())
	length := math.Sqrt(rand.Float64() * math.Pow(this.radius, 2))
	x := this.x_center + length*math.Sin(t)
	y := this.y_center + length*math.Cos(t)
	return []float64{x, y}
}
