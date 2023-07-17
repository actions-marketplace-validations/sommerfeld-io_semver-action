FROM alpine:3.18.2 AS run
LABEL maintainer="sebastian@sommerfeld.io"

RUN apk add --no-cache bash=5.2.15-r5 \
    && apk add --no-cache curl=8.1.2-r0

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
