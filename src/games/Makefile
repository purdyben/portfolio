run: 
	GOARCH=wasm GOOS=js go build -o wasm/snake.wasm ./cmd/snake/main.go
	go run server.go

build:
	go build -o bin/server server.go 
buildr: build
	./bin/server