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

RUN printf "export PS1='%s '\n" '\u:\W\$' >.profile
RUN printf "export PATH=/usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/bin:/sbin\n" >>.profile

CMD exec /bin/sh -l
