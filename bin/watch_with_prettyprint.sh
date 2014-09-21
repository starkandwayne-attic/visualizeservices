#!/usr/bin/env bash

docker_image=starkandwayne/visualizeservices

host_ip=${DOCKER_IP:-192.168.59.103}
consul_ip=${CONSUL_IP:-$host_ip}
consul_port=${CONSUL_PORT:-8500}
CONSUL_HTTP_ADDR=${CONSUL_HTTP_ADDR:-${consul_ip}:${consul_port}}
kv_store=http://${CONSUL_HTTP_ADDR}/v1/kv

docker run -v /var/run/docker.sock:/var/run/docker.sock \
  -e "CONSUL_HTTP_ADDR=${CONSUL_HTTP_ADDR}" \
  ${docker_image} \
    consul watch -http-addr=$CONSUL_HTTP_ADDR \
      -type services visualizeservices pretty
