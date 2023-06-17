#!/bin/bash
set -e

docker compose run web bash -c "npm ci && npm run lint"
