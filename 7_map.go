package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	fmt.Println(m, len(m))
	fmt.Println(m["one"], m["notKey"]) // 1 0   不存在的key返回0或者nil

	value, ok := m["notKey"]
	fmt.Println(value, ok) // 0 false

	delete(m, "one")
	fmt.Println(m)

	m1 := map[string]int{"one": 1, "two": 2}
	fmt.Println(m1)
}
