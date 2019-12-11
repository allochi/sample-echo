default: test clean main.go
	go build -o bin/main .

test:
	go test

.PHONY: clean
clean:
	rm -f bin/*
