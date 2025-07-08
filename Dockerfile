FROM ubuntu:jammy-20240911.1

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Jakarta

WORKDIR /opt

RUN apt-get update -qq \
    && apt-get upgrade -y -qq \
    && apt-get install -y --no-install-recommends -qq \
        wget ca-certificates make curl git \
    && rm -rf /var/lib/apt/lists/*

ENV GOPROXY=https://proxy.golang.org
ENV GOROOT=/usr/local/go
ENV GOPATH=/usr/local/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN wget https://golang.org/dl/go1.23.2.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz \
    && rm -f go1.23.2.linux-amd64.tar.gz \
    && echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile.d/go.sh \
    && chmod +x /etc/profile.d/go.sh

WORKDIR /app

ENV GO111MODULE=on

WORKDIR /app

ADD cmd /app/cmd
ADD internal /app/internal
ADD pkg /app/pkg
ADD config.json /app/config.json
ADD go.mod /app/go.mod
ADD go.sum /app/go.sum
ADD Makefile /app/Makefile

RUN make build