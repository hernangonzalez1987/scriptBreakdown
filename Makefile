BINARY_NAME=script-breakdown

build: test
	go mod download
	go build -o ${BINARY_NAME} cmd/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test ./... -cover

lint:
	golangci-lint run --enable-all --disable depguard --disable testpackage  --disable exhaustruct --disable nilnil --exclude-files "mock.*" --sort-results




