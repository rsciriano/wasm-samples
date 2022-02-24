'use strict';

var wasm;

function wasm_init(url) {

    const go = new Go();
    if ('instantiateStreaming' in WebAssembly) {
        WebAssembly.instantiateStreaming(fetch(url), go.importObject).then(function (obj) {
            wasm = obj.instance;
            go.run(wasm);
        })
    } else {
        fetch(WASM_URL).then(resp =>
            resp.arrayBuffer()
        ).then(bytes =>
            WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
                wasm = obj.instance;
                go.run(wasm);
            })
        )
    }
}
