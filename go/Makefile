update_version:
	node ../bin/update_version.js go.mod 'github.com/bloock/bloock-sdk-go/v2 v(.*)' $(shell cat ../VERSION)
	go mod tidy

run:
	find samples/**/* -type f -name '*.go' | rush 'echo {}; go run {}' 
