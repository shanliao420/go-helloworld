package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().Unix())
	secretNum := rand.Intn(maxNum)
	fmt.Println("please guss the num(0-100):")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error!")
			continue
		}
		input = strings.TrimSuffix(input, "\n") // 去掉换行符

		gussNum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input!")
			continue
		}

		if gussNum == secretNum {
			fmt.Println("you guss right!")
			break
		} else if gussNum < secretNum {
			fmt.Println("you guss smaller")
		} else if gussNum > secretNum {
			fmt.Println("you guss bigger")
		}
	}
}
