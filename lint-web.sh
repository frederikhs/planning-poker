#!/bin/bash
set -e

docker compose run web npm run lint
