#!/bin/bash
# @file semver.sh
# @brief Build, test and run the semver app and Docker image locally.
#
# @description This script is designed to build and test the Semver Docker image using Docker Compose. It
# starts a Docker Compose Stack that performs a series of essential steps, including initializing Go modules,
# running Go tests, building the binary from the Go app, and creating the Docker image with the binary.
# Additionally, the script executes tests within several containers to ensure the functionality of the Docker
# image. 
#
# === Script Arguments
#
# * *$@* (...): The same parameters as ``semver`` app.
#
# === Script Example
#
# [source, bash]
# ```
# ./semver.sh --help
# ```
#
# == Prerequisites
#
# The only prerequisite is a working a Docker installation. Docker compose is shipped alongside Docker. It
# has been developed and tested with Docker version 24.0.2 on top of Ubuntu 22.10. With Docker installed,
# users can execute the script to build and test the Semver Docker image without any additional dependencies.

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


MY_USERNAME="$(whoami)" MY_UID="$(id -u)" MY_GID="$(id -g)" docker compose down -v
MY_USERNAME="$(whoami)" MY_UID="$(id -u)" MY_GID="$(id -g)" docker compose up --build --force-recreate --remove-orphans
