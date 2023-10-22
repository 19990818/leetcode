package main

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	if xCenter <= x2 && xCenter >= x1 && yCenter <= y2 && yCenter >= y1 {
		return true
	}
	if xCenter <= x2 && xCenter >= x1 {
		return abs(yCenter-y2) <= radius || abs(yCenter-y1) <= radius
	}
	if yCenter <= y2 && yCenter >= y1 {
		return abs(xCenter-x2) <= radius || abs(xCenter-x1) <= radius
	}
	return cntInCir(x1, y1, xCenter, yCenter, radius) || cntInCir(x1, y2, xCenter, yCenter, radius) || cntInCir(x2, y2, xCenter, yCenter, radius) || cntInCir(x2, y1, xCenter, yCenter, radius)
}
func cntInCir(x, y, xcenter, ycenter, r int) bool {
	return (x-xcenter)*(x-xcenter)+(y-ycenter)*(y-ycenter) <= r*r
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
