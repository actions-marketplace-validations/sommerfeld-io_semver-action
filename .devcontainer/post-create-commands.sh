#!/bin/bash
# @file post-create-commands.sh
# @brief This script configures the devlopment container and installs some development tools.
#
# @description This script serves as a ``postCreateCommand`` for the devcontainer, executing after the devcontainer
# has been created. Its purpose is to set up the necessary environment for development. First, the script installs
# link:https://github.com/sebastian-sommerfeld-io/jarvis[Jarvis], a productivity tool, to enhance the development
# workflow. After that, it adds custom bash aliases to simplify frequently used commands and enhance productivity.
#
# === Script Arguments
#
# The script does not accept any parameters.
#
# === Script Example
#
# The following snippets demonstrate how to use this script as a postCreateCommand in the ``.devcontainer.json``.
#
# [source, yaml]
# ```
# [...]
# "postCreateCommand": ".devcontainer/post-create-command.sh"
# [...]
# ```
#
# == Prerequisites
#
# The devcontainer plugin for VSCode is required to enable and manage development containers within the VSCode
# editor, providing a seamless and consistent development environment and ensure every developer is working with
# the same environment. 


set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace


echo "alias ll='ls -alF'" >> "$HOME/.bashrc"
curl https://raw.githubusercontent.com/sebastian-sommerfeld-io/jarvis/main/src/main/install.sh | bash -
