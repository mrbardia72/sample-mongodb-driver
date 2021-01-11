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
	gor datafaker/data.go
