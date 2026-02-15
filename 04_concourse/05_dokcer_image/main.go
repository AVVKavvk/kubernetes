package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello world from image")

	count := 0

	for true && count <= 100 {
		fmt.Println("Count: ", count)
		count++
		time.Sleep(2 * time.Second)
	}
}
