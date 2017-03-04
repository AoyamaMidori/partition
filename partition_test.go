package partition

import "fmt"

func ExamplePartition() {
	fmt.Println(Partition(10, 3))
	fmt.Println(Partition(10, -3))
	fmt.Println(Partition(-10, 3))
	fmt.Println(Partition(-10, -3))
	// Output:
	// [[0 4] [4 7] [7 10]]
	// [[10 6] [6 3] [3 0]]
	// [[-10 -6] [-6 -3] [-3 0]]
	// [[0 -4] [-4 -7] [-7 -10]]
}
