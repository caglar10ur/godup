all:
	go build dup.go
format:
	gofmt -s -w -tabs=false -tabwidth=4 dup.go
clean:
	rm -f dup
