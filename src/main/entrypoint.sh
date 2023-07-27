#!/bin/sh
# @file entrypoint.sh
# @brief Entrypoint for Github Action.
#
# @description The script is intended to serve as the entrypoint for a Docker image running in an Alpine-based
# container. It acts as a bridge to the ``semver`` binary, which is the main functionality of the containerized
# application.
#
# CAUTION: Since the base image is Alpine, it doesn't have Bash available and uses Sell instead. So, the
# entrypoint must be specified as ``#!/bin/sh`` instead of ``#!/bin/bash``. running this script directly outside
# the Docker container could potentially cause trouble since it relies on ``#!/bin/sh`` as the shebang. If
# executed on a system that expects #!/bin/bash, it may lead to syntax errors or unexpected behavior.
#
# The purpose of this script is to provide an interface between the user and the semver binary inside the container.
# It allows users to pass arguments to the semver binary .
#
# === Script Arguments
#
# * *$@* (string): All arguments passed to the container. These arguments are all passed to the ``semver`` binary.
#
# === Script Example
#
# [source, bash]
# ```
# ./t.sh
# ```
#
# == Prerequisites
# Ensure that Docker is installed and on your system.


echo "----- " "$@"
echo "===== $1"
semver "$1"