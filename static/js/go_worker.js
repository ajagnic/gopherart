self.importScripts('wasm_exec.js')

const go = new Go()

if (!WebAssembly.instantiateStreaming) {
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer()
    return await WebAssembly.instantiate(source, importObject)
  }
}
WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then(
  (result) => {
    go.run(result.instance)
  }
)

self.onmessage = function (e) {
  const img = e.data
  const img64 = global.processImage(img.b, img.len, img.opt)
  self.postMessage(img64)
}
