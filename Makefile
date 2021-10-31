
deploy:
	go build -o $(GOPATH)/bin/awg

generate:
	go generate -v ./example