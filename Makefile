
init:
	@cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./example


build:
	@GOOS=js GOARCH=wasm go build -o example/main.wasm src/main.go 

serve:
	@cd example && npm run serve