package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("PATH:\n", os.Getenv("PATH"))

	fmt.Println(os.Setenv("key", "value"))

	buff, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buff))
}
