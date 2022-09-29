SHELL := /bin/bash

.PHONY: dev-jq
dev-jq:
	source .env && cd cmd && go run . 2>&1 | jq

.PHONY: dev
dev:
	source .env && cd cmd && go run .