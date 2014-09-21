#!/usr/bin/env bash

docker_image=starkandwayne/visualizeservices

if [[ -f Dockerfile ]]; then
  docker build -t ${docker_image} .
fi
