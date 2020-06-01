FROM jrmsdev/ucsh:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.05.31"

USER root:root
WORKDIR /root

RUN apk --no-cache add go

RUN mkdir -vp -m 0750 /go && mkdir -vp -m 0750 /go/src && \
	mkdir -vp -m 0750 /go/src/ucsh && chown -vR ucsh:ucsh /go
VOLUME /go/src/ucsh

ARG UCSH_UMASK

RUN printf 'umask %s\n' "${UCSH_UMASK}" >>/home/ucsh/.profile
RUN printf 'export GOPATH=/go\n' >>/home/ucsh/.profile

USER ucsh:ucsh
WORKDIR /go/src/ucsh

ENV GOPATH /go
RUN go env && go version

CMD exec /bin/sh -l
