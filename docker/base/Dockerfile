FROM alpine:3

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.05.31"

USER root:root
WORKDIR /root

RUN apk --no-cache -U upgrade

ARG UCSH_UID
ARG UCSH_GID

RUN addgroup -g ${UCSH_GID} ucsh && \
	adduser -h /home/ucsh -g ucsh -s /bin/sh -G ucsh -D -u ${UCSH_UID} ucsh

USER ucsh:ucsh
WORKDIR /home/ucsh

ENV USER ucsh
ENV HOME /home/ucsh

RUN printf "PS1='%s '\n" '\u:\W\$' >.profile

CMD exec /bin/sh -l
