GOROOT=$(shell go env GOROOT)

rebuild:
	@make wasm

wasm:
	GOOS=js GOARCH=wasm go build -mod vendor -ldflags="-s -w" -o static/wasm/whosonfirst_placetypes.wasm cmd/placetypes/main.go

example:
	go run -mod vendor cmd/example/main.go -port 8080
