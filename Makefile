# nested-logrus-formatter

.PHONY: all
all: test demo

.PHONY: test
test:
	go test ./tests/* -v -count 1 -v

cover:
	go test ./tests/* -v -covermode=count -coverprofile=coverage.out

.PHONY: demo
demo:
	go run example/main.go
