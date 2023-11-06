
.PHONY: examples
examples: clean build
	./build/mlogplus-linux-$(shell go env GOARCH) ./examples

build: 
	mkdir -p build
	env GOOS=darwin GOARCH=arm64 go build -o build/mlogplus-darwin-arm64 ./cmd/mlogplus/ 
	env GOOS=darwin GOARCH=amd64 go build -o build/mlogplus-darwin-amd64 ./cmd/mlogplus/ 
	env GOOS=linux GOARCH=amd64 go build -o build/mlogplus-linux-amd64 ./cmd/mlogplus/ 
	env GOOS=linux GOARCH=arm64 go build -o build/mlogplus-linux-arm64 ./cmd/mlogplus/ 
	env GOOS=windows GOARCH=amd64 go build -o build/mlogplus-win-amd64.exe ./cmd/mlogplus/ 

clean:
	rm -rf build examples/compiled