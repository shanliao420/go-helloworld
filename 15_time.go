package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2023-05-12 16:23:06.651363 +0800 CST m=+0.000192035
	myTime := time.Date(2023, 5, 12, 16, 23, 22, 0, time.UTC)
	fmt.Println(myTime)
	fmt.Println(myTime.YearDay())
	fmt.Println(myTime.Year(), myTime.Month(), myTime.Day())
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 使用2006-01-02 15:04:05格式化而不是yyyy-MM-dd
	diff := now.Sub(myTime)
	fmt.Println(diff)

	parse, err := time.Parse("2006-01-02 15:04:05", "2023-05-12 16:23:22")
	if err != nil {
		panic(err)
	}
	fmt.Println(parse)
	fmt.Println(parse == myTime)
	fmt.Println(now.Unix()) // 1683880427

}
