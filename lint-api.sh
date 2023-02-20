#!/bin/bash
set -e

docker run -t --rm -v $(pwd)/api:/api -v ~/.cache/golangci-lint/v1.51.2:/root/.cache -w /api golangci/golangci-lint:v1.51.2 golangci-lint run -v
