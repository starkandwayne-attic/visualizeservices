#!/usr/bin/env bash


docker_image=starkandwayne/visualizeservices

host_ip=${DOCKER_IP:-192.168.59.103}
consul_ip=${CONSUL_IP:-$host_ip}
consul_port=${CONSUL_PORT:-8500}
consul_http_addr="${consul_ip}:${consul_port}"
kv_store=http://${consul_http_addr}/v1/kv

docker run -v /var/run/docker.sock:/var/run/docker.sock \
  -e "consul_http_addr=${consul_http_addr}" \
  ${docker_image} \
    consul watch -http-addr=$consul_http_addr \
      -type services \
      /scripts/handle_service_changes.sh
