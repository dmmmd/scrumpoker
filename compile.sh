#!/usr/bin/env bash
docker build -t scrumpoker-compiler -f docker/server/Dockerfile.compiler .
docker run --detach --name=scrumpoker-compiler scrumpoker-compiler
docker cp scrumpoker-compiler:/scrumpoker/scrumpoker docker/server/scrumpoker
docker container stop scrumpoker-compiler
docker container rm scrumpoker-compiler
docker image rm scrumpoker-compiler
