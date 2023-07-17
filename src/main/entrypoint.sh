#!/bin/bash
# @file entrypoint.sh
# @brief Lorem ipsum dolor sit amet, consetetur sadipscing elitr.
#
# @description Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore
# et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita
# kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
#
# IMPORTANT: Do not run this script directly. This script is intended to run as part of a Github Actions job.
#
# === Script Arguments
#
# * *$1* (string): The version to validate
#
# === Script Example
#
# [source, bash]
# ```
# ./entrypoint.sh
# ```


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


echo -e "[INFO] Run $0"
