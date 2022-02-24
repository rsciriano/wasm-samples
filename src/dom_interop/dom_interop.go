package main

import (
	"fmt"
	"syscall/js"
)

func hello(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello from WASM, " + name + "!"

}

//export writeHello
func writeHello() {
	document := js.Global().Get("document")
	name := document.Call("getElementById", "nameInput").Get("value").String()
	document.Call("getElementById", "helloPanel").Set("value", hello(name))
}

func main() {
	fmt.Println("Hello, WebAssembly!")
}
