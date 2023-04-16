package main

func isRobotBounded(instructions string) bool {
	now := 0
	x, y := 0, 0
	for i := 0; i < 4; i++ {
		for _, v := range instructions {
			if v == 'L' {
				now = (now + 3) % 4
			} else if v == 'R' {
				now = (now + 1) % 4
			} else {
				switch now {
				case 0:
					y += 1
				case 1:
					x += 1
				case 2:
					y -= 1
				case 3:
					x -= 1
				}
			}
		}
	}
	return x == 0 && y == 0
}
