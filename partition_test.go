package partition

import "fmt"

func ExamplePartition() {
	PrintZip := func(x, n int64) {
		fmt.Print("[")
		for z := Partition(x, n); z.Valid(); z.Next() {
			i, j := z.Fetch()
			fmt.Print("[", i, " ", j, "]")
			if !z.Last() {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	}
	PrintZip(10, 3)
	PrintZip(10, -3)
	PrintZip(-10, 3)
	PrintZip(-10, -3)
	// Output:
	// [[0 4] [4 7] [7 10]]
	// [[10 6] [6 3] [3 0]]
	// [[-10 -6] [-6 -3] [-3 0]]
	// [[0 -4] [-4 -7] [-7 -10]]
}
