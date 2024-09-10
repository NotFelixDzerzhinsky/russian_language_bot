.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot bot/main.go

run: build
	./.bin/bot