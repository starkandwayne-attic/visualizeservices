Visualization of services
=========================

```
Nodes   Services
----------------
node 1: Re
node 2: rne
node 3: Cem
node 4: cnM
node 5: rcnm
```

In the example above, there is a Redis cluster (master `R` on node 1, slaves `r` on nodes 2 and 5), a Mongo cluster (primary `M` on node 4, secondaries on nodes 3 and 5), and so on (e.g. NATS is `n`, Elastic Search is `e`, Cassandra is `c`).

It is assumed that this application is run as a handler for a `consul watch` on any services changes:

```bash
consul watch -type=services visualizeservices
```

The input data from `consul watch` is something similar to:

```
{"consul":[],"consul-53":["udp"],"consul-8400":[],"consul-8500":[],"redis-1":["uuid-redis-8976","redis","master","uuid-redis-25973","slave","uuid-redis-20509"]}
```

The application can also fetch the services catalog for itself if only being run one time.

Run within boot2docker
----------------------

```
export DOCKER_IP=192.168.59.103
docker build -t starkandwayne/visualizeservices .
docker run -e "CONSUL_HTTP_ADDR=$DOCKER_IP:8500" starkandwayne/visualizeservices visualizeservices
```

Run watcher within boot2docker
------------------------------

```
./bin/visualization_start.sh
```
