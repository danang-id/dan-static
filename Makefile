build:
	@make build-darwin
	@make build-linux
	@make build-windows

build-darwin:
	@GOOS=darwin GOARCH=amd64 go build -o bin/dan-static-darwin-amd64

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/dan-static-linux-amd64

build-windows:
	@GOOS=windows GOARCH=amd64 go build -o bin/dan-static-windows-amd64.exe

run:
	@go run server.go

clean:
	@npx rimraf bin