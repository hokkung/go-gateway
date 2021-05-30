FROM golang:1.16 AS builder
COPY . /build
WORKDIR /build
ARG CI_JOB_TOKEN
RUN GOOS=linux \
	BUILDFLAGS="-v -o go-gateway -ldflags \
	make build

FROM debian:buster
RUN apt-get update \
	&& rm -rf /var/lib/apt/lists/* \
	&& useradd -d /app -M -u 1000 -s /bin/false app
COPY docker-entrypoint.sh /docker-entrypoint.sh
COPY --from=builder /build/go-gateway /app/go-gateway
USER 1000
ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["/docker-entrypoint.sh"]
