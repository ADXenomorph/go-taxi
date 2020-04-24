.PHONY: build test bench bench-race apache-bench wrk cover cover-html help godoc

TMP_COVER:=$(shell mktemp)

build: ## Build the binary
	go build ./cmd/taxid/taxid.go

test: ## Run tests
	go clean -testcache && go test ./... -v

bench: ## Runs parallel benchmark
	go test -benchmem -bench=. -cpu=1,2,3,4 ./...

bench-race: ## Runs parallel benchmark with race detector
	go test -benchmem -bench=. -race -cpu=1,2,3,4 ./...

apache-bench: build ## Runs apache bench
	./taxid & 
	sleep 0.5 
	ab -n 50000 -c 1000 localhost:8080/request
	pkill taxid

wrk: build ## Runs wrk bench
	./taxid &
	sleep 0.5
	wrk -t 4 -c 16 -d 10 http://localhost:8080/request
	pkill taxid

cover: ## Show coverage in CLI
	go test ./... -coverprofile cover.out \
	&& go tool cover -func cover.out \
	&& rm cover.out

cover-html: ## Show coverage in browser
	go test -coverprofile=${TMP_COVER} ./... && go tool cover -html=${TMP_COVER} && unlink ${TMP_COVER}

godoc: ## Starts godoc server
	godoc -http=:6060 &
	xdg-open http://localhost:6060/pkg/github.com/ADXenomorph/go-taxi/?m=all \
	&& sleep 0.1 \
	&& echo "Dont forget to stop the godoc server: pkill godoc"

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

