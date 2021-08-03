#!/bin/bash

# text color
RED='\033[0;31m'   # RED
BLUE='\033[0;34m'  # Blue

# Underline color
UCyan='\033[4;36m' # Cyan

# shellcheck disable=SC2059
printf "${BLUE}Generating swagger file into ${UCyan}\"./docs\"${BLUE} folder\n${RED}${UCyan}" 1>&2
swag init --parseDependency --dir ./ --output ./docs -g ./main.go