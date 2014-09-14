#!/usr/bin/env bash

host_ip=${DOCKER_IP:-192.168.59.103}
consul_ip=${CONSUL_IP:-$host_ip}
consul_port=${CONSUL_PORT:-8500}
consul_http_addr=${consul_http_addr:-${consul_ip}:${consul_port}}

visualizeservices -consul-addr $consul_http_addr
