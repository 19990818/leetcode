package main

func sampleStats(count []int) []float64 {
	cnt := 0
	c := make([][]int, 0)
	mcnt := 0
	mi, ma := float64(300), float64(0)
	mc := float64(0)
	sum := 0
	for i, v := range count {
		if v > 0 {
			if v > mcnt {
				mc = float64(i)
				mcnt = v
			}
			sum += v * i
			mi = float64(min(int(mi), i))
			ma = float64(max(int(ma), i))
			cnt += v
			c = append(c, []int{cnt, i})
		}
	}
	mid1 := (cnt + 1) / 2
	mid2 := (cnt + 1) / 2
	if cnt%2 == 0 {
		mid2 = cnt/2 + 1
	}
	f1, f2 := false, false
	for _, v := range c {
		if mid1 <= v[0] && !f1 {
			f1 = true
			mid1 = v[1]
		}
		if mid2 <= v[0] && !f2 {
			f2 = true
			mid2 = v[1]
		}
	}
	return []float64{mi, ma, float64(sum) / float64(cnt), float64(mid1+mid2) / float64(2), mc}
}
