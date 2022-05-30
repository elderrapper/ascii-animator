build:
	go build -v -o ./bin/ascii-animator github.com/davidhsingyuchen/ascii-animator/cmd

run: build
	./bin/ascii-animator

.PHONY: build run
