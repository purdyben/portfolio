build: 
	
	GOARCH=wasm GOOS=js go build -o wasm/pong.wasm ./cmd/pong/main.go
	GOARCH=wasm GOOS=js go build -o wasm/snake.wasm ./cmd/snake/main.go
# 	go run server.go

serve:
	go run server.go
