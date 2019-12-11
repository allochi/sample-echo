default: clean main.go
	go build -o bin/main .

.PHONY: clean
clean:
	rm -f bin/*
