BINARY_NAME=script-breakdown

build:
	go mod download
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME} start

clean:
	go clean
	rm ${BINARY_NAME}

test: 
	go test --cover ./... 

lint:
	golangci-lint run --enable-all --disable depguard --disable testpackage --disable godot --disable gochecknoinits --disable exhaustruct --disable nilnil --disable gochecknoglobals --exclude-files "mock.*" --sort-results




