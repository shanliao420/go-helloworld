package main

import "fmt"

func add2(num int) {
	num += 2
}

func add2Pointer(num *int) {
	*num += 2
}

func main() {
	num := 0
	add2(num)
	fmt.Println(num)

	add2Pointer(&num)
	fmt.Println(num)

}
