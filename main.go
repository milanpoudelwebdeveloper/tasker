package main

import "fmt"

func main() {
	store := NewStore()
	api := newAPIServer(":3000", nil)
	fmt.Println("Hello, World!")
}
