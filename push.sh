#!/bin/bash
set -e

CURRENT_VERSION=$(curl -s https://registry.hrgn.dk/v2/planning-poker-api/tags/list | jq '.tags | map(select(. | contains ("v"))) |join("\n")' --raw-output | sort -V - | tac | head -n 1)
NEXTVERSION=$(echo ${CURRENT_VERSION} | awk -F. -v OFS=. '{$NF += 1 ; print}')

echo "curren tag: ${CURRENT_VERSION}"
echo "next tag: ${NEXTVERSION}"

TAG=$NEXTVERSION docker compose -f docker-compose.yml -f docker-compose.prod.yml build
TAG=$NEXTVERSION docker compose -f docker-compose.yml -f docker-compose.prod.yml push
