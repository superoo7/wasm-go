const path = require("path");
const CopyPlugin = require("copy-webpack-plugin");

module.exports = {
  entry: "./index.js",
  output: {
    path: path.resolve(__dirname, "dist"),
    filename: "index.js"
  },
  plugins: [
    new CopyPlugin([
      { from: "main.wasm", to: "dist/main.wasm" },
      {
        from: "wasm_exec.js",
        to: "dist/wasm_exec.js"
      }
    ])
  ],
  mode: "development"
};
