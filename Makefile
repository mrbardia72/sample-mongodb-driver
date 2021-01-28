LOGFILE=$(LOGPATH) `date +'%A-%b-%d-%Y-%H-%M-%S'`

.PHONY: help
help: ## help command
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: cm
cm: ## push to github
	@echo '************👇  run command 👇************'
	git add .
	git commit -m "🌱#gopher#💙-${LOGFILE}"
	git push -u origin main

.PHONY: fakeruser
fakeruser: ## run data facker
	gor internal/datafaker/data.go

.PHONY: run
run: ## run data facker
	go run cmd/main.go

.PHONY: config
config: ## You can store your credentials using the following command
	git config credential.helper store
	git push http://github.com/your-user-name/your-epo.git
