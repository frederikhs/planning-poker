#!/bin/bash
set -e

TAG=v0.0.3 docker compose -f docker-compose.yml -f docker-compose.prod.yml build
TAG=v0.0.3 docker compose -f docker-compose.yml -f docker-compose.prod.yml push
