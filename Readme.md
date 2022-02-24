# WASM Samples

[Compiling a New C/C++ Module to WebAssembly - WebAssembly | MDN](https://developer.mozilla.org/en-US/docs/WebAssembly/C_to_wasm)

```
docker run --rm -v $(pwd):/src -u $(id -u):$(id -g) \
  emscripten/emsdk emcc helloworld.c -o helloworld.html --shell-file template.html
```

https://github.com/tinygo-org/tinygo/tree/release/src/examples/wasm

```
tinygo build -o ./simple_call.wasm -target wasm ./simple_call.go
tinygo build -o ./advanced_call.wasm -target wasm ./advanced_call.go
tinygo build -o ./dom_interop.wasm -target wasm ./dom_interop.go
```

Links:

- [Using WebAssembly | TinyGo](https://tinygo.org/docs/guides/webassembly/)
- [tinygo/src/examples/wasm at release · tinygo-org/tinygo](https://github.com/tinygo-org/tinygo/tree/release/src/examples/wasm)
- [WebAssembly using Go (Golang) | Run Go programs in the browser](https://golangbot.com/webassembly-using-go/)
- [js package - syscall/js - pkg.go.dev](https://pkg.go.dev/syscall/js#pkg-overview)
- [Passing strings between TinyGo and JavaScript — Blog documentation](https://www.alcarney.me/blog/2020/passing-strings-between-tinygo-wasm/)
- [syscall/js.finalizeRef not implemented · Issue #1140 · tinygo-org/tinygo](https://github.com/tinygo-org/tinygo/issues/1140)