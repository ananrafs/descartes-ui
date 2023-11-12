build-fe:
	cd ./frontend && yarn build && cd ..

build: build-fe
	@go run main.go

.PHONY: build build-fe
.DEFAULT_GOAL:= build