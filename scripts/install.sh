#!/bin/bash
set -e

# Mere Installation Script
# This script downloads and installs the Mere CLI tool

VERSION="${VERSION:-latest}"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
REPO="Enriquefft/Mere"
GITHUB_BASE="https://github.com/${REPO}/releases"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Print colored output
info() {
    echo -e "${GREEN}[INFO]${NC} $1" >&2
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1" >&2
}

error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
}

# Detect OS
detect_os() {
    local os
    case "$(uname -s)" in
        Linux*)     os=linux;;
        Darwin*)    os=darwin;;
        MINGW*|MSYS*|CYGWIN*) os=windows;;
        *)          error "Unsupported OS: $(uname -s)"; exit 1;;
    esac
    echo "$os"
}

# Detect architecture
detect_arch() {
    local arch
    case "$(uname -m)" in
        x86_64|amd64)    arch=amd64;;
        arm64|aarch64)    arch=arm64;;
        i386|i686)       arch=386;;
        armv6l|armv7l)   arch=arm;;
        *)               error "Unsupported architecture: $(uname -m)"; exit 1;;
    esac
    echo "$arch"
}

# Get latest version from GitHub
get_latest_version() {
    if [ "$VERSION" != "latest" ]; then
        echo "$VERSION"
        return
    fi

    info "Fetching latest version from GitHub..."
    local latest_version
    latest_version=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

    if [ -z "$latest_version" ]; then
        error "Failed to fetch latest version"
        exit 1
    fi

    echo "$latest_version"
}

# Download the binary
download_binary() {
    local version="$1"
    local os="$2"
    local arch="$3"
    local filename="mere-${os}-${arch}.tar.gz"
    local download_url="${GITHUB_BASE}/download/${version}/${filename}"

    info "Downloading Mere ${version} for ${os}/${arch}..."
    info "From: ${download_url}"

    # Create temp directory for download
    local tmp_dir=$(mktemp -d)
    local tmp_archive="${tmp_dir}/${filename}"
    local tmp_file="${tmp_dir}/mere"

    # Download archive
    if ! curl -fsSL "$download_url" -o "$tmp_archive"; then
        error "Failed to download Mere from ${download_url}"
        exit 1
    fi

    # Extract binary from archive
    tar -xzf "$tmp_archive" -C "$tmp_dir" mere

    # Make binary executable
    chmod +x "$tmp_file"

    echo "$tmp_file"
}

# Install the binary
install_binary() {
    local source="$1"
    local target="${INSTALL_DIR}/mere"

    # Check if install directory exists
    if [ ! -d "$INSTALL_DIR" ]; then
        warn "Install directory ${INSTALL_DIR} does not exist. Creating it..."
        sudo mkdir -p "$INSTALL_DIR"
    fi

    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        info "Installing Mere to ${target}..."
        cp "$source" "$target"
    else
        info "Installing Mere to ${target} (requires sudo)..."
        sudo cp "$source" "$target"
    fi

    # Make sure it's executable
    sudo chmod +x "$target"

    # Clean up temp file
    rm "$source"
    rmdir "$(dirname "$source")"
}

# Check PATH
check_path() {
    if ! echo "$PATH" | grep -q "${INSTALL_DIR}"; then
        warn "${INSTALL_DIR} is not in your PATH"
        warn "Add the following to your shell configuration:"
        warn "  export PATH=\"\$PATH:${INSTALL_DIR}\""
    fi
}

# Main installation
main() {
    echo ""
    echo "Mere Installation Script"
    echo "========================"
    echo ""

    # Detect OS and architecture
    local os=$(detect_os)
    local arch=$(detect_arch)
    info "Detected OS: ${os}"
    info "Detected Architecture: ${arch}"
    echo ""

    # Get version
    local version
    version=$(get_latest_version) || exit 1
    info "Installing version: ${version}"
    echo ""

    # Download binary
    local binary_path
    binary_path=$(download_binary "$version" "$os" "$arch") || exit 1

    # Install binary
    install_binary "$binary_path"

    echo ""
    info "Mere installed successfully!"
    echo ""
    info "Version: $(mere version 2>/dev/null || echo 'unknown')"
    echo ""

    # Check PATH
    check_path

    echo ""
    info "Next steps:"
    echo "  1. Restart your shell or run: source ~/.bashrc (or ~/.zshrc)"
    echo "  2. Run: mere init"
    echo "  3. Use /module to create your first deep module"
    echo ""
}

# Run main function
main
