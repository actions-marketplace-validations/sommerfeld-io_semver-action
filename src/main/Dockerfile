FROM alpine:3.18.2 AS run
LABEL maintainer="sebastian@sommerfeld.io"

RUN apk add --no-cache bash=5.2.15-r0 \
    && apk add --no-cache curl=7.88.1-r0

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
