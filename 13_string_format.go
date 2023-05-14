package main

import "fmt"

type pointer struct {
	x, y int
}

func main() {
	n := 12
	s := "hello"
	p := pointer{1, 2}
	fmt.Printf("s = %v\n", s)
	fmt.Printf("n = %v\n", n)
	fmt.Printf("p = %v\n", p) // p = {1 2}

	fmt.Printf("p = %+v\n", p) //p = {x:1 y:2}
	fmt.Printf("p = %#v\n", p) // p = main.pointer{x:1, y:2}

	f := 3.1415926
	fmt.Printf("f = %.2f\n", f) // f = 3.14

}
