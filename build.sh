#!/bin/bash

APP_NAME="shitkey"
VERSION="1.4"
ARCH="amd64"
MAINTAINER="Nathanne Isip <nathanneisip@gmail.com>"
DESCRIPTION="A secure command-line tool for file encryption and decryption."
INSTALL_PATH="/usr/local/bin"

BUILD_DIR="."
PACKAGE_ROOT="${APP_NAME}_${VERSION}_${ARCH}"
DEBIAN_DIR="${PACKAGE_ROOT}/DEBIAN"
INSTALL_BIN_DIR="${PACKAGE_ROOT}${INSTALL_PATH}"
DEB_FILE="${APP_NAME}_${VERSION}_${ARCH}.deb"

print_info() {
    echo -e "\x1b[36mINFO:\x1b[0m $1"
}

print_success() {
    echo -e "\x1b[32mSUCCESS:\x1b[0m $1"
}

print_error() {
    echo -e "\x1b[31mERROR:\x1b[0m $1"
    exit 1
}

print_info "Starting build and packaging process for ${APP_NAME}..."
print_info "Cleaning up previous build artifacts..."
rm -f "${APP_NAME}" "${DEB_FILE}"
rm -rf "${PACKAGE_ROOT}"

print_info "Compiling Go program into an ELF executable..."
go build -ldflags="-s -w"

if [ ! -f "${APP_NAME}" ]; then
    print_error "Compiled executable '${APP_NAME}' not found. Build failed."
fi

print_info "Stripping debug symbols from the executable..."
strip "${APP_NAME}" || print_error "Failed to strip executable!"

print_info "Creating Debian package directory structure..."
mkdir -p "${DEBIAN_DIR}" || print_error "Failed to create DEBIAN directory!"
mkdir -p "${INSTALL_BIN_DIR}" || print_error "Failed to create installation directory!"

print_info "Copying executable to package directory..."
cp "${APP_NAME}" "${INSTALL_BIN_DIR}/" || print_error "Failed to copy executable!"

chmod 0755 "${INSTALL_BIN_DIR}/${APP_NAME}" || print_error "Failed to set executable permissions!"

print_info "Creating DEBIAN/control file..."
cat <<EOF > "${DEBIAN_DIR}/control"
Package: ${APP_NAME}
Version: ${VERSION}
Architecture: ${ARCH}
Maintainer: ${MAINTAINER}
Description: ${DESCRIPTION}
EOF

print_info "Building the .deb package..."
dpkg-deb --build "${PACKAGE_ROOT}" || print_error "Failed to build .deb package!"

print_info "Cleaning up temporary package directory..."
rm -rf "${PACKAGE_ROOT}"

print_success "Debian package created successfully: ${DEB_FILE}"
print_info "You can install it using: sudo dpkg -i ${DEB_FILE}"

exit 0
