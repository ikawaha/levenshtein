package tiny

func min(x, y, z int) int {
	if x > y {
		if y > z {
			return z
		}
		return y
	} else if x > z {
		return z
	}
	return x
}

func Distance(x, y []rune) int {
	h, w := len(x)+1, len(y)+1
	m := make([][]int, h)
	for i := 0; i < h; i++ {
		m[i] = make([]int, w)
		m[i][0] = i
	}
	for j := 0; j < w; j++ {
		m[0][j] = j
	}
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if x[i-1] == y[j-1] {
				m[i][j] = m[i-1][j-1]
			} else {
				m[i][j] = 1 + min(m[i-1][j-1], m[i-1][j], m[i][j-1])
			}
		}
	}
	return m[len(x)][len(y)]
}
