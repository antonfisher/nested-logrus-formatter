# nested-logrus-formatter

.PHONY: all
all: test demo

.PHONY: test
test:
	go test ./tests/* -v -count 1

.PHONY: demo
demo:
	go run example/main.go
