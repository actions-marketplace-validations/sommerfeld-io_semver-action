# @file Dockerfile
# @brief Build and run the semver application.
#
# @description This Dockerfile is used for the Github Action. This image is published to
# link:https://hub.docker.com/r/sommerfeldio/semver[DockerHub] as well and can be used standalone.
# This Dockerfile bundles the Go application. The Go application serves as the core logic for
# managing and handling all tasks related to semantic versions.
#
# This is a Multistage Dockerfile that separates building the binary for Go application from
# running the binary in minimalistic linux environment.
#
# === Environment Variables
#
# Inputs, that are defined in the ``action.yml``, are available as environment variables.
# Environment variables are prefixed with ``INPUT``. So the input param ``command`` becomes
# the environment variable ``$INPUT_COMMAND``.
#
# * *INPUT_COMMAND* (string): This is the command string passed from the Github Action. Can be
# empty. When running this image inside a Github Action workflow this param probably won't be
# empty because otherwise only the applications help is printed and no validations take place.
# When running this image standalone, this param is empty and allows you to pass your own commands
# through the command line.
#
# === Example
#
# For standalone usage of the Docker image, run ``docker run --rm local/semver:dev`` (showing the
# applications help) or ``docker run --rm local/semver:dev validate v.1.0`` to validate a version
# against semantic versioning rules.
#
# To build and run the image for testng purposes locally use ``docker build -t local/semver:dev .``
# and ``docker run --rm local/semver:dev``.
#
# == Prerequisites
#
# This image has been developed and tested with Docker version 24.0.2 on top of Ubuntu 22.10. 
#
# == Stage 1: ``build``
#
# In the ``build`` stage, the application's source code is copied into the Docker image, and the
# binary is then built from the Go sources. The link:https://hub.docker.com/_/golang[``golang``] 
# image is utilized, which is based on Alpine Linux and comes pre-configured with all the necessary
# tools for compiling Go applications. Nno extra software is installed.
#
# == Stage 2: ``run``
#
# In the ``run`` stage, the ``semver`` binary is copied from the ``build`` stage and serves as the
# entry point for the Docker image. The base image used for this stage is link:https://hub.docker.com/_/alpine[Alpine].
#
# == See also
#
# * For the compose config, that orchestrates the local build process, see xref:AUTO-GENERATED:docker-docs/src/main/docker-compose-yml-docker-docs.adoc[``src/main/docker-compose.yml``]
# * To run this Compose config, use xref:AUTO-GENERATED:bash-docs/src/main/semver-sh.adoc[``src/main/semver.sh``] bash script.


FROM golang:1.20-rc-alpine AS build
LABEL maintainer="sebastian@sommerfeld.io"

COPY app /app
WORKDIR /app

RUN go mod download \
    && go test ./... \
    && go build .


FROM alpine:3.18.2 AS run
LABEL maintainer="sebastian@sommerfeld.io"

COPY --from=build app/semver /usr/bin/semver

ARG USER=semver
RUN adduser -D $USER

USER $USER
ENTRYPOINT [ "/usr/bin/semver $INPUT_COMMAND" ]
