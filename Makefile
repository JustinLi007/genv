target = bin/genv
main = ./cmd/genv
src = $(shell find . -name '*.go')

default: $(target)

$(target): $(src) | bin
	go build -o $@ $(main)

bin:
	mkdir -p $@

run: $(target)
	./$(target)

install:
	go install $(main)

test:
	go clean -testcache
	go test ./... $(ARGS)

tidy:
	go mod tidy

clean:
	rm -f $(target)

.PHONY: clean default run test tidy
