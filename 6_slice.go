package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := make([]string, 3) // 构造容量为3的字符串切片
	slice1[0] = "hello"
	slice1[1] = "world"
	slice1[2] = "!"
	fmt.Println(slice1, slice1[2], len(slice1)) // [hello world !] ! 3

	slice1 = append(slice1, "beautiful")
	slice1 = append(slice1, "world", "!")
	fmt.Println(slice1, len(slice1)) // [hello world ! beautiful world !] 6
	/**
	注意 append 的用法，你必须把 append 的结果赋值为原数组。
	因为 slice 的原理实际上是它有一个它存储了一个长度和一个容量，加一个指向一个数组的指针，在你执行 append 操作的时候，如果容量不够的话，会扩容并且返回新的 slice。
	*/

	copy1 := make([]string, len(slice1))
	copy(copy1, slice1)
	fmt.Println(copy1)

	fmt.Println(slice1[2:5])
	fmt.Println(slice1[:5])
	fmt.Println(slice1[2:])
	// 不支持负数索引

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good, reflect.TypeOf(good))
}
