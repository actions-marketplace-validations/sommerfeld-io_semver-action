#!/bin/bash
# @file semver.sh
# @brief Build, test and run the semver app locally.
#
# @description Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy
# eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero
# eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata
# sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing
# elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed
# diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd
# gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.
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

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


# Download and include logging library
rm -rf /tmp/bash-lib
mkdir -p /tmp/bash-lib
curl -sL https://raw.githubusercontent.com/sebastian-sommerfeld-io/jarvis/main/src/main/modules/bash-script/assets/lib/log.sh --output /tmp/bash-lib/log.sh
source /tmp/bash-lib/log.sh


readonly IMAGE="local/semver:dev"


# @description Wrapper function to encapsulate ``go`` in a Docker container (``go`` commands
# are delegated to the link:https://hub.docker.com/_/golang[golang] Docker image).
#
# The current working directory is mounted into the container and selected as working directory
# so all files are available to ``go``. Paths are preserved. The working directory is placed
# in ``$(pwd)`` (in the container) to make sure paths to the go app are the same everywhere (Go
# wrapper container, Dev Container and all images built from ``src/main/Dockerfile``). Keep in
# mind that most functions in this script (which call this ``go`` wrapper function) first ``cd``
# into the ``go`` folder. So most of the time the current working direktory is not ``src/main``
# (where this script is placed) but ``src/main/go``.
#
# The go wrapper container runs with the current user.
#
# @example
#    go version
#
# @arg $@ String The ``semver`` commands (1-n arguments) - $1 is mandatory
#
# @exitcode 8 If param with ``go`` command is missing
function go() {
  if [ -z "$1" ]; then
    LOG_ERROR "No command passed to the go container"
    LOG_ERROR "exit" && exit 8
  fi

  mkdir -p "/tmp/$USER/.cache"

  docker run --rm \
    --volume /etc/passwd:/etc/passwd:ro \
    --volume /etc/group:/etc/group:ro \
    --user "$(id -u):$(id -g)" \
    --volume "/tmp/$USER/.cache:/home/$USER/.cache" \
    --volume "$(pwd):$(pwd)" \
    --workdir "$(pwd)" \
    golang:1.20.6-alpine3.18 go "$@"
}


# @description Format go source code. Before formatting, the function ``cd``s into the
# ``go`` folder.
function format() {
  LOG_HEADER "Format code"
  (
    cd go || exit
    go fmt ./...
  )
}


# @description Run all test cases and security scanner. 
#
# Before testing, the function ``cd``s into the ``go`` folder.
function test() {
  LOG_HEADER "Run tests"
  (
    cd go || exit
    go test ./...
  )
}


# @description Build ``local/semver:dev`` Docker image.
function build() {
  LOG_HEADER "Build $IMAGE Docker image"
  docker build -t "$IMAGE" .
}

# @description Run ``semver`` app in Docker container.
#
# @arg $@ String The ``semver`` commands (1-n arguments) - $1 is mandatory
function run() {
  LOG_HEADER "Run app in Docker container" "$@"
  docker run --rm \
    "$IMAGE" "$@"
}


# @description Initialize the go application in needed. Before initializing, the function
# ``cd``s into the ``go`` folder.
function init() {
  (
    cd go || exit
    if [ ! -f go.mod ]; then
      local MODULE="github.com/sommerfeld-io/semver"
      readonly MODULE

      LOG_HEADER "Initialize $MODULE"
      go mod init "$MODULE"

      go get -u github.com/spf13/cobra@latest
      
      go mod tidy
    fi
  )
}


init
format
test
build
run "$@"
