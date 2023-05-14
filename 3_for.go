package main

import "fmt"

func main() {

	for {
		fmt.Println("loop")
		break
	}
	i := 1
	for i <= 3 {
		fmt.Println("i = ", i)
		i++
	}
	for j := 0; j < 6; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println("j = ", j)
	}
}
