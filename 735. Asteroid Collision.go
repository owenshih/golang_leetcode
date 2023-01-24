package golang

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, v := range asteroids {

		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}

		if stack[len(stack)-1]*v > 0 || (stack[len(stack)-1] < 0 && v > 0) {
			stack = append(stack, v)
			continue
		}

		collisionEnd := false
		for !collisionEnd {
			if len(stack) == 0 || stack[len(stack)-1]*v > 0 {
				stack = append(stack, v)
				collisionEnd = true
				continue
			}
			lastInt := abs(stack[len(stack)-1])
			newInt := abs(v)
			if lastInt > newInt {
				collisionEnd = true
				continue
			}
			if lastInt == newInt {
				stack = stack[:len(stack)-1]
				collisionEnd = true
				continue
			}
			if lastInt < newInt {
				stack = stack[:len(stack)-1]
			}
		}

	}

	return stack
}

func abs(i int) int {
	if i < 0 {
		i = 0 - i
	}
	return i
}
