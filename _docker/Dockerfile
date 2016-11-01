FROM golang:latest

RUN \
  apt-get update \
  && apt-get install -y --no-install-recommends pkg-config python-dev zip \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /tmp

ADD entrypoint /usr/local/bin/entrypoint

ENTRYPOINT ["/usr/local/bin/entrypoint"]