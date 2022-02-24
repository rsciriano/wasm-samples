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

func sayHello(this js.Value, args []js.Value) interface{} {

	//syscall/js.finalizeRef not implemented
	//https://github.com/tinygo-org/tinygo/issues/1140
	name := args[0].String()

	return js.ValueOf(hello(name))
}

func main() {
	fmt.Println("Hello, WebAssembly!")
	wait := make(chan struct{}, 0)
	js.Global().Set("wasm_sayHello", js.FuncOf(sayHello))
	<-wait
}
