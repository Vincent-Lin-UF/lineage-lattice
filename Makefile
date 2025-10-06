APP=lltnode

EXT=
ifeq ($(OS),Windows_NT)
EXT=.exe
endif

BIN=$(APP)$(EXT)

.PHONY: build run test tidy

build:
	go build -o $(BIN) ./cmd/lltnode

run: 
	go run ./cmd/lltnode

test:
	go test ./...

tidy:
	go mod tidy