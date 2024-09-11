.DEFAULT_GOAL=run

MAIN_FILE := ./cmd/app/main.go

name ?= default

run:
	go run $(MAIN_FILE)

run-flag:
	go run $(MAIN_FILE) -name $(name)

test:
	go test ./... -v

build:
	go build -o ./gosearch $(MAIN_FILE)

# make run-flag name=golang