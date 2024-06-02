package main

import (
	"farmish/config"
	"fmt"
)

func main() {
	ids := []int{}
	num, err := config.GenNewID(ids)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}
