package partition

// Zip64 is designed to gracefully return two ints. For use, see
// example for Partition64.
type Zip64 struct {
	// len(Left), len(Right) may be not equalvent.
	Left, Right []int64
	Pos         int
}

// Valid reports whether further iteration is valid or not.
func (z *Zip64) Valid() bool {
	return z.Pos < len(z.Left) && z.Pos < len(z.Right)
}

// Next increments z.Pos by 1.
func (z *Zip64) Next() {
	z.Pos++
}

// Fetch returns two elements pointed by z.Left and z.Right.
func (z Zip64) Fetch() (i int64, j int64) {
	return z.Left[z.Pos], z.Right[z.Pos]
}

// First reports whether the iteration is first or not.
func (z Zip64) First() bool {
	return z.Pos == 0 && len(z.Left) > 0 && len(z.Right) > 0
}

// Last reports whether the iteration is last or not.
func (z Zip64) Last() bool {
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
func Partition64(x, n int64) Zip64 {
	if x == 0 {
		return laid64([]int64{0, 0})
	}
	if n == 0 {
		return Zip64{}
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
		return laid64(s)
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
	return laid64(s)
}

func laid64(s []int64) Zip64 {
	return Zip64{
		Left:  s,
		Right: s[1:],
	}
}

// Zip is designed to gracefully return two ints. For use, see
// example for Partition64.
type Zip struct {
	// len(Left), len(Right) may be not equalvent.
	Left, Right []int
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
func (z Zip) Fetch() (i int, j int) {
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
func Partition(x, n int) Zip {
	if x == 0 {
		return laid([]int{0, 0})
	}
	if n == 0 {
		return Zip{}
	}
	var u, r, y int
	if n > 0 {
		if x < 0 {
			y = x
			x = -x
		}
		u, r = x/n, x%n
		s := make([]int, n+1)
		for i := range s {
			z := y + u
			if r > 0 {
				r--
				z++
			}
			s[i] = y
			y = z
		}
		return laid(s)
	}

	n *= -1
	if x > 0 {
		y = x
	} else {
		x *= -1
	}
	u, r = x/n, x%n
	s := make([]int, n+1)
	for i := range s {
		z := y - u
		if r > 0 {
			r--
			z--
		}
		s[i] = y
		y = z
	}
	return laid(s)
}

func laid(s []int) Zip {
	return Zip{
		Left:  s,
		Right: s[1:],
	}
}
