.PHONY: precheckin test build run

precheckin: test build

clean:
	rm -rf target

target:
	mkdir target

build: target
	go build -o target/colluders main.go

test:
	go test ./...

run: build
	./target/colluders

