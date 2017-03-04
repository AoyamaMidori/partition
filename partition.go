package partition

// Partition partitions a slice with length x into n sub slices.
// direction towards which partition goes is determined by the sign of
// n, and the start of the partitioned slices is by x.
//
// For example:
//
//   n > 0: 0 -> x
//   n < 0: x -> 0
//
func Partition(x, n int64) [][2]int64 {
	if x == 0 {
		return [][2]int64{{0, 0}}
	}
	if n == 0 {
		return [][2]int64{}
	}
	var u, r, y int64
	if n > 0 {
		if x < 0 {
			y = x
			x = -x
		}
		u, r = x/n, x%n
		s := make([][2]int64, n)
		for i := range s {
			z := y + u
			if r > 0 {
				r--
				z++
			}
			s[i] = [2]int64{y, z}
			y = z
		}
		return s
	}

	n *= -1
	if x > 0 {
		y = x
	} else {
		x *= -1
	}
	u, r = x/n, x%n
	s := make([][2]int64, n)
	for i := range s {
		z := y - u
		if r > 0 {
			r--
			z--
		}
		s[i] = [2]int64{y, z}
		y = z
	}
	return s
}
