FROM mischief/docker-golang
MAINTAINER Dr Nic Williams <drnicwilliams@gmail.com>

RUN apt-get install zip -y

# ADD https://dl.bintray.com/mitchellh/consul/0.4.0_linux_amd64.zip /tmp/consul.zip
ADD blobs/consul-0.4.0_linux_amd64.zip /tmp/
RUN cd /tmp; unzip consul-0.4.0_linux_amd64.zip; mv consul /bin/

# ADD https://get.docker.io/builds/Linux/x86_64/docker-1.2.0 /bin/docker
ADD blobs/docker-1.2.0 /bin/docker
RUN chmod +x /bin/docker

ADD . /root/go/src/github.com/starkandwayne/visualizeservices
RUN cd /root/go/src/github.com/starkandwayne/visualizeservices; go get ./...
