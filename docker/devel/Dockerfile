FROM jrmsdev/ucsh:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.06.04"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk --no-cache add go sudo

ENV GOPATH /root/go
ENV GOBIN /usr/local/bin

RUN apk add --no-cache git && \
	go get -v golang.org/x/tools/cmd/godoc && \
	go generate -v golang.org/x/tools/godoc/static && \
	go install -v -i golang.org/x/tools/cmd/godoc && \
	apk del --no-cache --purge git && \
	rm -rf /root/go /root/.cache

RUN printf '\n%%ucsh ALL=(ALL) NOPASSWD: ALL\n' >>/etc/sudoers

RUN mkdir -vp -m 0750 /go && mkdir -vp -m 0750 /go/src && \
	mkdir -vp -m 0750 /go/src/ucsh && chown -vR ucsh:ucsh /go

VOLUME /go/src/ucsh

ARG UCSH_UMASK

RUN printf 'umask %s\n' "${UCSH_UMASK}" >>/home/ucsh/.profile
RUN printf 'export GOPATH=/go\n' >>/home/ucsh/.profile

ADD --chown=ucsh entrypoint.sh /usr/local/bin/entrypoint.sh

USER ucsh:ucsh
WORKDIR /go/src/ucsh

ENV USER ucsh
ENV HOME /home/ucsh

ENV GOBIN ''
ENV GOPATH /go
RUN go env && go version

ENTRYPOINT /usr/local/bin/entrypoint.sh
