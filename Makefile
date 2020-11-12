LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

lint:
	golangci-lint run ./...

test:
	go test ./... -count 1 -v

build:
	$(LINUX_AMD64) go build -o gorecipes .

dk-build: build
	docker build -t gorecipes .

dk-start:
	docker run --name gorecipes gorecipes:latest

dk-logs:
	docker logs --tail 100 -f gorecipes