#!/bin/bash
set -e

docker compose run web sh -c "npm ci && npm run lint"
