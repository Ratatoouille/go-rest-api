.PHONY: clean create-binary-dir api-server

api-server:
	cd ./cmd && go build -v -o ../bin/api

create-binary-dir:
	mkdir -p bin

clean:
	rm bin/api

.PHONY: test

test: 
	go test -v -race -timeout 30s ./...
