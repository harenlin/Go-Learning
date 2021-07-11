package test_eg

func Add(x int, y int) int {
	return x + y
}

func Found(target string, names []string) bool {
	for _, n := range names {
		if n == target {
			return true
		}
	}
	return false
}