package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	var sum int
	for index, value := range nums {
		sum += value
		fmt.Println("num -", index, " = ", value)
	}
	fmt.Println("sum = ", sum)

	m := map[string]string{"shanliao": "man", "king": "girl"}
	for key, value := range m {
		fmt.Println(key, "is", value)
	}

}
