Visualization of services
=========================

```
Nodes Services
----------------
1.    Re
2.    rne
3.    Cem
4.    cnM
5.    rcnm
```

In the example above, there is a Redis cluster (master `R` on node 1, slaves `r` on nodes 2 and 5), a Mongo cluster (primary `M` on node 4, secondaries on nodes 3 and 5), and so on (e.g. NATS is `n`, Elastic Search is `e`, Cassandra is `c`).

It is assumed that this application is run as a handler for a `consul watch` on any services changes:

```bash
consul watch  -type=services visualizeservices
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
docker run starkandwayne/visualizeservices -consul-addr $DOCKER_IP:8500
```
