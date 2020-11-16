LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

get-lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.32.2

lint:
	golangci-lint run ./...

test:
	go test ./... -count 1 -v

build:
	$(LINUX_AMD64) go build -o gorecipes .

# Usage: HOST=0.0.0.0 PORT=3000 make start
start: build
	./gorecipes

dk-build: build
	docker build -t gorecipes .

dk-start: dk-build
	docker run --rm -p 3000:3000 --env-file .env --name gorecipes gorecipes:latest

dk-logs:
	docker logs --tail 100 -f gorecipes
