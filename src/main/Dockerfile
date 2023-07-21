# The Dockerfile used for the Github Action. This image is published to DockerHub as well and can
# be used standalone. This Dockerfile bundles the Go application. The Go application serves as the
# core logic for managing and handling all tasks related to semantic versions.
#
# This is a Multistage Dockerfile that separates building the binary for Go application from
# running the binary in minimalistic linux environment.
# 
# For standalone usage of the Docker image, run ``docker run --rm local/semver:dev`` (showing the
# applications help) or ``docker run --rm local/semver:dev validate v.1.0`` to validate a version
# against semantic versioning rules.

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
FROM alpine:3.18.0 AS run
LABEL maintainer="sebastian@sommerfeld.io"

COPY --from=build app/semver /usr/bin/semver

ARG USER=semver
RUN adduser -D $USER

USER $USER
ENTRYPOINT [ "/usr/bin/semver" ]
