package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world!"

	fmt.Println(strings.Contains(str, "t"))
	fmt.Println(strings.Count(str, "l"))
	fmt.Println(strings.HasPrefix(str, "h"))
	fmt.Println(strings.HasSuffix(str, "ld!"))
	fmt.Println(strings.Join([]string{"hello", "beautiful", "world"}, "-")) // hello-beautiful-world
	fmt.Println(strings.Repeat(str, 2))
	fmt.Println(strings.Replace(str, "h", "H", 1)) // n为替换几次，替换几个 n为-1时没有限制
	fmt.Println(strings.Split("hello beautiful world !", " "))
	fmt.Println(strings.ToTitle(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(len(str))

	chinese := "你好"
	fmt.Println(len(chinese)) // 6
	// 中文字符串大小
}
