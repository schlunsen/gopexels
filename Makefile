all:
	go build -o gopexels main.go

test: 
	go test -v tests/*

