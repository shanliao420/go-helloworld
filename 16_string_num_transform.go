package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string convert
	f, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("121", 10, 64) // base 进制  十进制
	fmt.Println(i)

	i, _ = strconv.ParseInt("121", 8, 64) // base 进制  八进制
	fmt.Println(i)

	hex, _ := strconv.ParseInt("0x100", 0, 64) // 如果为0则由字符串前缀确定
	fmt.Println(hex)

	n1, _ := strconv.Atoi("323")
	fmt.Println(n1)

	n2, err := strconv.Atoi("A")
	fmt.Println(n2, err)
}
