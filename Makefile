all:
	go build -o gopexels main.go

install:
	cp gopexels $$GOPATH/bin

test: 
	go test -v tests/*

