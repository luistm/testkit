test: ## Runs tests
	go test ./... -v

style: ## Formats the project
	goimports -w .
	go fmt ./...
