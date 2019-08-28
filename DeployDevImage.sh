#!/bin/sh
docker service rm dermaster_pikabu_signaling 2>/dev/null
docker stack deploy -c docker-stack.develop.yml dermaster
