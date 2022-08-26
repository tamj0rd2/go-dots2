include .bingo/Variables.mk

.PHONY: test
.DEFAULT_GOAL := test

setup:
	git config core.hooksPath .hooks
	go install github.com/bwplotka/bingo@latest
	bingo get

t: test
test: lint
	$(GOTEST) --race --count=1 ./...

ci:
	git pull -r
	make test
	git push

lint:
	$(GOLANGCI_LINT) run --timeout=5m ./...

lf: lintfix
lintfix:
	@$(GOLANGCI_LINT) run ./... --fix
	@$(GCI) write -s standard -s default -s "prefix(github.com/tamj0rd2/go-dots2)" $$(find . -type f -name '*.go' -not -path "./vendor")

generate:
	. .bingo/variables.env
	@go generate ./...
