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


MY_USERNAME="$(whoami)" MY_UID="$(id -u)" MY_GID="$(id -g)" docker compose down -v
MY_USERNAME="$(whoami)" MY_UID="$(id -u)" MY_GID="$(id -g)" docker compose up --build --force-recreate --remove-orphans
