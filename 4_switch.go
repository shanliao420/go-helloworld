package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	condition := rand.Int() % 8
	fmt.Println("condition = ", condition)
	switch condition {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}

	now := time.Now()
	switch { // 在 switch 后面不加任何的变量，然后在 case 里面写条件分支 这样代码相比你用多个 if else 代码逻辑会更为清晰
	case now.Hour() <= 12:
		fmt.Println("good morning!")
	case now.Hour() >= 12 && now.Hour() <= 18:
		fmt.Println("good afternoon!")
	default:
		fmt.Println("good evening!")

	}

}
