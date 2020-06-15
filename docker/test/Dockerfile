FROM jrmsdev/ucsh:devel

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.05.31"

USER root:root
WORKDIR /root

ADD --chown=ucsh entrypoint.sh /usr/local/bin/entrypoint.sh

USER ucsh:ucsh
WORKDIR /go/src/ucsh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
