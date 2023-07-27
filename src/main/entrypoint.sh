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
# * *$1* (string): The ``semver`` command
# * *$2* (string): The arguments to pass to the ``semver`` command
# * *$2* (string): Flags to pass to the ``semver`` command (optional)
#
# === Script Example
#
# [source, bash]
# ```
# ./entrypoint.sh validate v0.1.0
# ./entrypoint.sh validate v0.1.0 --json
# ```
#
# == Prerequisites
# Ensure that Docker is installed and on your system.


readonly CMD="$1"
readonly ARGS="$2"
readonly FLAGS="$3"


# Check if $3 is NULL
if [ -z "$FLAGS" ]; then
    semver "$CMD" "$ARGS"
else
    semver "$CMD" "$ARGS" "$FLAGS"
fi
