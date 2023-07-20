# ---------------------------------------------------------
# Build
# ---------------------------------------------------------
FROM golang:1.20-rc-alpine AS build
LABEL maintainer="sebastian@sommerfeld.io"

COPY app /app
WORKDIR /app

RUN go mod download \
    && go test ./... \
    && go build .

# ---------------------------------------------------------
# Run
# ---------------------------------------------------------
# FROM alpine:3.18.2 AS run
# LABEL maintainer="sebastian@sommerfeld.io"
# 
# COPY entrypoint.sh /entrypoint.sh
# ENTRYPOINT ["/entrypoint.sh"]

FROM alpine:3.18.0 AS run
LABEL maintainer="sebastian@sommerfeld.io"

COPY --from=build app/semver /usr/bin/semver

ARG USER=semver
RUN adduser -D $USER

USER $USER
ENTRYPOINT [ "/usr/bin/semver" ]
