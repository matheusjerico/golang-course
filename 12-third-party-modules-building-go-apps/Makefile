.PHONY: help
help: ## Show help menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-35s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: go-build
go-build: ## Building golang application
	go build -o bin/monsterslayer

.PHONY: go-run
go-run: ## Running golang application
	./bin/monsterslayer

