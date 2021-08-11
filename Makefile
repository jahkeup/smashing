build: smashing

smashing: $(shell find . -name '*.go')
	go build . 

test:
	go test $(V) ./...

fmt:
	goimports -local github.com/jahkeup/smashing -w .
