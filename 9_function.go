package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add1(a, b int) int {
	return a + b
}

func exists(m map[string]string, key string) (string, bool) {
	value, ok := m[key]
	return value, ok
}

func exists1(m map[string]string, key string) (value string, ok bool) {
	value, ok = m[key]
	return value, ok
}

func main() {
	result := add(1, 3)
	fmt.Println(result)

	m := map[string]string{"key": "value"}
	fmt.Println(exists(m, "key"))
	fmt.Println(exists1(m, "notKey"))
}
