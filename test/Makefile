LOGFILE=$(LOGPATH) `date +'%A-%b-%d-%Y-%H-%M-%S'`

.PHONY: hp
hp: ## help command
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: tpass
tpass: ## run test pass
	go test