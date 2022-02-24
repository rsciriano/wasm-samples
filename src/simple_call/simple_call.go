package main

import (
	"fmt"
)

//export sum
func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, WebAssembly!")
}
