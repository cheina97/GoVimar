cli:
	go build -ldflags="-s -w"  ./cmd/govimar-cli
cli-install:
	go install -ldflags="-s -w" ./cmd/govimar-cli