#!/bin/bash

set -x
set -o pipefail

# this flow-go sub folder will be where all the TPS tests will be run
# this will keep the TPS automation code separate from the code that's being tested so we won't run into issues
# of having old versions of automation code just because we happen to be testing an older version flow-go
git clone https://github.com/onflow/flow-go.git
cd flow-go/integration/localnet

git fetch
git fetch --tags

while read -r branch_hash; do
    hash="${branch_hash##*:}"
    branch="${branch_hash%%:*}"

    git checkout "$branch" || continue
    git reset --hard "$hash"  || continue

    git log --oneline | head -1
    git describe

    make -C ../.. crypto_setup_gopath

    # instead of running "make stop" which uses docker-compose for a lot of older versions,
    # we explicitly run the command here with "docker compose"
    DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker compose -f docker-compose.metrics.yml -f docker-compose.nodes.yml down -v --remove-orphans

    make clean-data
    make -e COLLECTION=12 VERIFICATION=12 NCLUSTERS=12 LOGLEVEL=INFO bootstrap

    DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker compose -f docker-compose.nodes.yml build || continue
    DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker compose -f docker-compose.nodes.yml up -d || continue

    sleep 30;
    go run -tags relic ../benchmark/cmd/ci -log-level debug -git-repo-path ../../ -tps-initial 800 -tps-min 1 -tps-max 1200 -duration 30m

    # instead of running "make stop" which uses docker-compose for a lot of older versions,
    # we explicitly run the command here with "docker compose"
    DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker compose -f docker-compose.metrics.yml -f docker-compose.nodes.yml down -v --remove-orphans

    docker system prune -a -f
    make clean-data
done </opt/master.recent
