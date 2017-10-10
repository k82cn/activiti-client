all: init
	go build -o _output/acli main.go

init:
	mkdir -p _output
