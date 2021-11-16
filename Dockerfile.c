FROM alpine:3.14.3

RUN addgroup -S unprivileged && \
	adduser -S unprivileged -G unprivileged

RUN  apk add \
	--no-cache \
	build-base

RUN mkdir /src

USER unprivileged

COPY main.c /src

WORKDIR /src

USER root
RUN gcc -o main main.c
USER unprivileged

CMD ["./main"]
