package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var a = "main"
	fmt.Println(a, reflect.TypeOf(a)) // main string
	// go使用自动类型推导

	var b, c int = 12, 22
	fmt.Println(b, c, reflect.TypeOf(b)) // 可以指定类型

	var d = true
	fmt.Println(d, reflect.TypeOf(d)) // ture bool

	var e float64 = 2.22 // 未赋值初始化为0
	f := 32              // 使用:= 与var作用相同
	fmt.Println(e, f)    // 0 32

	g := int(e)
	fmt.Println(g) // 强制类型转换 2

	h := a + " hello world!"
	fmt.Println(h)

	const str string = "const"
	fmt.Println(str)
	const i = 50000
	const j = 3e20 // 3 * 10 ^ 20
	const k = i / j
	fmt.Println(i, j, k, math.Sin(i), math.Sin(k))
}
