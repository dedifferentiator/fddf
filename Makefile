CC=go
BINNAME=bin/fddf

build:
	$(CC) build -o $(BINNAME) main.go

clean:
	rm $(BINNAME)

run:
	$(CC) run main.go

test:
	$(CC) test

fmt:
	$(CC) fmt ./...

lint:
	golint . internal/...

all: build fmt lint test
