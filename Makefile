SHELL := /bin/bash

.PHONY: dev
dev:
	cd cmd && go run . 2>&1 | jq