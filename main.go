package main

import "fmt"

func main() {
	api := newAPIServer(":3000", nil)
	fmt.Println("Hello, World!")
}
