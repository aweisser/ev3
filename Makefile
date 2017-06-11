build:
	go build ./cmd/ev3rest/ev3rest.go

tests:
	go test ./robot	
	go test ./goev3

default:
	build

.phony: build test