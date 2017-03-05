package partition

// Zip object is designed to gracefully return two ints. For use, see
// example for Partition.
type Zip struct {
	// len(Left), len(Right) may be not equalvent.
	Left, Right []int64
	Pos         int
}

// Valid reports whether further iteration is valid or not.
func (z *Zip) Valid() bool {
	return z.Pos < len(z.Left) && z.Pos < len(z.Right)
}

// Next increments z.Pos by 1.
func (z *Zip) Next() {
	z.Pos++
}

// Fetch returns two elements pointed by z.Left and z.Right.
func (z Zip) Fetch() (i int64, j int64) {
	return z.Left[z.Pos], z.Right[z.Pos]
}

// First reports whether the iteration is first or not.
func (z Zip) First() bool {
	return z.Pos == 0 && len(z.Left) > 0 && len(z.Right) > 0
}

// Last reports whether the iteration is last or not.
func (z Zip) Last() bool {
	return z.Pos+1 == len(z.Left) || z.Pos+1 == len(z.Right)
}

// Partition partitions a slice with length x into n sub slices.
// direction towards which partition goes is determined by the sign of
// n, and the start of the partitioned slices is by x.
//
// For example:
//
//   n > 0: 0 -> x
//   n < 0: x -> 0
//
func Partition(x, n int64) Zip {
	if x == 0 {
		return lean([]int64{0, 0})
	}
	if n == 0 {
		return Zip{}
	}
	var u, r, y int64
	if n > 0 {
		if x < 0 {
			y = x
			x = -x
		}
		u, r = x/n, x%n
		s := make([]int64, n+1)
		for i := range s {
			z := y + u
			if r > 0 {
				r--
				z++
			}
			s[i] = y
			y = z
		}
		return lean(s)
	}

	n *= -1
	if x > 0 {
		y = x
	} else {
		x *= -1
	}
	u, r = x/n, x%n
	s := make([]int64, n+1)
	for i := range s {
		z := y - u
		if r > 0 {
			r--
			z--
		}
		s[i] = y
		y = z
	}
	return lean(s)
}

func lean(s []int64) Zip {
	return Zip{
		Left:  s,
		Right: s[1:],
	}
}
