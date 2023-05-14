package main

import "fmt"

func main() {

	var array [5]int
	fmt.Println(array, len(array)) // [0 0 0 0 0] 5

	var b = [3]int{1, 2, 3}
	fmt.Println(b)

	var array2D [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			array2D[i][j] = i + j
		}
	}
	fmt.Println("array2D\n", array2D)

	// 在真实业务中，我们很少使用数组，因为它的长度是固定的，我们一般使用切片
}
