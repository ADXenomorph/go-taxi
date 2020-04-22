build: ## Build the binary
	go build -race ./cmd/taxid/taxid.go

bench: ## Runs parallel benchmark
	go test -bench=. -cpu=1,2,3,4 ./cmd/taxid

bench-race: ## Runs parallel benchmark with race detector
	go test -bench=. -race -cpu=1,2,3,4 ./cmd/taxid

apache-bench: build ## Runs apache bench
	./taxid & 
	sleep 0.5 
	ab -n 30000 -c 300 localhost:8080/request
	pkill taxid

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

