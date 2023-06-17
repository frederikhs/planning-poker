#!/bin/bash
set -e

docker compose --env-file .env.test run --build --rm web npm run lint
2