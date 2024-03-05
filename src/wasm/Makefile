run: 
	GOARCH=wasm GOOS=js go build -o main.wasm ./cmd/snake/main.go
	go run main.go

build:
	go build -o bin/main main.go 
buildr: build
	./bin/main