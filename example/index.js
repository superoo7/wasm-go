const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
  async result => {
    // go.run(result.instance);
    window.mod = result.module;
    window.inst = result.instance;
    await go.run(inst);
  }
);

global.fireEvent = value => {
  console.log(value);
};
