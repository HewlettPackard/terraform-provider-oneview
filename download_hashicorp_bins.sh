#!/bin/bash
set -e
set -o pipefail

TERRAFORM_VERSION=0.8.8
REQUIRED_APPS=("curl" "unzip")

function check_installed_apps() {
    # Define a list of apps as requirements, that we can later check, if they're
    # not installed, we can install them
    local not_installed_apps=()

    # Check if any of the given programs is not installed, as requirements
    for i in "${REQUIRED_APPS[@]}"
    do
        # Check if the app is not installed, if not, install
        if ! [ -x "$(command -v ${i})" ]; then
            not_installed_apps=("${not_installed_apps[@]}" "${i}")
        fi
    done

    # Check if we need to install extra dependencies
    if ! [ ${#not_installed_apps[@]} -eq 0 ]; then
        echo -e "The following apps are a requirement and they aren't installed."
        echo -e "	Please install them first: ${not_installed_apps[@]}"
    fi
}

# Download terraform binaries for different platforms
# to an specific folder.
# $1: platform
# $2: output folder
function download_terraform_binary() {
    # Define the platform, the zip file name, and the path to the binary
    local platform=$1
    local destination=$2
    local filename="terraform_${TERRAFORM_VERSION}_${platform}"
    local zipfile="${filename}_amd64.zip"
    local url_path="https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/${zipfile}"

    # We also need some temporary folders to store the temporary binary
    local random_uid="$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)"
    local temp_folder="/tmp/terraform_download_${random_uid}"

    # Create the directory we need
    mkdir -p ${temp_folder}

    # Now download the file into the destination with ${zipfile}
    # as the name
    echo -e "Downloading Terraform into ${temp_folder}/${zipfile}..."
    curl -sS -o "${temp_folder}/${zipfile}" "${url_path}"
    [ $? -eq 0 ] || { echo -e "Unable to download Terraform zip file. Exiting..."; exit 1; }

    # Once we downloaded it, we will extract it
    unzip -q "${temp_folder}/${zipfile}" -d "${temp_folder}/"

    # Once we have it, let's move it to a folder with the OS Architecture
    # First, we need to know the binary filename
    local bin_filename="terraform"
    if [ "$platform" == "windows" ]; then
        bin_filename="terraform.exe"
    elif [ "$platform" == "darwin" ]; then
        platform="macos"
    fi

    # Now create the output directory (which will fail silently if the directory already exists)
    mkdir -p "${destination}/${platform}"

    # Now we move the file to that destination
    mv "${temp_folder}/${bin_filename}" "${destination}/${platform}/"
    echo -e "Binary file saved at \"${destination}/${platform}/${bin_filename}\""

    # And then we remove the temporary directory
    rm -rf "${temp_folder}"

    # Line break
    echo -e ""
}

# - file $(find terraform/ -type f -name "*" | sed "s/^\.\///")
function start() {
    local platforms=("windows" "darwin" "linux")
    local outputdir="$1"

    # First, remove any previous directory so we have a clean build
    if [ -d "$outputdir" ]; then
        rm -rf "${outputdir}"
    fi

    # Then download all the required binaries to "./bin"
    for i in "${platforms[@]}"; do
        download_terraform_binary "${i}" "${outputdir}"
    done

    # Print an output to see if it was successful
    echo -e "Downloaded Terraform binaries for: ${platforms[@]}"
    echo -e "-----"
    file $(find "${outputdir}" -type f -name "*" | sed "s/^\/\///")
}

start $1
