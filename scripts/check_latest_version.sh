#!/bin/bash
set -e
set -o pipefail

# Configuration variables
HASHICORP_RELEASE_URL="https://api.github.com/repos/hashicorp/terraform/releases/latest"
TERRAFORM_MINOR_VERSION="0.9"
TERRAFORM_TESTED_VERSION="0.9.3"

# This function checks against the Github API url to find
# the latest version in the Hashicorp binary.
function get_latest_patch_version() {
    # Download the content of the JSON API from Github
    local content=$(curl -s -X GET ${HASHICORP_RELEASE_URL})

    # Then extract the data we need
    local version=$(echo ${content} | get_json_value "tag_name")

    # We also need to remove the "v" at the beginning
    # and print...
    version=$(echo ${version} | cut -c 2-)

    # Check if the new version matches the major.minor format
    local ispatch=$(compare_to_patch ${version})

    # Finally compare and execute the proper flow.
    if [[ ${ispatch}  == "true" ]]; then
        echo -e "${version}"
    else
        echo -e "${TERRAFORM_TESTED_VERSION}"
    fi
}

# Retrieve a JSON value from a given data. Data is
# $1 and the name of the JSON value to be retrieved is $2.
function get_json_value() {
    awk -F"[,:}]" '{for(i=1;i<=NF;i++){if($i~/'$1'\042/){print $(i+1)}}}' | tr -d '"' | sed -n ${2}p
}

# Compare to patch sees the difference
function compare_to_patch() {
    # We take the length to create a substring
    local len=$(( 10#${#TERRAFORM_MINOR_VERSION} ))

    # Get the substring
    local substr=${1:0:${len}}

    if [[ ${substr} ==  ${TERRAFORM_MINOR_VERSION} ]]; then
        echo -e "true"
    else
        echo -e "false"
    fi
}

# Run the actual flow
echo $(get_latest_patch_version)
