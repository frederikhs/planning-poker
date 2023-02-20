#!/bin/bash
set -e

CURRENT_VERSION=$(curl -s https://registry.hrgn.dk/v2/planning-poker-api-new/tags/list | jq '.tags | sort | reverse | .[0]' --raw-output)
NEXTVERSION=$(echo ${CURRENT_VERSION} | awk -F. -v OFS=. '{$NF += 1 ; print}')

echo "next tag: ${NEXTVERSION}"

TAG=$NEXTVERSION docker compose -f docker-compose.yml -f docker-compose.prod.yml build
TAG=$NEXTVERSION docker compose -f docker-compose.yml -f docker-compose.prod.yml push
