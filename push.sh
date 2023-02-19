#!/bin/bash
set -e

TAG=v0.0.2 docker compose -f docker-compose.yml -f docker-compose.prod.yml build
TAG=v0.0.2 docker compose -f docker-compose.yml -f docker-compose.prod.yml push
