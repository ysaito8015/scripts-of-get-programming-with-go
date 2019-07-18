package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %s\n", name)
	fmt.Println("Hello, world")
	fmt.Println("Hello, 世界")
}
